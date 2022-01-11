.PHONY: test
test: data.go
	golangci-lint run
	go test -v -covermode=count -coverprofile=coverage.out


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


