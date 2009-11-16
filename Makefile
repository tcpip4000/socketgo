main:
	6g -o ./complex/complex.6 ./complex/complex.go
	6g -o ./complex/msg.6 ./complex/msg.go
	6g -o hello.6 -I complex/ hello.go
	6l -o hello hello.6

clean:
	rm -rf hello.6 hello 

