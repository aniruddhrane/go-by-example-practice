package main 
import(
	"fmt"
	"maps"
)

func main(){
	m:=make(map[string]int)

	m["k1"]=7
	m["k2"]=13

	fmt.Println("map",m)
	v1:=m["k1"]
	fmt.Println("v1:",v1)

	v3:=m["k3"]
	fmt.Println("v3",v3)

	fmt.Println("len",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	clear(m)
	fmt.Println("map",m)

	_, prs:=m["k2"]
	fmt.Println("prs",prs)

	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map",n)

	n2:=map[string]int{"foo":1,"bar":2}
	if maps.Equal(n,n2){
		fmt.Println("n==n2")
	}
}

/*
========================================================
MAPS IN GO
========================================================

Map:
----
A map stores data in:

    key -> value

format.

Similar to:
- dictionary in Python
- unordered_map in C++
- hash map in Java

========================================================
CREATING A MAP
========================================================

m := make(map[string]int)

--------------------------------------------------------

Breakdown:
-----------

map[string]int

means:

KEY TYPE:
    string

VALUE TYPE:
    int

--------------------------------------------------------

Examples:

"k1" -> 7
"k2" -> 13

--------------------------------------------------------

make()
-------
Allocates and initializes the map.

Without make:

var m map[string]int

map would be nil.

You cannot insert into nil map.

========================================================
ADDING VALUES
========================================================

m["k1"] = 7

--------------------------------------------------------

Syntax:
--------
map[key] = value

--------------------------------------------------------

Key:
    "k1"

Value:
    7

--------------------------------------------------------

Now map contains:

{
    "k1": 7
}

========================================================

m["k2"] = 13

Now map becomes:

{
    "k1": 7,
    "k2": 13
}

========================================================

fmt.Println("map", m)

Possible output:

map map[k1:7 k2:13]

--------------------------------------------------------

IMPORTANT:
------------
Maps are unordered.

Output order may change.

========================================================
ACCESSING VALUES
========================================================

v1 := m["k1"]

--------------------------------------------------------

Gets value associated with key "k1".

Result:
    7

========================================================

v3 := m["k3"]

VERY IMPORTANT.

--------------------------------------------------------

"k3" does NOT exist.

Go returns ZERO VALUE of value type.

Value type:
    int

Zero value for int:
    0

--------------------------------------------------------

So:

v3 = 0

========================================================
len()
========================================================

len(m)

Returns number of key-value pairs.

Currently:
    2

========================================================
delete()
========================================================

delete(m, "k2")

--------------------------------------------------------

Syntax:
--------
delete(map, key)

--------------------------------------------------------

Arguments:
-----------
1. map
2. key to remove

--------------------------------------------------------

After delete:

{
    "k1": 7
}

========================================================
clear()
========================================================

clear(m)

--------------------------------------------------------

Removes ALL elements from map.

Map becomes empty.

--------------------------------------------------------

Result:
---------
{}

========================================================
CHECKING IF KEY EXISTS
========================================================

_, prs := m["k2"]

VERY IMPORTANT CONCEPT.

--------------------------------------------------------

Map lookup actually returns TWO values.

Syntax:
--------
value, exists := map[key]

--------------------------------------------------------

1st value:
------------
actual value

2nd value:
------------
whether key exists

--------------------------------------------------------

Example:

v, ok := m["k1"]

If key exists:
    ok = true

Otherwise:
    ok = false

========================================================

_ , prs := m["k2"]

--------------------------------------------------------

_
---
blank identifier

Means:
    ignore first returned value

--------------------------------------------------------

prs
----
stores whether key exists.

Since map was cleared:

"k2" does not exist.

So:

prs = false

========================================================
DIRECT MAP INITIALIZATION
========================================================

n := map[string]int{
    "foo": 1,
    "bar": 2,
}

--------------------------------------------------------

Creates map directly with values.

========================================================
maps.Equal()
========================================================

Package:
---------
"maps"

Provides utility functions for maps.

========================================================

maps.Equal(n, n2)

--------------------------------------------------------

Checks whether two maps are identical.

--------------------------------------------------------

Syntax:
--------
maps.Equal(map1, map2)

--------------------------------------------------------

Arguments:
-----------
1. first map
2. second map

--------------------------------------------------------

Checks:
---------
- same keys
- same values

--------------------------------------------------------

Returns:
---------
true or false

========================================================

n:
{
    "foo":1,
    "bar":2
}

n2:
{
    "foo":1,
    "bar":2
}

--------------------------------------------------------

Result:
---------
true

========================================================
IMPORTANT MAP CHARACTERISTICS
========================================================

1. Maps are unordered
-----------------------
No guaranteed order.

========================================================

2. Keys must be comparable
---------------------------
Allowed:
- string
- int
- bool

Not allowed:
- slices
- maps
- functions

========================================================

3. Maps are reference types
----------------------------
When passed/copied,
underlying data is shared.

Unlike arrays.

========================================================

4. Reading missing key is safe
-------------------------------
Returns zero value.

========================================================

5. Writing to nil map causes panic
-----------------------------------

Example:

var m map[string]int

m["a"] = 1

ERROR.

Need make():

m := make(map[string]int)

========================================================
FUNCTION SUMMARY
========================================================

make(map[keyType]valueType)
----------------------------
Creates map.

--------------------------------------------------------

map[key] = value
----------------
Insert/update value.

--------------------------------------------------------

value := map[key]
------------------
Get value.

--------------------------------------------------------

delete(map, key)
-----------------
Remove key.

--------------------------------------------------------

clear(map)
-----------
Remove all entries.

--------------------------------------------------------

len(map)
---------
Number of entries.

--------------------------------------------------------

value, ok := map[key]
----------------------
Check if key exists.

--------------------------------------------------------

maps.Equal(a, b)
-----------------
Compare maps.

========================================================
FINAL OUTPUT
========================================================

map map[k1:7 k2:13]
v1: 7
v3 0
len 2
map: map[k1:7]
map map[]
prs false
map map[bar:2 foo:1]
n==n2

========================================================
VERY IMPORTANT DIFFERENCE
========================================================

ARRAY:
-------
index-based

SLICE:
-------
dynamic index-based

MAP:
----
key-based

--------------------------------------------------------

Examples:

Array/Slice:
    arr[0]

Map:
    m["name"]

========================================================
*/