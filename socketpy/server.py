import socket

socket_s = socket.socket()
socket_s.bind(("localhost", 9998))
socket_s.listen(5)

socket_c, (host_c, puerto_c) = socket_s.accept()  
while True:
	recibido = socket_c.recv(1024)  
	if recibido == "quit":
		break
	print "Recibido: ", recibido  
	socket_c.send(recibido)

print "adios"
socket_c.close()
socket_s.close()

