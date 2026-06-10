build:
	go build -o govuk-job-request-annotator-bin .

run:
	go run gov-uk-job-request-annotator-bin .

test:
	go test -race -covermode=atomic -coverprofile=c.out -v ./...

coverage:
	make test
	go tool cover -html=c.out

fmt:
	go fmt ./...

