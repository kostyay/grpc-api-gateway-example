# grpc-api-gateway-example

A simple example demonstrating API Gateway pattern using gRPC. For a blog post at << insert blog post address >>.

The repo contains 3 microservices:
* `users` - Internal Users microservice
* `orders` - Internal Orders microservice
* `api-gw` - External API gateway microservice
* `apigw-client` - A simple client calling ListOrdersWithUser endpoint.

The api gateway is exposed externally and offers public api. Any call to the api gateway translates into 1 or multiple 
requests to internal microservices.


## Working with the project
* `make build` - Will build the binaries for each of the microservices
* `docker-compose up` - Will start all the microservices

