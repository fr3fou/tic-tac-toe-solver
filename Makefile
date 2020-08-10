static-build:
	go build -a -ldflags '-extldflags "-static"' .