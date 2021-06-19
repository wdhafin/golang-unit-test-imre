test:
	go test -v -cover ./...

test-out:
	go test -v -cover -coverprofile=cover.out ./...

bench:
	go test -v -run=TestNoTest -bench=. ./...

push:
	go fmt ./... && git push origin HEAD

push-force:
	go fmt ./... && git push -f origin HEAD