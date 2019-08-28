all : dexter migrate web

dexter : *.go cmd/dexter/main.go api/alerts/alerts.pb.go api/data/data.pb.go
	go build -o dexter cmd/dexter/main.go

migrate : *.go cmd/migrate/main.go
	go build -o migrate cmd/migrate/main.go

api/alerts/alerts.pb.go : api/alerts/alerts.proto
	cd api; protoc -I alerts/ alerts/alerts.proto --go_out=plugins=grpc:alerts

api/data/data.pb.go : api/data/data.proto
	cd api; protoc -I data/ data/data.proto --go_out=plugins=grpc:data

web : demo/src/data_grpc_web_pb.js demo/src/data_pb.js

demo/src/data_grpc_web_pb.js demo/src/data_pb.js : api/data/data.proto
	protoc -I api/data data.proto --js_out=import_style=commonjs:demo/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:demo/src

clean :
	rm -f dexter migrate
