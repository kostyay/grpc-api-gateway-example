
proto-api:
	@echo "--> Generating gRPC clients for API"
	@rm -rf `pwd`/api-gw/api/goclient
	@docker run -v `pwd`/api-gw/api:/api -v `pwd`/api-gw/api/goclient:/goclient jfbrandhorst/grpc-web-generators protoc -I /api \
	  --go_out=plugins=grpc,paths=source_relative:/goclient \
	 v1/api_orders_service.proto v1/api_users_service.proto
	@echo "Done"

proto-users:
	@echo "--> Generating gRPC clients for Users"
	@rm -rf `pwd`/users/api/goclient
	@docker run -v `pwd`/users/api:/api -v `pwd`/users/api/goclient:/goclient jfbrandhorst/grpc-web-generators protoc -I /api \
	  --go_out=plugins=grpc,paths=source_relative:/goclient \
	 v1/users_service.proto
	@echo "Done"

proto-orders:
	@echo "--> Generating gRPC clients for Orders"
	@rm -rf `pwd`/orders/api/goclient
	@docker run -v `pwd`/orders/api:/api -v `pwd`/orders/api/goclient:/goclient jfbrandhorst/grpc-web-generators protoc -I /api \
	  --go_out=plugins=grpc,paths=source_relative:/goclient \
	 v1/orders_service.proto
	@echo "Done"


proto: proto-api proto-orders proto-users

build:
	mkdir -p ./out
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/apigw ./api-gw
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/users ./users
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/orders ./orders
	docker-compose build

clean:
	rm -rf ./out

run-servers:
	@echo "--> Starting servers"
	@docker-compose up