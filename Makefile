build: 
	@echo "Building Windows Target"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/win/ ./cmd/...
	@echo "Building Linux Target"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./out/linux/ ./cmd/... 

