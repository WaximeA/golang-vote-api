# Golang vote API

A GOLang API that provide routes to vote as a user.

## API Requirements

[API Requirements](api-requirements.md)

## Run app :
Simply use this command on the project root : `$ docker-compose up --build`

## Routes : 

First you have to login :
- `POST http://localhost:8080/login` you can test with body
```
{
	"username":	"mike"
}
```
It will give you a token useful for other routes. 

Response example : 
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1MTQ4MDgwMDAsInVzZXIiOiJtaWtlIn0.DK2NUAAnDq_0wC9a7NufyhTri0g7f1DZJjPG1kHy6mA","message":"logged in"}
```

Then, **for each route** you need to specify the Authorization bearer into the header :
`Authorization: Bearer {TOKEN}`

- Home : `GET http://127.0.0.1:8080/` 

- Create user : `POST http://127.0.0.1:8080/users` with body : 
```json
{
	"Id":       	2,
	"access_level": 1,
	"first_name":	"John",
	"last_name": 	"DOE",
	"email":		"john.doe@gmail.com",
	"user_votes": 		0,
	"password": "pass",
	"birth_date": "10-10-2000"
}
```

- Get users : `GET http://127.0.0.1:8080/users`

- Get user : `GET http://127.0.0.1:8080/users/1`

- Update user : `PATCH http://127.0.0.1:8080/users/1` with body for example
```
{
	"first_name":	"Jack"
}
```

- Delete user : `DELETE http://127.0.0.1:8080/users/5`

- Create vote : `POST http://127.0.0.1:8080/votes` with body for example 
```
{
	"Id":       	2,
	"Title": "Second vote",
	"Desc":	"Second desc"
}
```

- Get votes : `GET http://127.0.0.1:8080/votes` 

- Get vote : `GET http://127.0.0.1:8080/votes/1`