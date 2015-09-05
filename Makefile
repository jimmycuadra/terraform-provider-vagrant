FILES ?= ./...

test: fmt
	go test $(FILES) -parallel=4

fmt:
	go fmt $(FILES)

install: fmt
	go install

plan:
	terraform plan

apply:
	terraform apply
