compile:
	protoc \
		--proto_path=plumber \
		--go_out=plumber \
		--go-grpc_out=plumber \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		plumber/*.proto

# docker-build-dev:
#	docker build -t muxfarm:dev ./
