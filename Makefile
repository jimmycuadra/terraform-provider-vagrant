FILES ?= ./...

test:
	go test $(FILES) -parallel=4

fmt:
	go fmt $(FILES)
