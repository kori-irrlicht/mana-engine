
test:
	go test -json -coverprofile=coverage.out > test-report.out
	go vet ./... > vet.out
	golint ./... > lint.out
	gometalinter ./... > metalint.out

install:
	go get -t ./...

devinstall:
	go get -u golang.org/x/lint/golint
	go get -u gopkg.in/alecthomas/gometalinter.v2

clean:
	rm *.out
