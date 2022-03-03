package computer

type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}

// if the name of struct starts with a upper case, it is called "exported class"
// and can be accessed in other packages

// the same to field

// just like public
// but it just be distinguished by case here
