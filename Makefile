.PHONY: build
build:
	@go build -o httpget main.go 


.PHONY: build-win
build-win:
	@GOOS=windows go build -o httpget.exe main.go 

.PHONY: run
run:
	@go run main.go


.PHONY: test
test: build
	./httpget --url https://http.cat/201.jpg -o test.jpg