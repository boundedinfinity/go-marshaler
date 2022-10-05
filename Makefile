makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -rf objecttype/*.enum.go
	rm -rf stringformat/*.enum.go

build: generate
	go build

generate: purge
	go generate ./...

install: generate
	go install

test: generate
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -fa $(tag) -m "$(tag)"
	git push -f origin $(tag)

publish: test
	make commit m=$(m)
	make tag tag=$(m)