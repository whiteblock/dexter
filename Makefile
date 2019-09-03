all : dexter migrate web docs

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

docs : docs/alerts.md docs/data.md

docs/alerts.md : docs/alerts.header.md docs/alerts.body.md
	cat docs/alerts.header.md docs/alerts.body.md > docs/alerts.md

docs/alerts.body.md : api/alerts/alerts.proto
	protoc -I api/alerts alerts.proto --doc_out=./docs --doc_opt=markdown,alerts.body.md

docs/data.md : docs/data.header.md docs/data.body.md
	cat docs/data.header.md docs/data.body.md > docs/data.md

docs/data.body.md : api/data/data.proto
	protoc -I api/data data.proto --doc_out=./docs --doc_opt=markdown,data.body.md

clean-docs :
	rm -f docs/alerts.md docs/data.md

clean : clean-docs
	rm -f dexter migrate
