package complex

const Hello = "Hello complex!"

type MyComplex struct { real, imaginary float; }

func (c *MyComplex) GetReal() float { return c.real; }

func (c *MyComplex) SetReal(value float) { c.real = value; }

func (c *MyComplex) GetImaginary() float { return c.imaginary; }

func (c *MyComplex) SetImaginary(value float) { c.imaginary = value; }

