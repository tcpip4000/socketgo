main:
	6g -I pkg/numbers/ -o hello.6  hello.go
	6l -L pkg/numbers -o hello hello.6

clean:
	rm -rf hello.6 hello 

