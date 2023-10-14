# Welcome to Godo!

Godo is a simple TODO web service written in Go where users can:
1. Sign in using 3rd party OAuth services, in this example, it is **github** only.
2. Add a TODO item.
3. Delete a TODO item.
4. List all TODO items.
5. Mark a TODO item as completed.

This application utilizes following frameworks:
- [gin](https://github.com/gin-gonic/gin) for web server
-  [gorm](https://gorm.io/index.html) for ORM

## Installing

## Running

## Testing

## Modules

### TODO module
This is the main TODO resource, it communicates with the data repository to do actual CRUD operations.

### Auth module
Provide OAuth authentication service, can be extended to include more service providers. This example will provide 

### User module
Manages User resource.

### Main module
Entry point of the web service, it bootstraps the application with web server, registers the route handlers and ensure all web requests are proper authenticated.

## Security
### Request authentication
JWT token will be used to authenticate each web requests. Once authenticated with OAuth service provider, a user object will be retrieved/ or created, and generate a long lived JWT token.

### Audit trails
For simplicity, this project only logs all web requests where:
	- GET/ DELETE operations: Logs user, route and query params
	- POST/ PATCH/ DELETE operations: Logs user, route, query params and body


## Out of scope
- User RBAC, eg: no Admin role.
- Revisions on TODO
- JWT token refresh
