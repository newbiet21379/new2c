# new2c

# An Golang project using MongoDB as database and can use Docker to build 

**Use Docker**
1. `docker build --tag new2c . ` - Build docker image
2.  `docker run --publish 8080:8080 new2c` - Deploy docker container on port 8080

# Setup
1. `go get`  - install all dependency
2. `go build server.go` - Run the project on port 8080

**Have a demo Dockerfile for deploy container purpose**

Github: https://github.com/newbiet21379
Author: ***trungpq***

Source:
- [CRUD Source](https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c)
- [Go Basic](https://go.dev/tour/basics/1)
- [Go Gin Example](https://go.dev/doc/tutorial/web-service-gin)

