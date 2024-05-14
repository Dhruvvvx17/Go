# GO

- Go is a compiled language.

- Go is a statically typed language - declare variables explicitly or the type should be inferred. These variables cannot be changed until type conversion.

    Example:

    - Explicit declaration: `var name string`

    - Type inferred: `var name = "Dhruv"` or `name := "Dhruv"`


- Go is also a strongly typed language. I.e., the operation you can perform on a variable depends on the variable type.

- Statically and Strongly typed language enables the Go compiler to thoroughly check code for errors.

- Go has built in concurrency - Goroutines

## Data Types

- Interger:

    - int, int8, int16, int32, int64
    - uint, uint8, uint16, uint32, uint64

- Float:

    - float32, float64

- Cannot perform operations between different types i.e., between a int and a float.

- Integer division returns an integer rounded down

- String:
    - "enter string here" or \`enter sting here\` (for exact formatting)
    - len(variable) -> length of the strings

- Runes:
    - single character in go within single quotes

- Boolean:
    - either true or false


## Arrays - Fixed Length List

- Array initialization: `var arr [3]int32`

    - 0 indexed initialization
    - initally all values would be 0 by default
    - python like slicing
    - &arr[0] is the memory location


## Slices - Variable Length List

- Slices as a data structure are wrappers around arrays

    - initalization: `var slice []int32`
    - a slice doesn't take a size
    - insert element at the end of the slice: `slice = append(slice, num)`
    - insert multiple elements at the end of the slice: `slice1 = append(slice1, slice2...)`
    - the `...` is called the spread operator and unpacks the value in a container
    - `len(slice)` -> to fetch the current length of the slice (active elements)
    - `cap(slice)` -> to fetch the current capacity of the slice (possible number of elements before capacity reallocation)
    - if `append()` is called and the length is equal to the capacity, the capacity is doubled and the new element is appended
 

 ## Maps

- Initalization: `var cache map[string]uint8 = make(map[string]uint8)`

- Initialization with values: `var cache = map[string]uint8{"name": dhruv, "age":24}`

- Accessing a map will always return something even if the key does not exist.
    - the default value is returned based on the datatype of the value. 0 for int.
    - a second value is returned when a map is accessed which is true if the key exists or false otherwise

- Deleting from a map:
    - `delete(cache, "age")` this does not return a new map

 ## String

 - string are immutable in go

 - strings are byte arrays of the UTF-8 character represention forming the string

 - indexing a string returns the UTF-8 value of the character at the index

 - len(string) results in the number of bytes the character uses

 - a forloop based string concatenation is inefficient as it creates a new string at each iteration

 -  instead use `string.Builder.WriteString()` from the string package in go. This appends each character to the existing builder and at the end a new string is created. This is much more efficient than the generic += concatenation.