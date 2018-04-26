setup:
	cd goslib && protoc -I src/gos_rpc_proto --go_out=plugins=grpc:src/gos_rpc_proto src/gos_rpc_proto/*.proto
	cd game && ./tools/gen_routes
	cd game && ./tools/gen_protocol
	cd game && bundle exec rake generate_tables

build:
	sh build_gos.sh

dep_install:
	sh dep_install.sh

rpc_proto:
	cd goslib && protoc -I src/gos_rpc_proto --go_out=plugins=grpc:src/gos_rpc_proto src/gos_rpc_proto/*.proto

tcp_protocol:
	cd game && ./tools/gen_routes
	cd game && ./tools/gen_protocol

generate_tables:
	cd game && bundle exec rake generate_tables

build_linux:
	sudo docker run --rm -v $(shell pwd):/usr/src/gos -w /usr/src/gos -e GOOS=linux -e GOARCH=amd64 golang:latest make build