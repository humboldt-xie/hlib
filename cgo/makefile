all:
	go build -o c.so -buildmode=c-shared c.go cd.go
	gcc -o client c.cc ./c.so
