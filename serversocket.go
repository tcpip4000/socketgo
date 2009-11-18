package main

import (
	"fmt";
	"net";
	"os";
)

func main() {
	var host = "127.0.0.1";
	var port = "9998";
	var remote = host + ":" + port;
	//var data []byte;
	var data2 = make([]byte, 1024);

	lis, error := net.Listen("tcp", remote);
	defer lis.Close();
	if error != nil { fmt.Printf("Error creating listener: %s\n", error ); os.Exit(1); }
	for {
		con, error := lis.Accept();
		defer con.Close();
		if error != nil { fmt.Printf("Error accepting: %s\n", error); os.Exit(2); }
		for {
			n, error := con.Read(data2);
			if error != nil { fmt.Printf("Error reading: %s, %d", error, n); os.Exit(3); }
			fmt.Println(string(data2));
		}
	}

	fmt.Println("Closing server...");
}

