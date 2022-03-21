run:
	go run .

build:
	go build -o /bin/biblika-api .

protobuf:
	protoc [PATH TO PROTO FILE] \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative

gateway:
	protoc [PATH TO PROTO FILE] \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=paths=source_relative \
		--grpc-gateway_opt=generate_unbound_methods=true

# mkdir -p google/api
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto
