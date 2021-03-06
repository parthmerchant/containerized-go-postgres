package main

// Import packages
import (
	"net/http"
	"strconv"
  
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// List ORM struct
type List struct {

	gorm.Model

	ListID      int
	Title    	string
	Info	 	string
	Items    	[]Item
}
  

// Item ORM struct
type Item struct {

	gorm.Model
  
	Text		string
	ListID		int
}

var db *gorm.DB
var err error

var (

	lists = []List{
		{ListID: 1, Title: "Grocery", Info: "Get from Trader Joes"},
		{ListID: 2, Title: "Expenses", Info: "Pay by end of the month"},
		{ListID: 3, Title: "Guestlist", Info: "Guestlist for bday party"},
	}
	items = []Item{
		{Text: "Avocados", ListID: 1},
		{Text: "Kale", ListID: 1},
		{Text: "Phone Bill", ListID: 2},
		{Text: "Internet", ListID: 2},
		{Text: "Dan", ListID: 3},
		{Text: "Chris", ListID: 3},
		{Text: "Marc", ListID: 3},
		{Text: "Linda", ListID: 3},
	}
)


func main() {

	e := echo.New()
  
	db, err = gorm.Open( "postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")
  
	if err != nil {
	  panic("Failed to connect database.")
	}
  
	defer db.Close()
  
	db.AutoMigrate(&List{})
	db.AutoMigrate(&Item{})
  
	for index := range items {
		db.Create(&items[index])
	}
  
	for index := range lists {
		db.Create(&lists[index])
	}

	e.GET("/items", GetItems)
	e.GET("/lists/:id", GetList)
	e.GET("/lists", GetLists)

	e.POST("/items", CreateItem)
	e.POST("/lists", CreateList)

	e.DELETE("/items/:id", DeleteItem)
	e.DELETE("/lists/:id", DeleteList)
  
	handler := cors.Default().Handler(e)
  
	e.Logger.Fatal(http.ListenAndServe(":8080", handler))
}


func GetItems(c echo.Context) error {
	var items []Item
	db.Find(&items)
	return c.JSON(http.StatusOK, &items)
}

func GetList(c echo.Context) error {
	var list List
	var items []Item

	id := c.Param("id")

	db.Find(&list, id)
	db.Where("list_id = ?", list.ListID).Find(&items)

	list.Items = items
	
	return c.JSON(http.StatusOK, &list)
}

func GetLists(c echo.Context) error {
	var lists []List
	var response_list []List
	db.Find(&lists)
	
	for i := 0; i < len(lists); i++ {
		var items []Item
		var list List
		id := lists[i].ListID
		db.Find(&list, id)
		db.Where("list_id = ?", list.ListID).Find(&items)
		list.Items = items
		response_list = append(response_list, list)
	}
	
	return c.JSON(http.StatusOK, &response_list)
}

func CreateItem(c echo.Context) error {
	listid, _ := strconv.Atoi(c.QueryParam("listid"))
	text := c.QueryParam("text")
	item := Item{Text: text, ListID: listid}
	db.Create(&item)
	return c.JSON(http.StatusOK, &item)
}

func CreateList(c echo.Context) error {
	listid, _ := strconv.Atoi(c.QueryParam("listid"))
	title := c.QueryParam("title")
	info := c.QueryParam("info")
	
	list := List{ListID: listid, Title: title, Info: info}
	db.Create(&list)
	return c.JSON(http.StatusOK, &list)
}

// 
func DeleteItem(c echo.Context) error {
	var item Item
	id := c.Param("id")
	db.Find(&item, id)
	db.Delete(&item)

	var items []Item
	db.Find(&items)
	return c.JSON(http.StatusOK, &items)
}

// Delete List and all associated Items
// given listid 
func DeleteList(c echo.Context) error {
	var list List
	var items []Item
	id := c.Param("id")
	db.Find(&list, id)
	db.Delete(&list)

	db.Where("list_id = ?", id).Delete(&items)
	return c.JSON(http.StatusOK, "Deleted List")
}
