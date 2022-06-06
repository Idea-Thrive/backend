all: build image run
	echo "backed is up running on http://localhost:8080"

build:
	CGO_ENABLED=0 GOOS='linux' GOARCH='amd64' go build main.go

image:
	docker build -t backend:latest .

run:
	docker run -p 8080:8080 backend:latest

lint:
	golangci-lint run ./... --timeout=4m --tests=false
