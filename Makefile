test:
	go test ./... -covermode=count -coverprofile=.coverprofile.out
	go tool cover -func=.coverprofile.out
