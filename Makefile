suni: *.go
	go build -o suni

sunix: *.go
	env GOOS=linux GOARCH=amd64 go build -o sunix
