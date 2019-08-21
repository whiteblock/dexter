lib/data_grpc_web_pb.js lib/data_pb.js: ../../proto/data.proto lib
	protoc -I ../../proto data.proto --js_out=import_style=commonjs:lib --grpc-web_out=import_style=commonjs,mode=grpcwebtext:lib

lib:
	mkdir lib

clean:
	rm lib/data_grpc_web_pb.js
	rm lib/data_pb.js
