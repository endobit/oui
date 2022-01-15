.PHONY: test
test: data.go
	go test -v -coverprofile=coverage.txt -covermode=count ./...

.PHONY: lint
lint:
	golangci-lint run

data.go: oui.txt
	go build ./cmd/gen
	go generate .

oui.txt:
	curl https://linuxnet.ca/ieee/oui/nmap-mac-prefixes -o $@

.PHONY: nuke
nuke:
	-rm oui.txt

.PHONY: clean
clean:
	-rm gen


