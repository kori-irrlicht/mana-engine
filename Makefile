
test:
	go test -json -coverprofile=coverage.out > test-report.out
	go vet ./... > vet.out
	golint ./... > lint.out
	gometalinter ./... > metalint.out

install:
	go get -t ./...

clean:
	rm *.out
