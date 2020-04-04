GOCOMP=go

shell: main.go config.json
	$(GOCOMP) build -o shell