main:
	6g -o serversocket.6  serversocket.go
	6l -o serversocket serversocket.6
	6g -I pkg/numbers/ -o socketgo.6  socketgo.go
	6l -L pkg/numbers -o socketgo socketgo.6

clean:
	rm -rf socketgo.6 socketgo serversocket.6 serversocket

