package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()
	vm.Run(`
	    abc = 2 + 2;
		console.log("The value of abc is " + abc); // 4
	`)
	if value, err := vm.Get("abc"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			fmt.Println(value_int, err)
		}
	}
	vm.Set("def", 11)
	vm.Run(`
	    console.log("The value of def is " + def);
		// The value of def is 11
	`)
	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
	    console.log(xyzzy.length); // 16
	`)
}
