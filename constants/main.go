package main

import(
	"fmt"
	"math"
)

const s string = "constant"

func main(){

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))


	fmt.Println(math.Sin(n))
}

/*

========================
GO CONSTANTS EXPLANATION
========================

1) package main
----------------
Every Go program starts with a package declaration.

"main" is a special package.
If the package name is main, Go understands that this file
can be executed as a program.

Execution starts from the main() function.


2) import
----------
import is used to include external packages/modules.

"fmt"
- Used for formatted input/output.
- fmt.Println() prints data to the console.

"math"
- Provides mathematical functions.
- Example:
    math.Sin()
    math.Sqrt()
    math.Pow()


3) const keyword
----------------
const is used to declare CONSTANT values.

Syntax:
    const name datatype = value

Example:
    const s string = "constant"

A constant value:
- cannot be changed later
- is fixed during program execution


4) const s string = "constant"
------------------------------
Here:

s       -> variable/identifier name
string  -> datatype
constant -> actual value

This creates a string constant.


5) func main()
--------------
main() is the entry point of the program.

Code inside main() executes when the program runs.


6) fmt.Println(s)
-----------------
Prints the value stored in s.

Output:
    constant


7) const n = 500000000
----------------------
Creates a numeric constant named n.

Go automatically infers the type when possible.

Since it is a constant,
Go keeps high precision internally.


8) const d = 3e20 / n
---------------------
3e20 means:

3 × 10^20

Scientific notation is supported in Go.

So:
    3e20 = 300000000000000000000

Then division happens:
    3e20 / 500000000

Result is stored in constant d.


9) Constant Arithmetic
----------------------
Go constants support arbitrary precision arithmetic.

Meaning:
- calculations are very accurate
- precision is maintained until a concrete type is needed


10) fmt.Println(d)
------------------
Prints the value of d.

Go automatically formats numeric output.


11) int64(d)
------------
This is TYPE CONVERSION.

Syntax:
    datatype(value)

Here:
    int64(d)

Converts d into int64 datatype.

Why needed?
Because d originally has no fixed datatype.


12) Untyped Constants
---------------------
In Go, numeric constants are often UNTYPED.

Example:
    const x = 10

Go decides the datatype only when required.

This gives flexibility and precision.


13) fmt.Println(math)
---------------------
This line is INVALID in actual Go programs.

Reason:
- math is a package
- packages cannot be printed directly

You normally use:
    math.Sin()
    math.Sqrt()

instead of printing the package itself.


14) math.Sin(n)
---------------
math.Sin() calculates sine value.

Syntax:
    math.Sin(number)

math.Sin expects:
    float64 datatype

So Go automatically treats n as float64 here.

This demonstrates:
- context-based type conversion for constants


15) Why constants are useful
----------------------------
Constants are used for:
- fixed configuration values
- mathematical values
- PI values
- application limits
- values that should never change

Example:
    const PI = 3.14159
    const MAX_USERS = 100


16) Difference between var and const
------------------------------------

var:
- value can change
- stored in memory normally

Example:
    var x = 10
    x = 20

const:
- value cannot change
- fixed permanently

Example:
    const y = 10


17) Important Rule
------------------
Only values known at compile time can be constants.

Allowed:
- numbers
- strings
- booleans
- constant expressions

Not allowed:
- runtime values
- user input
- function return values


*/