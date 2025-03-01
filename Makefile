sync:
	cd cities.json;git pull
	cp cities.json/cities.json data/

all:
	go fmt ./...
	golangci-lint run ./...
	go test -v ./...
	cargo test
	cd citiespy && make test
