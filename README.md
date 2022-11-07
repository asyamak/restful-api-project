## restful-api-project ##

This project is a simple RESTAPI HTTP service.

In this project I used PostgresSQL database. Docker & Docker-compose. Also you can see UNIT-tests for handlers and MOCK tests for database.
Database will contain following table "users" with fields:
{
    `id` INT PRIMARY KEY, 
    `data` VARCHAR
}
`data` field itself contains json format:
{
"first_name": "First", "last_name": "Last", "interests": "coding,golang"
}

To initialize program:
`go run main.go`

To run through docker:
`docker-compose up --build -d`
`docker exec -it postgres bash
psql -U postgres`

To check handlers use Postman with address: "localhost:4040/user/{id}" with all following methods:
POST
GET
PUT
DELETE

For PUT method updates user through values. For example:
"localhost:4040/user/1?first_name=Ola&last_name=Sun&interests=golang"

Author:
makasya
m.a_k@mail.ru