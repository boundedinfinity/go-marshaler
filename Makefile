makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -rf objecttype/*.enum.go
	rm -rf stringformat/*.enum.go

build:
	go build

install:
	go install

generate:
	go generate ./...

commit:
	git add . || true
	git commit -m "$(m)"
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish: generate
	make commit m=$(tag)
	make tag tag=$(tag)