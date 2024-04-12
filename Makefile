start:
	go run ./main.go

build:
	go build -o ./dist/rm-by-pattern ./main.go
	chmod +x ./dist/rm-by-pattern

install-local:
	go build -o $(GOBIN)/rm-by-pattern ./main.go
	chmod +x $(GOBIN)/rm-by-pattern

install:
	go install github.com/asolopovas/dsync@latest

test:
	 go run ./main.go -r ./dsync-config.json

tag-push:
	$(eval VERSION=$(shell cat version))
	git tag $(VERSION)
	git push origin $(VERSION)
	if git rev-parse latest >/dev/null 2>&1; then git tag -d latest; fi
	git tag latest
	git push origin latest --force
