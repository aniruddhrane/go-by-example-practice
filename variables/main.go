package main
import "fmt"

func main(){
	var a="initial"
	fmt.Println(a)
	var b,c int =1,2
	fmt.Println(b,c)
	var d=true
	fmt.Println(d)
	var e int 
	fmt.Println(e)

	f :="apple"
	fmt.Println(f)
}
/*
var keyword in Go:
-------------------
- `var` is used to declare variables in Go.
- Syntax:
    
    var variableName dataType = value

- Example:
    
    var b int = 1

- Unlike C++, Go places the datatype AFTER the variable name.
- C++:
    
    int a = 2;

- Go:
    
    var a int = 2

------------------------------------------------------------

Go is Statically Typed:
-----------------------
- Go is a statically typed language.
- This means the datatype of variables is decided at compile time, NOT runtime.
- Example:
    
    var x int = 5

  Here `x` can only store integers.

- If you later do:
    
    x = "hello"

  Go compiler gives an error before program runs.

------------------------------------------------------------

Line-by-line Explanation:
-------------------------

var a = "initial"
-----------------
- Go automatically detects datatype using the assigned value.
- `"initial"` is a string, so Go infers:
    
    a is of type string

- This is called TYPE INFERENCE.

------------------------------------------------------------

var b, c int = 1, 2
-------------------
- Multiple variables can be declared together.
- Both `b` and `c` are integers.
- `int` applies to both variables.

Equivalent in C++:
    
    int b = 1, c = 2;

------------------------------------------------------------

var d = true
-------------
- Go infers datatype as bool.
- `true` and `false` are boolean values.

------------------------------------------------------------

var e int
----------
- Variable declared without assigning value.
- Go automatically gives DEFAULT ZERO VALUE.

For int:
    
    0

For string:
    
    ""

For bool:
    
    false

------------------------------------------------------------

f := "apple"
-------------
- `:=` is SHORT VARIABLE DECLARATION.
- Go automatically:
    1. creates variable
    2. infers datatype
    3. assigns value

Equivalent to:
    
    var f string = "apple"

- `:=` can only be used INSIDE functions.

------------------------------------------------------------

fmt.Println(...)
----------------
- `fmt` is Go's formatting package.
- `Println` prints values with newline.
- Similar to:
    
    cout << value << endl;

  in C++.

------------------------------------------------------------

Important Difference Between Go and C++:
----------------------------------------
C++:
    datatype comes first

    int a = 5;

Go:
    variable name comes first

    var a int = 5;

Reason:
- Go syntax is designed to be cleaner and more readable.
- It also supports easy type inference.

------------------------------------------------------------

Three Common Ways to Declare Variables in Go:
---------------------------------------------

1. Full declaration:
    
    var x int = 10

2. Type inference:
    
    var x = 10

3. Short declaration:
    
    x := 10

------------------------------------------------------------

Memory & Typing:
----------------
- Go variables are statically typed.
- Memory size/type is known during compilation.
- But Go still feels flexible because of type inference.

So:
- Static typing
- Compile-time checking
- Cleaner syntax
- Automatic type inference

all exist together in Go.
*/