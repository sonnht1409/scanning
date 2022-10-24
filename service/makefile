build:
	CGO_ENABLED=0 go build -installsuffix 'static' -o ./cmd/server
server:
	cmd/server api
worker:
	cmd/server worker