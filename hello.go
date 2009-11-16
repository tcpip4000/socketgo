package main

import (
	"fmt";
	"numbers";
)

func main() {
	var real,imaginary float;
	var msg = "The ";
	var msg1 = "result is: ";
	var option bool = true;
	
	real = 1;
	imaginary=87;
	total := real + imaginary;
	//totalStr := string(total);
	msg = msg + msg1;
	fmt.Println(msg, total, option);
	fmt.Printf("%s: %f: %t\n", msg, total, option);
	fmt.Println(numbers.Hello);
	var i numbers.MyComplex;
	i.SetReal(23);
	fmt.Println(i.GetReal());
	fmt.Println(numbers.Msg1);

}

