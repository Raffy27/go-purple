<img src="/public/images/cherry.png" alt="Purple cherry" align="right">

# go-purple &middot; [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
> A simple Gin project scaffolding with MongoDB and JWT Auth

This project aims to create a reliable and reusable scaffolding for future server applications making use of Gin. It supports connecting to a MongoDB instance and interacting with it using the official driver for Go. Authorization is implemented with JSON Web Tokens (JWT) that are digitally signed, thus securing information exchange with minimal performance overhead.

:warning: This project is not ready for production use.

As a secondary purpose, go-purple is an attempt at structuring server-side logic in a way consistent with existing documentation, but also one that I personally prefer and find maintainable.

## Getting started

To test go-purple for yourself, or use it as the base for an actual project, clone it in a directory of your choice, and start it up:

```shell
git clone https://github.com/Raffy27/go-purple.git
cd go-purple

go run .
```

The first run is expected to take a little longer. Go should automatically download and add all the required dependencies. You may also build it first, and then use the executable (some firewalls act friendlier this way).

This will launch go-purple in debug mode. The default port is **3000**. 

The server instance should be listening on all available interfaces, you may access the test page at http://localhost:3000/ to verify if everything is working correctly.

## Developing

### Built With
* Go 1.16.5
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [MongoDB](https://www.mongodb.com/)
* [JWT](https://jwt.io/)
* [Viper](https://github.com/spf13/viper)

### Setting up Dev

Use an IDE of your choice to test and tinker with go-purple. Make sure it has internal support for Go.

Also make sure to properly set up and configure an instance of MongoDB, and change the configuration file accordingly.
Main config entries are found in `config/config.json` (code can be adapted to load other formats of configuration as well).

This project uses the official Go bindings for MongoDB, so it is expected to work with future versions as well, without any significant changes.

## Structure

```ini
├── config
|   ├── config.go           # This loads and checks the configuration files
│   └── config.json         # The main configuration file
├── controllers
|   ├── auth.go             # Controller for handling authentication
│   └── user.go             # Controller for handling user data
├── forms
│   └── user.go             # Login and Register forms
├── middleware
│   └── auth.go             # Middleware for checking authorization
├── models
│   └── user.go             # User model for the database
├── public                  # Static files
├── server                  # Server entry point
│   ├── db
│   │   └── mongo.go        # Database controller with utility functions
│   ├── router.go           # This handles route setup and API Groups
│   └── server.go           # This takes care of the initialization and graceful shutdown of the app
├── tests
│   ├── future_test.go
│   ├── future_test.go
│   └── future_test.go
├── LICENSE                 
├── main.go                 # Entry point
└── README.md
```

## Configuration

Aspects of the project that have built-in configuration support are as follows:

| Key | Value | Description |
| :-: | :---: | :---------- |
| server.address | `:3000` | The address the server will listen on |
| server.maxShutdown | `5` | The maximum number of seconds to wait for the server to gracefully shut down |
| database.url | `mongodb://localhost:27017` | MongoDB connection string |
| database.main | `test` | The name of the main database to use |
| jwt.algo | `HS256` | The algorithm used to sign the JWT tokens |
| jwt.secret | `secret` | The secret used to sign the JWT tokens |
| jwt.expires | `259200` | The number of seconds before the JWT tokens expire |

## Tests

...are not yet implemented, but when they will be, you will be able to run them like so:

```shell
go test ./...
```

## API Reference

go-purple exposes multiple API endpoints by defult, mainly used for authentication and testing purposes. Here is a list of all the endpoints and their respective methods. You can test these directly by issuing the requests in `tests/*.http` files.

| Endpoint | Method | Parameters | Description |
| :------: | :----: | :--------- | :---------- |
| / | GET | | Static binding for files in `public` |
| /api/v1/auth/login | POST | username, password | Attempts to authenticate a user. If successful, it returns a web token |
| /api/v1/auth/logout | POST | | Invalidates the current session |
| /api/v1/auth/register | POST | username, password, email | Attempts to create a new user. The provided username **must** be unique |
| /api/v1/users | GET | | Returns a list of all users. Authorization is required to access this endpoint |
| /api/v1/users/`:user` | GET | | Returns a specific user. Authorization is required to access this endpoint |

Responses are formatted in JSON. A typical response will contain a `msg` property, which will contain a human-readable message. Requests that run into errors will also contain an `error` property, which will contain a reference to the actual error encountered.

## Database :cricket:

## License

This project is licensed under the MIT License. You may use it for any purpose, including commercial ones. See the [LICENSE](/LICENSE) file for more details.
