package main

import (
	"fmt";
	"net";
	"os";
	"io";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		data = make([]byte, 10);
	)
	fmt.Println("Initiating server... (Ctrl-C to stop)");

	lis, error := net.Listen("tcp", remote);
	defer lis.Close();
	if error != nil { 
		fmt.Printf("Error creating listener: %s\n", error ); 
		os.Exit(1); 
	}
	for {
		var response string;
		var read = true;
		con, error := lis.Accept();
		defer con.Close();
		if error != nil { fmt.Printf("Error: Accepting data: %s\n", error); os.Exit(2); }
		for read {
			n, error := con.Read(data);
			switch error { 
			case io.ErrUnexpectedEOF:
				fmt.Printf("Warning: End of data reached: %s \n", error); 
				read = false;
			case nil:
				//fmt.Println(string(data[0:n])); // Debug
				response = response + string(data[0:n]);
			default:
				fmt.Printf("Error: Reading data : %s \n", error); 
				read = false;
			}
		}
		fmt.Println("Full response: " + response); 
	}
}

