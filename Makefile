all : dexter migrate

dexter : *.go cmd/dexter/main.go api/alerts/alerts.pb.go api/data/data.pb.go
	go build -o dexter cmd/dexter/main.go

migrate : *.go cmd/migrate/main.go
	go build -o migrate cmd/migrate/main.go

api/alerts/alerts.pb.go : api/alerts/alerts.proto
	cd api; protoc -I alerts/ alerts/alerts.proto --go_out=plugins=grpc:alerts

api/data/data.pb.go : api/data/data.proto
	cd api; protoc -I data/ data/data.proto --go_out=plugins=grpc:data

clean :
	rm -f dexter migrate
