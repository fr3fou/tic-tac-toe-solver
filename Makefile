static-build:
	go build -ldflags '-extldflags "-static"' .