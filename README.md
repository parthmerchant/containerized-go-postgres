# List and Items Web Service
--------------------------------------
Go, PostgreSQL, Docker

A simple REST endpoint for querying and modifying lists and their items.


## Setup
Run the following command to build the application using Docker:
```
docker compose up
```

## Usage 

There are currently 3 REST API endpoints:
```
GET /items
GET /lists/{id}
DELETE /items/{id}
```

If you want to hit these endpoints, use the following `curl` commands to perform the aforementioned GET and DELETE: 

1. `GET /items`
```
curl -X GET localhost:8080/items
```

2. `GET /lists/{id}`
```
curl -X GET localhost:8080/lists/{id}
```

3. `DELETE /items/{id}
```
curl -X DELETE localhost:8080/items/{id}
```
