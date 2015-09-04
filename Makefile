FILES ?= ./...

test:
	go test $(FILES) -parallel=4

fmt:
	go fmt $(FILES)

install:
	go install

plan:
	terraform plan

apply:
	terraform apply
