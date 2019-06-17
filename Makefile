all : dexter

dexter : *.go cmd/dexter/main.go
	go build -o dexter cmd/dexter/main.go

clean :
	rm -f dexter
