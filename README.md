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

## License
This project is licensed under the MIT License. See the LICENSE file for details.