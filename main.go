package main

// Import packages
import (
	"encoding/json"
	"log"
	"net/http"
  
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

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
	Done 		bool
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
		{Text: "Avocados", Done: true, ListID: 1},
		{Text: "Kale", Done: false, ListID: 1},
		{Text: "Phone Bill", Done: false, ListID: 2},
		{Text: "Internet", Done: true, ListID: 2},
		{Text: "Dan", Done: false, ListID: 3},
		{Text: "Chris", Done: true, ListID: 3},
		{Text: "Marc", Done: true, ListID: 3},
		{Text: "Linda", Done: false, ListID: 3},
	}
)


func main() {

	router := mux.NewRouter()
  
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
  
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/lists/{id}", GetList).Methods("GET")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
  
	handler := cors.Default().Handler(router)
  
	log.Fatal(http.ListenAndServe(":8080", handler))
}


func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	db.Find(&items)
	json.NewEncoder(w).Encode(&items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	db.First(&item, params["id"])
}

func GetList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var list List
	var items []Item

	db.First(&list, params["id"])
	db.Model(&list).Related(&items)

	list.Items = items
	
	json.NewEncoder(w).Encode(&list)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
  	var item Item
  	db.First(&item, params["id"])
  	db.Delete(&item)

  	var items []Item
  	db.Find(&items)
  	json.NewEncoder(w).Encode(&items)
}
