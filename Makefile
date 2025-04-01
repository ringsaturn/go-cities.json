sync:
	git submodule update
	cp cities.json/cities.json data/
	cp cities.json/LICENSE data/LICENSE

all:
	go fmt ./...
	golangci-lint run ./...
	go test -v ./...
	cargo test
	cd citiespy && make test
