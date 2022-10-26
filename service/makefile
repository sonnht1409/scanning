build:
	CGO_ENABLED=0 go build -installsuffix 'static' -o ./cmd/server
server:
	cmd/server api
worker:
	cmd/server worker
dockerize-server:
	docker build -t scanning-server -f dockerfile.server .
dockerize-run-server:
	docker run -p 8080:8080 scanning-server
dockerize-worker:
	docker build -t scanning-worker -f dockerfile.worker .
dockerize-run-worker:
	docker run scanning-worker