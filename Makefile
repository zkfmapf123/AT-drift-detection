clear:
	clear

.PHONY: build
build: clear
	go build -o at-plan main.go

run: build
	sudo mv at-plan /usr/local/bin/ && \
	at-plan plan \
	-g github_token \
	-u https://atlantis.example.com \
	-t atlantis_token \
	-r atlantis_repository \
	-c atlantis_config_file
