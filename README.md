
  

  

# Welcome to Godo!

Godo is a simple TODO web service written in Go where users can:

1. Sign in using 3rd party OAuth services, in this example, it is **github** only.
2. Add a TODO item.
3. Delete a TODO item.
4. List all TODO items.
5. Mark a TODO item as completed.


This application utilizes following frameworks:

- [gin](https://github.com/gin-gonic/gin) for web server
- [gorm](https://gorm.io/index.html) for ORM

  
  

## Installing

1. Build the docker file ```docker build -t godo .```
2. Copy the .env.sample into .env , adjust the environment variables accordingly

  
  

## Running

Demo JWT token (if your JWT_SECRET is godo): ```eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoibG9jYWw6MSIsIm5hbWUiOiJKb2huIERvZSJ9LCJleHAiOjE3Mjg5NzE3OTB9.8x3H6eq9gNKhc52XbHPXmGN1nju4_f6gaRiKe9fgyiM```

  

### Start up server

```
docker-compose up -d
```
The TODO web service will listen at `http://localhost:8080`

  
  

## API Endpoints

### Request for Github OAuth verifier url
**Endpoint**: POST /auth/logins
**Body**:
```json
{
	"data":{
		"strategy": "github"
	}
}
```
**Response**
```json
{
	"data": "https://github.com/login/oauth/authorize?client_id=&response_type=code&state=state"
}
```


### Verify Github OAuth callback
  **Endpoint**: POST /auth/logins
**Body**:
```json
{
	"data":{
		"strategy": "githubVerify",
		"verificationCode": "234f12bd95774e8f370a"
	}
}
```
**Response**
```json
{
	"data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoiZ2l0aHViOjczMzgxNzMiLCJuYW1lIjoiQ2hvb2kgS2FoIFdhaSJ9LCJleHAiOjE3Mjg5ODM4NDgsInN1YiI6ImdpdGh1Yjo3MzM4MTczIn0.aUp5UPLsJ6lJGnepOI88Yao_TfM8EjM9JaaHYBAQjbE"
}
```

### Add a TODO item
**Endpoint**: POST /todos
 **Authorization Header**: Bearer JWT_TOKEN
**Body**:
```json
{
	"data":{
		"title": "Commit to Github"
	}
}
```
**Response**
```json
{
	"data": {
	"id": 1,
	"title": "Commit to Github",
	"completed": false,
	"userId": "github:7338173",
	"createdAt": "2023-10-15T16:48:02.4211646+08:00",
	"updatedAt": "2023-10-15T16:48:02.4211646+08:00"
	}
}
```

### List all TODO items
  **Endpoint**: GET /todos
  **Authorization Header**: Bearer JWT_TOKEN
**Response**
```json
{
  "data": [
    {
      "id": 1,
      "title": "hello there",
      "completed": false,
      "userId": "github:7338173",
      "createdAt": "2023-10-15T16:48:02Z",
      "updatedAt": "2023-10-15T16:48:02Z"
    },
    {
      "id": 2,
      "title": "Commit to Github",
      "completed": false,
      "userId": "github:7338173",
      "createdAt": "2023-10-15T09:24:35Z",
      "updatedAt": "2023-10-15T09:24:35Z"
    },
  ]
}
```
  
  
### Delete a TODO item
  **Endpoint**: DELETE /todos/:id
  **Authorization Header**: Bearer JWT_TOKEN
**Response**
```json
{
	"id": 1
}
```

### Mark a TODO item as completed
**Endpoint**: PATCH /todos/:id
 **Authorization Header**: Bearer JWT_TOKEN
**Body**:
```json
{
	"data":{
		"completed": true
	}
}
```
**Response**
```json
{
  "data": {
    "id": 1,
    "title": "hello there",
    "completed": true,
    "userId": "github:7338173",
    "createdAt": "2023-10-15T16:48:02Z",
    "updatedAt": "2023-10-15T16:48:02Z"
  }
}
```


## Testing

  

  

  

## Modules

### TODO module

This is the main TODO resource, it communicates with the data repository to do actual CRUD operations.

  

### Auth module

Provides OAuth authentication service, can be extended to include more service providers. This example includes github authentication only. Once authenticated, a JWT token will be generated with User in the payload.

  

User is a simple logical object with fields {ID, Name} where id is in the format of "provider:remoteId". eg: "github:123456"

### Main module

Entry point of the web service, it bootstraps the application with web server, registers the route handlers and ensure all web requests are proper authenticated.

  
  

## Security

### Request authentication

JWT token will be used to authenticate each web requests. Once authenticated with OAuth service provider, a user object will be retrieved/ or created, and generate a long lived JWT token.

  
  

### Audit trails

For simplicity, this project only logs all web requests where:

- GET/ DELETE operations: Logs user, route and query params
- POST/ PATCH operations: Logs user, route, query params and body

  

## Out of scope

- User management
- Revisions on TODO
- JWT token refresh
- Pagination & dynamic filters for result set