package main
import (
	"fmt"
	"time"
)

func main(){
	i:=2
	fmt.Println("Write",i," as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
    
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}){
		switch t:= i.(type){
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Dont know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

/*
========================================================
PROGRAM OVERVIEW
========================================================

This Go program demonstrates:
1. switch statement
2. multiple case matching
3. switch without condition
4. anonymous functions
5. interface{}
6. type switch

========================================================
PART 1 : BASIC SWITCH
========================================================

i := 2

- Creates variable i
- Stores integer value 2

--------------------------------------------------------

fmt.Println("Write", i, "as")

Output:
    Write 2 as

--------------------------------------------------------

switch i {

- Go compares value of i with cases.

Equivalent logic:

if i == 1
else if i == 2
else if i == 3

--------------------------------------------------------

case 2:
    fmt.Println("two")

Since i = 2

Output:
    two

========================================================
PART 2 : SWITCH WITH CURRENT DAY
========================================================

switch time.Now().Weekday()

--------------------------------------------------------

time.Now()
-----------
- Gets current date and time.

Weekday()
----------
- Extracts current day.
- Returns:
    Monday
    Tuesday
    Wednesday
    etc.

--------------------------------------------------------

case time.Saturday, time.Sunday:

Means:
    if today is Saturday OR Sunday

Go allows multiple values in same case.

--------------------------------------------------------

default:
-----------
Runs when no case matches.

Equivalent to:
    else

========================================================
PART 3 : SWITCH WITHOUT VARIABLE
========================================================

t := time.Now()

Stores current time in variable t.

--------------------------------------------------------

switch {

This is special syntax.

No variable is given.

Internally Go treats it like:

switch true

--------------------------------------------------------

case t.Hour() < 12:

Hour()
-------
Returns current hour.

Example:
    9
    14
    22

--------------------------------------------------------

If hour is less than 12:
    "It's before noon"

Otherwise:
    "It's after noon"

Equivalent normal if-else:

if t.Hour() < 12 {
    ...
} else {
    ...
}

========================================================
MOST IMPORTANT PART
========================================================

whatAmI := func(i interface{}) {

========================================================
STEP 1 : ANONYMOUS FUNCTION
========================================================

func(i interface{}) {
}

This is a function WITHOUT a name.

Called anonymous function.

--------------------------------------------------------

Stored inside variable:

whatAmI :=

So now:
    whatAmI behaves like a function.

--------------------------------------------------------

Equivalent normal function:

func whatAmI(i interface{}) {

}

========================================================
STEP 2 : interface{}
========================================================

interface{}

means:

    "accept ANY datatype"

--------------------------------------------------------

So this parameter can store:
- int
- bool
- string
- float
- struct
- anything

--------------------------------------------------------

Examples:

whatAmI(true)
whatAmI(1)
whatAmI("hey")

All valid.

--------------------------------------------------------

In modern Go:

any

is alias for:

interface{}

========================================================
STEP 3 : TYPE SWITCH
========================================================

switch t := i.(type)

This is VERY IMPORTANT.

--------------------------------------------------------

Normally switch checks VALUES.

Example:

switch x {
case 1:
}

But here:

switch t := i.(type)

checks DATATYPE.

--------------------------------------------------------

(type)
-------
This special syntax asks:

    "What actual datatype is stored inside i?"

--------------------------------------------------------

This ONLY works in type switches.

========================================================
UNDERSTANDING EACH CALL
========================================================

--------------------------------------------------------
CALL 1
--------------------------------------------------------

whatAmI(true)

Parameter i contains:

    true

Datatype:

    bool

--------------------------------------------------------

switch checks datatype:

case bool:

MATCHED.

Output:
    I am bool

========================================================

--------------------------------------------------------
CALL 2
--------------------------------------------------------

whatAmI(1)

Datatype of 1:

    int

--------------------------------------------------------

case int:

MATCHED.

Output:
    I am an int

========================================================

--------------------------------------------------------
CALL 3
--------------------------------------------------------

whatAmI("hey")

Datatype:

    string

--------------------------------------------------------

No matching case exists for string.

So:

default:

runs.

========================================================
STEP 4 : %T
========================================================

fmt.Printf("Dont know type %T\n", t)

--------------------------------------------------------

%T
----
Prints datatype.

Example:

string
int
bool

--------------------------------------------------------

Output:
    Dont know type string

========================================================
VERY IMPORTANT DETAIL
========================================================

switch t := i.(type)

Here:
- i = interface{}
- t = actual value after type extraction

Example:

If:
    i contains "hey"

Then:
    t becomes string value "hey"

--------------------------------------------------------

So:
    %T prints datatype of t

========================================================
WHY TYPE SWITCHES ARE USEFUL
========================================================

Used when:
- function accepts many datatypes
- generic handling is needed
- working with unknown data
- JSON parsing
- APIs
- middleware
- logging systems

========================================================
FINAL OUTPUT EXAMPLE
========================================================

Write 2 as
two
It's a weekday
It's after noon
I am bool
I am an int
Dont know type string

========================================================
SHORT SUMMARY
========================================================

whatAmI:
---------
- anonymous function
- accepts any datatype using interface{}
- uses type switch
- checks actual datatype at runtime
- prints message based on datatype

This is Go's way of doing runtime type checking.
switch without variable or expression is treated as 
switch true
*/