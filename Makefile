GO      = go
OUT     = out
GO111MODULE ?= on

DOCKER  = docker
TAG     = calini/draco:dev

run:
	$(GO) run .

build:
	$(GO) build -o $(OUT)/main .

test:
	$(GO) test -count=1 ./...

test-coverage:
	$(GO) test -count=1 -coverprofile $(OUT)/tests/cover.out . >> $(OUT)/tests/test.json

docker:
	$(docker) build -t $(TAG) .

