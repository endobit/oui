.PHONY: test
test: data.go
	go test -v -covermode=count -coverprofile=coverage.out

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


