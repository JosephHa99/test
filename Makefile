build:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go install -v -a -installsuffix cgo
