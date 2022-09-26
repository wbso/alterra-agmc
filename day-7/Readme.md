# Day 7 Create Docker Image And Push To Registry

docker image url: https://hub.docker.com/r/wbso/agmc-go


task: implement Hexagonal architecture using echo framework use gorm

## diagram
> (user) ----http---> [handler] ---dto---[app] ---dao---> [repository]

## handler package
Handler package just an entry point to the application.
1. decode http request from user
2. pass data to business logic
3. encode response from business logic to http response

# app package
app package contains business logic.
this package doesn't care where the request came from(http, grpc, cli).
this package doesn't care where the data is stored

# repository package

repository package only has responsibility to store and read data from one or many storages(database, 3rd-party apis)
