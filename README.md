<img src="/public/images/cherry.png" alt="Purple cherry" align="right">

# go-purple &middot; [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
> A simple Gin project scaffolding with MongoDB and JWT Auth

This project aims to create a reliable and reusable scaffolding for future server applications making use of Gin. It supports connecting to a MongoDB instance and interacting with it using the official driver for Go. Authorization is implemented with JSON Web Tokens (JWT) that are digitally signed, thus securing information exchange with minimal performance overhead.

As a secondary purpose, go-purple is an attempt at structuring server-side logic in a way consistent with existing documentation, but also one that I personally prefer and find maintainable.

## Getting started

To test go-purple for yourself, or use it as the base for an actual project, clone it in a directory of your choice, and start it up:

```shell
git clone https://github.com/Raffy27/go-purple.git
cd go-purple

go run .
```

The first run is expected to take a little longer. Go should automatically download and add all the required dependencies. You may also build it first, and then use the executable (some firewalls act friendlier this way).

This will launch go-purple in debug mode. The default port is **3000**. The server instance should be listening on all available interfaces, you may access the test page at http://localhost:3000/ to verify if everything is working correctly.

## Developing

### Built With
* Go 1.16.5
* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [MongoDB](https://www.mongodb.com/)
* [JWT](https://jwt.io/)

<!--
### Prerequisites
What is needed to set up the dev environment. For instance, global dependencies or any other tools. include download links.


### Setting up Dev

Here's a brief intro about what a developer must do in order to start developing
the project further:

```shell
git clone https://github.com/your/your-project.git
cd your-project/
packagemanager install
```

And state what happens step-by-step. If there is any virtual environment, local server or database feeder needed, explain here.

### Building

If your project needs some additional steps for the developer to build the
project after some code changes, state them here. for example:

```shell
./configure
make
make install
```

Here again you should state what actually happens when the code above gets
executed.

### Deploying / Publishing
give instructions on how to build and release a new version
In case there's some step you have to take that publishes this project to a
server, this is the right time to state it.

```shell
packagemanager deploy your-project -s server.com -u username -p password
```

And again you'd need to tell what the previous code actually does.

## Versioning

We can maybe use [SemVer](http://semver.org/) for versioning. For the versions available, see the [link to tags on this repository](/tags).


## Configuration

Here you should write what are all of the configurations a user can enter when using the project.

## Tests

Describe and show how to run the tests with code examples.
Explain what these tests test and why.

```shell
Give an example
```

## Style guide

Explain your code style and show how to check it.

## Api Reference

If the api is external, link to api documentation. If not describe your api including authentication methods as well as explaining all the endpoints with their required parameters.


## Database

Explaining what database (and version) has been used. Provide download links.
Documents your database design and schemas, relations etc...
-->

## License

This project is licensed under the MIT License. You may use it for any purpose, including commercial ones. See the [LICENSE](/LICENSE) file for more details.
