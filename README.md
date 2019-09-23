# Golang vote API

A GOLang API that provide routes to vote as a user.

## API Requirements

[API Requirements](api-requirements.md)

## Run app : 
- build project : `$ go build`
- init modules : `$ go mod init`
- run the app : `$ go run main.go`
- Then go on http://localhost:8001 on postman and see "Welcome home!" message in response block.

## Routes : 

- Home : `GET http://localhost:8001/`
- Create user : `POST http://127.0.0.1:8001/users` with body : 
```json
{
	"Id":       	2,
	"first_name":	"John",
	"last_name": 	"DOE",
	"email":	"john.doe@gmail.com",
	"user_votes":   0
}
```
- Get users : `GET http://127.0.0.1:8001/users` 
