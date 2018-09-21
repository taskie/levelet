.PHONY: build install test coverage dep-ensure dep-graph pre-commit

CMD_DIR := cmd/levelet

build:
	go build -v -ldflags "-s -w -X fwv.Revision=$(shell git rev-parse --short HEAD)"
	$(MAKE) -C $(CMD_DIR) build

install:
	go install -v -ldflags "-s -w -X fwv.Revision=$(shell git rev-parse --short HEAD)"
	$(MAKE) -C $(CMD_DIR) install

test:
	go test

coverage:
	mkdir -p test/coverage
	go test -coverprofile=test/coverage/cover.out
	go tool cover -html=test/coverage/cover.out -o test/coverage/cover.html

dep-ensure:
	dep ensure
	dep status

dep-graph:
	mkdir -p images
	dep status -dot | dot -Tpng -o images/dependency.png

pre-commit:
	$(MAKE) build
	$(MAKE) coverage
	rm -rf vendor/
	$(MAKE) dep-ensure
	$(MAKE) dep-graph
