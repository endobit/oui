.PHONY: test
test: data.go
	go test -v -coverprofile=coverage.txt -covermode=count ./...

.PHONY: lint
lint:
	golangci-lint run

data.go: oui.csv
	go build ./cmd/gen
	go generate .

oui.csv:
	curl https://standards-oui.ieee.org/oui/oui.csv -o oui.csv

.PHONY: nuke
nuke:
	-rm oui.csv

.PHONY: clean
clean:
	-rm gen


