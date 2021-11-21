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

There are currently 4 REST API endpoints:
```
GET /items
GET /items/{id}
GET /lists/{id}
DELETE /items/{id}
```
