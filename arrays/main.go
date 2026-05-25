package main
import "fmt"

func main(){
	var a [5]int
	fmt.Println("emp:",a)
	
	a[4]=100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	fmt.Println("len",len(a))

	b:= [5]int{1,2,3,4,5}
	fmt.Println("dcl:",b)
    
	b=[...]int{1,2,3,4,5}
	fmt.Println("dcl:",b)

	b=[...]int{100,3:400,500}
	fmt.Println("idx:",b)

	var twoD [2][3]int
	for i :=range 2{
		for j:=range 3{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d:",twoD)
	twoD= [2][3] int{
		{1,2,3},
		{1,2,3},
	}
	fmt.Println("2d: ",twoD)

}

/*
========================================================
ARRAYS IN GO
========================================================

Array:
------
- Collection of fixed-size elements.
- All elements must be same datatype.
- Size is part of datatype in Go.

Example:
    [5]int

means:
    array of 5 integers

========================================================
DECLARING ARRAY
========================================================

var a [5]int

--------------------------------------------------------

Breakdown:
-----------
var
    keyword for variable declaration

a
    variable name

[5]
    array size

int
    datatype of elements

--------------------------------------------------------

This creates array:

index:   0  1  2  3  4
value:   0  0  0  0  0

--------------------------------------------------------

Go automatically initializes arrays
with ZERO VALUES.

For int:
    0

For bool:
    false

For string:
    ""

========================================================

fmt.Println("emp:", a)

Output:
    emp: [0 0 0 0 0]

========================================================
SETTING VALUE
========================================================

a[4] = 100

--------------------------------------------------------

Array indexing starts from 0.

Indexes:

0 1 2 3 4

--------------------------------------------------------

So:
    a[4]

means LAST element.

Now array becomes:

[0 0 0 0 100]

========================================================

fmt.Println("get:", a[4])

Accesses value at index 4.

Output:
    100

========================================================
len()
========================================================

len(a)

Returns number of elements.

Output:
    5

--------------------------------------------------------

IMPORTANT:
------------
Length of array is FIXED.

You cannot increase/decrease it later.

========================================================
ARRAY INITIALIZATION
========================================================

b := [5]int{1,2,3,4,5}

--------------------------------------------------------

Creates array directly with values.

Array:

[1 2 3 4 5]

--------------------------------------------------------

:= means:
- declare variable
- infer datatype

========================================================
AUTO SIZE DETECTION
========================================================

b = [...]int{1,2,3,4,5}

--------------------------------------------------------

`...`

means:

    "Go, count the size yourself"

Go sees 5 elements.

So internally:

    [5]int

========================================================
INDEXED INITIALIZATION
========================================================

b = [...]int{100, 3:400, 500}

VERY IMPORTANT.

--------------------------------------------------------

Normal elements:
----------------
100

goes to index 0

--------------------------------------------------------

3:400

means:
    put 400 at index 3

--------------------------------------------------------

500

comes after index 3,
so it goes to index 4.

--------------------------------------------------------

Final array:

index:  0   1   2   3    4
value: [100 0   0 400 500]

--------------------------------------------------------

Output:
    idx: [100 0 0 400 500]

========================================================
2D ARRAYS
========================================================

var twoD [2][3]int

--------------------------------------------------------

Meaning:
---------
2 rows
3 columns

Like matrix:

[
  [0 0 0]
  [0 0 0]
]

========================================================
LOOPS
========================================================

for i := range 2 {

--------------------------------------------------------

NEW GO FEATURE:
----------------
range 2

means:
    iterate from 0 to 1

Equivalent:

for i := 0; i < 2; i++

--------------------------------------------------------

Similarly:

range 3

means:
    0 to 2

========================================================
FILLING 2D ARRAY
========================================================

twoD[i][j] = i + j

--------------------------------------------------------

Values become:

i=0:
    j=0 → 0+0 = 0
    j=1 → 0+1 = 1
    j=2 → 0+2 = 2

First row:
    [0 1 2]

--------------------------------------------------------

i=1:
    j=0 → 1+0 = 1
    j=1 → 1+1 = 2
    j=2 → 1+2 = 3

Second row:
    [1 2 3]

--------------------------------------------------------

Final matrix:

[
 [0 1 2]
 [1 2 3]
]

========================================================

fmt.Println("2d:", twoD)

Output:
    2d: [[0 1 2] [1 2 3]]

========================================================
DIRECT 2D INITIALIZATION
========================================================

twoD = [2][3]int{
    {1,2,3},
    {1,2,3},
}

--------------------------------------------------------

Creates matrix directly.

Row 1:
    [1 2 3]

Row 2:
    [1 2 3]

========================================================
IMPORTANT GO ARRAY CONCEPT
========================================================

In Go:

    [5]int

and

    [6]int

are DIFFERENT DATATYPES.

--------------------------------------------------------

Because size is part of type.

========================================================
ARRAY VS C++
========================================================

C++:
------
int arr[5];

Go:
-----
var arr [5]int

--------------------------------------------------------

Difference:
------------
Go puts datatype AFTER variable name.

========================================================
ARRAY MEMORY
========================================================

Arrays in Go:
-------------
- fixed size
- contiguous memory
- statically typed
- copied by value

--------------------------------------------------------

If you do:

b := a

Entire array gets copied.

Not reference.

========================================================
FINAL OUTPUT
========================================================

emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d: [[0 1 2] [1 2 3]]
2d: [[1 2 3] [1 2 3]]

========================================================
SHORT SUMMARY
========================================================

[5]int
    array of 5 integers

a[4]
    access index 4

len(a)
    array size

[...]int
    auto-detect size

[2][3]int
    2D array

range 2
    loop from 0 to 1

Arrays in Go:
- fixed size
- typed
- value copied
- size is part of datatype
*/