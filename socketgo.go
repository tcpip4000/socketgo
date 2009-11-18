package main

import (
	"fmt";
	"net";
	"os";
	"strings";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		msg string = "hola mundo2 y otras cosas más";
	)

	con, error := net.Dial("tcp", "", remote);
	defer con.Close();
	if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }

	in, error := con.Write(strings.Bytes(msg));
	if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }

	fmt.Println("Connection OK");

}


