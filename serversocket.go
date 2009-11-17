package main

import (
	"fmt";
	"net";
	"os";
)

func main() {
	var host = "127.0.0.1";
	var port = "8764";
	var remote = host + ":" + port;

	lis, error := net.Listen("tcp", remote);
	if error != nil { fmt.Printf("Error creating listener: %s\n", error ); os.Exit(1); }
	for {
	}

	lis.Close();
	fmt.Println("Listener closing.. ");
}

