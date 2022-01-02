# List and Items Web Service
--------------------------------------
Go, PostgreSQL, Docker

A simple REST endpoint for querying and modifying lists and list items.


## Setup
Run the following command to build the application using Docker:
```
docker compose up
```

## Usage 

There are currently 7 RESTful endpoints:
```
GET /items # Get all items
GET /lists/{id} # Get list by list id
GET /lists # Get all list and associated items

POST /items # Create item 
POST /lists # Create list

DELETE /items/{id} # Delete item by id 
DELETE /lists/{id} # Delete list by list id 
```

If you want to hit these endpoints, use the following `curl` commands to perform the aforementioned GET and DELETE: 

### `GET`
---------

1. `GET /items`
```
curl -X GET http://localhost:8080/items
```

2. `GET /lists/{id}`
```
curl -X GET http://localhost:8080/lists/{id}
```

3. `GET /lists`
```
curl -X GET http://localhost:8080/lists
```

### `POST`
----------

1. `POST /items`
```
curl -X POST http://localhost:8080/items\?text\=sampletext\&listid\=4
```

2. `POST /lists`
```
curl -X POST http://localhost:8080/lists\?listid\=samplelist\&title\=sampletitle\&info\=sampleinfo
```

### `DELETE`
-------

1. `DELETE /items/{id}`
```
curl -X DELETE http://localhost:8080/items/{id}
```

2. `DELETE /lists/{id}`
```
curl -X DELETE http://localhost:8080/lists/{id}
```
