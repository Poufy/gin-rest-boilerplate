# Gin Boilerplate
This is a boilerplate project for building RESTful APIs using the Gin Web Framework in Go. The project structure follows best practices and separates concerns such as configuration, routing, controllers, services, and data models.

## Project Structure
```go
my_project/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go
│   └── config.yaml
├── controllers/
│   ├── user_controller.go
│   └── product_controller.go
├── middlewares/
│   ├── auth.go
│   └── logger.go
├── models/
│   ├── user.go
│   └── product.go
├── routes/
│   ├── user_routes.go
│   ├── product_routes.go
│   └── api.go
├── services/
│   ├── user_service.go
│   └── product_service.go
├── repositories/
│   ├── user_repository.go
│   └── product_repository.go
├── storage/
│   └── database.go
├── go.mod
└── go.sum
```
## Getting Started
### Prerequisites
* Go (version 1.17 or higher)

### Setup
1. Clone this repository:
```sh
git clone git@github.com:Poufy/gin-rest-boilerplate.git
```

2. Navigate to the project directory:
```sh
cd gin-rest-boilerplate.git
```
3. Update the config.yaml file in the config/ directory with your desired configuration settings.


## Running the server
1. Run the server:
```sh
go run cmd/server/main.go
```

The server will listen on port 8080 by default. You can change the listening port in the config.yaml file.

## Endpoints
The boilerplate includes the following example endpoints:

* `GET /api/users`: Retrieves a list of users.
* `GET /api/products`: Retrieves a list of products.

## Customization
To customize the boilerplate for your own project, you can add, modify, or remove controllers, services, models, routes, and middlewares as needed. Update the config.yaml file with your specific configuration settings, such as database connection details. Replace the sample data in the models package with your own data structures and data sources.

## Layers
### Data Flow
```css
[HTTP Request] -> [Controller] -> [Service/Domain Logic] -> [Repository] -> [Database]
[HTTP Response] <- [Controller] <- [Service/Domain Logic] <- [Repository] <- [Database]
```
### Controllers
Controllers are responsible for handling the HTTP request/response flow in our application. They receive incoming requests, interpret them, and invoke the necessary logic to process the request. Controllers act as intermediaries between the client and the underlying application logic. They extract data from the request, perform any necessary validations, and coordinate the flow of data to and from the services. Controllers are typically associated with specific endpoints and define the actions to be taken for each endpoint.

### Routes
Routes define the available endpoints and the mapping between URLs and their corresponding controller methods. They provide a clear structure for the API and determine how incoming requests are routed to the appropriate controller. Routes enable us to handle different HTTP methods (such as GET, POST, PUT, DELETE) and define the necessary parameters. By defining routes, we establish the API's contract and map URLs to specific actions, ensuring that requests are handled by the correct controller methods.

### Repositories
Repositories encapsulate the data access logic in our application. They provide an abstraction layer that separates the database operations from the rest of the application. Repositories handle the interaction with the database, including CRUD (Create, Read, Update, Delete) operations, complex queries, and any other data-related functionalities. By encapsulating the data access logic, repositories promote reusability, modularity, and maintainability. They allow us to interact with the database using a consistent interface and abstract away the underlying database details.

### Services
Services contain the business logic of our application. They implement and encapsulate complex operations, validations, transformations, and any other domain-specific logic related to the entities we are working with. Services act as mediators between the controllers and the repositories. They coordinate the flow of data between the repositories, perform business logic operations that may span multiple repositories, and apply any necessary validations or transformations. Services provide a higher level of abstraction, allowing us to centralize and reuse business logic across different controllers or endpoints. They promote modularity, code organization, and maintainability by separating the business logic from the controllers and repositories.


## License
This project is licensed under the MIT License. See the LICENSE file for details.