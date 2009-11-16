package numbers

import (
	"testing"
	"fmt"
)

func TestComplexNumber(t *testing.T) {
	var c MyComplex;
	var real float := 34;
	var img float := 51;
	c.SetReal(real);
	c.SetImaginary(img);
	fmt.Println("r: ", c.GetReal, "c: ", c.GetImaginary);

	if c.GetReal() != real { 
		t.Errorf("Real returned %q, should be %q", c.GetReal(), real); 
	}
}

