# go-timetable-microservice

This is a microservice that handles the authentication and authorisation of a user. This microservice is built in conjunction with the following:

- [API Gateway](https://github.com/Bryan-BC/go-api-gateway)
- [Course Service](https://github.com/Bryan-BC/go-course-microservice)
- [Timetable Service](https://github.com/Bryan-BC/go-timetable-microservice)

## Prerequisites

Since this service utilises gRPC and protocol buffers, it requires the [protocol buffer compiler](https://grpc.io/docs/protoc-installation/) installed. Moreover, it also uses `Makefile` to compile the protobuf files. This means it also requires [make](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows) to be installed

## Setup

To run the service, simply cd into the repository and run the following command:

`make start`

Once this is run, the gateway will be running on port 4000.
