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

	var ar = [5]int {1, 2, 3};
	var sl = ar[0:2];
	var sl2 = make([]int, 100);
	var sl3 [100]int;
	var m = map[string]int {"hola":1, "dos":2};
	fmt.Println(sl, len(ar), len(sl), cap(sl), len(sl2), len(sl3), len(m));
	v, ok := m["toto"];
	if ok == false {
		fmt.Println("not found key");
	} else {
		fmt.Println(v);
	}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value);
	}
	m["dos"] = 0, false;
	fmt.Println(len(m));
	
	real = 1;
	imaginary=87;
	total := real + imaginary;
	msg = msg + msg1;
	fmt.Println(msg, total, option);
	fmt.Printf("%s: %f: %t\n", msg, total, option);
	fmt.Println(numbers.Hello);
	var i numbers.MyComplex;
	var j *numbers.MyComplex= new(numbers.MyComplex);
	i.SetReal(23);
	j = &i;
	i.SetReal(12);
	fmt.Println(i.GetReal(), j.GetReal());
	fmt.Println(numbers.Msg1);

}

