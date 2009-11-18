package main

import (
	"fmt";
	"net";
	"os";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		data = make([]byte, 10);
	)

	lis, error := net.Listen("tcp", remote);
	defer lis.Close();
	if error != nil { 
		fmt.Printf("Error creating listener: %s\n", error ); 
		os.Exit(1); 
	}
	for {
		var response string;
		con, error := lis.Accept();
		defer con.Close();
		if error != nil { fmt.Printf("Error accepting: %s\n", error); os.Exit(2); }
		for {
			n, error := con.Read(data);
			//fmt.Println(n, error);
			if error != nil { 
				fmt.Printf("Error reading: %s, %d \n", error, n); 
				break;
			}
			fmt.Println(string(data[0:n]));
			response = response + string(data[0:n]);
		}
		fmt.Println("Full response: " + response);
	}
	fmt.Println("Closing server...");
}

