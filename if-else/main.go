package main
import "fmt"

func main(){
	if 7%2 == 0{
		fmt.Println("7  is even")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4 == 0{
		fmt.Println("8 is divisible by 4")
	}
	if 8%2 == 0 || 7%2 ==0{
		fmt.Println("either 8 or 7 are even")
	}
	if num:=9 ; num<0{
      fmt.Println(num,"is negative")
	}else if num<10{
		fmt.Println(num,"has 1 digit")
	} else{
		fmt.Println(num,"has multiple digits")
	}
}

/*
========================================================
IF STATEMENT IN GO
========================================================

Basic Syntax:
--------------
if condition {
    // code
} else {
    // code
}

- Curly braces `{}` are compulsory in Go.
- Parentheses around condition are NOT required.

Example:
---------
if 7%2 == 0 {
    fmt.Println("even")
}

--------------------------------------------------------

7 % 2 == 0
-----------
- `%` is modulus operator.
- It gives remainder after division.

7 % 2
= 1

because:
    7 / 2 = 3 remainder 1

So:
    7 % 2 == 0

becomes:
    1 == 0

which is FALSE.

Therefore:
    else block executes.

--------------------------------------------------------

if 8%4 == 0
------------
8 % 4 = 0

Condition becomes:
    0 == 0

TRUE.

So:
    "8 is divisible by 4"

gets printed.

--------------------------------------------------------

LOGICAL OR (||)
----------------
if 8%2 == 0 || 7%2 == 0

- `||` means OR operator.
- If ANY ONE condition is true,
  whole condition becomes true.

8 % 2 == 0
TRUE

7 % 2 == 0
FALSE

TRUE || FALSE
= TRUE

So block executes.

========================================================
MOST IMPORTANT PART
========================================================

if num := 9; num < 0 {
    fmt.Println(num, "is negative")
}

This syntax is SPECIAL in Go.

--------------------------------------------------------
STRUCTURE
--------------------------------------------------------

if initialization; condition {
    // code
}

Go allows:
1. variable creation
2. condition checking

inside SAME if statement.

--------------------------------------------------------
BREAKING IT DOWN
--------------------------------------------------------

num := 9
---------
- Creates variable `num`
- Assigns value 9
- `:=` means short variable declaration.

Equivalent:
    var num int = 9

BUT:
- This variable exists ONLY inside this if-else block.
- Outside the block, `num` does not exist.

--------------------------------------------------------

num < 0
---------
Checks:
    Is 9 less than 0?

FALSE.

So first block does not run.

--------------------------------------------------------

else if num < 10
-----------------
Now Go checks:

    9 < 10

TRUE.

So this executes:

    fmt.Println(num, "has 1 digit")

Output:
    9 has 1 digit

========================================================
VERY IMPORTANT CONCEPT
========================================================

Scope of Variable
-----------------
`num` only exists inside:

if ... else if ... else

After block ends:

    num disappears.

Example:
---------

if num := 9; num < 0 {
    fmt.Println(num)
}

fmt.Println(num) // ERROR

Reason:
- num is local to the if statement.

========================================================
WHY GO PROVIDES THIS FEATURE
========================================================

This helps:
- keep variables temporary
- reduce unnecessary variables
- improve readability
- avoid polluting outer scope

Very commonly used with:
- error handling
- file operations
- database queries

Example:
---------

if err := someFunction(); err != nil {
    fmt.Println(err)
}

Here:
- err exists only for this check.

========================================================
EQUIVALENT NORMAL VERSION
========================================================

num := 9

if num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

Both work same.

But Go programmers often prefer:

    if initialization; condition

style because it keeps variable usage limited.
*/