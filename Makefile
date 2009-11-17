main:
	6g -I pkg/numbers/ -o socketgo.6  socketgo.go
	6l -L pkg/numbers -o socketgo socketgo.6

clean:
	rm -rf socketgo.6 socketgo 

