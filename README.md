# GO

- Go is a compiled language.

- Go is a statically typed language - declare variables explicitly or the type should be inferred. These variables cannot be changed until type conversion.

    Example:

    - Explicit declaration: `var name string`

    - Type inferred: `var name = "Dhruv"` or `name := "Dhruv"`


- Go is also a strongly typed language. I.e., the operation you can perform on a variable depends on the variable type.

- Statically and Strongly typed language enables the Go compiler to thoroughly check code for errors.

- Go has built in concurrency - Goroutines.

## Data Types

- Interger:

    - int, int8, int16, int32, int64
    - uint, uint8, uint16, uint32, uint64

- Float:

    - float32, float64

- Cannot perform operations between different types i.e., between a int and a float.

- Integer division returns an integer rounded down.

- String:
    - "enter string here" or \`enter sting here\` (for exact formatting).
    - len(variable) -> length of the strings.

- Runes:
    - Single character in go within single quotes.

- Boolean:
    - Either true or false.


## Arrays - Fixed Length List

- Array initialization: `var arr [3]int32`

    - 0 indexed initialization.
    - Initally all values would be 0 by default.
    - Python like slicing.
    - &arr[0] is the memory location.


## Slices - Variable Length List

- Slices as a data structure are wrappers around arrays

    - Initalization: `var slice []int32`.
    - A slice doesn't take a size.
    - Insert element at the end of the slice: `slice = append(slice, num)`.
    - Insert multiple elements at the end of the slice: `slice1 = append(slice1, slice2...)`.
    - The `...` is called the spread operator and unpacks the value in a container.
    - `len(slice)` -> to fetch the current length of the slice (active elements).
    - `cap(slice)` -> to fetch the current capacity of the slice (possible number of elements before capacity reallocation).
    - Ff `append()` is called and the length is equal to the capacity, the capacity is doubled and the new element is appended.
 

## Maps

- Initalization: `var cache map[string]uint8 = make(map[string]uint8)`.

- Initialization with values: `var cache = map[string]uint8{"name": dhruv, "age":24}`.

- Accessing a map will always return something even if the key does not exist.
    - The default value is returned based on the datatype of the value. 0 for int.
    - A second value is returned when a map is accessed which is true if the key exists or false otherwise.

- Deleting from a map:
    - `delete(cache, "age")` this does not return a new map.

## String

 - String are immutable in go.

 - Strings are byte arrays of the UTF-8 character represention forming the string.

 - Indexing a string returns the UTF-8 value of the character at the index.

 - len(string) results in the number of bytes the character uses.

 - A forloop based string concatenation is inefficient as it creates a new string at each iteration.

 - Instead use `string.Builder.WriteString()` from the string package in go. This appends each character to the existing builder and at the end a new string is created. This is much more efficient than the generic += concatenation.

## Pointers

- Same as pointers in C/C++. `var *p int32` is a pointer to a memory address which stores an `int32`.

- Technically in go, the above statement initally points to `nil` and has to call the `new()` method to point to a real location.

- `var *p int32 = new(int32)` now p points to a real location and the value at that location can be set by `*p = 5`.

- For p to point to the location of an existing variable, use `p = &i`.

- By passing an array to a function using its memory address `&` and then referencing that with a `*p` in the function allows operating right over the original array and save memory by not making duplicates for every function call.

## Go Routines

- GO routines are a way to launch multiple functions and have them execute concurrently.

- Concurrency is not the same as parallelism. Concurrency is context switching between the CPU so it doesn't stay idle while waiting on some action. The base CPU is still a single resource.

- Parallelism is when multiple CPUs are executing multiple tasks in parallel.

- To introduce concurrency in the code, identify the statement that can be executed concurrently (eg: waiting for a DB query response) and prefix it with the keyword `go`.

- Next, `import sync` and create a wait group `var wg = sync.WaitGroup{}`. Write before calling the GO routine, add the statement `wg.Add(1)` so the program keeps track of a pending routine.

- In the concurrent method itself, use `wg.Done()` at the end to mark that the current GO routine has finished execution. 

- To prevent the program flow from exiting while pending routines exists, use `wg.Wait()`. This checks the number of routines pending (as we increment each time calling `wg.Add(1)`) and proceed only after all routines are completed.

## Mutex

- When using GO routines, if different routines are trying to modify the same data structure, it might lead to some inconsistency in the results.

- Mutex are a way to prevent this. Mutex are locks that can guard a database read/write statement so that if multiple GO routines are to access it, only one can secure the lock and proceed while others have to wait until the lock is unlocked.

- `import sync` then create `var m = sync.Mutex{}`.

- Now use `m.Lock()` and `m.Unlock()` around the statement that is to be protected.

- Another type of lock offered is the read-write mutex. `var m = sync.RWMutex{}` this enables routines to access a shared variable for read access but not for any writes until a different routine completes the write and unlocking the mutex.

## Channels

- A way to enable GO routines to pass around information. The main features are:

    - They hold data - int, string, slice, etc.

    - They are thread safe - avoid data races when reading and writing from memory.

    - The code can listen for data addition/removal to/from a channel and execute code accrodingly.

- Creating a channel: `var c = make(chan int)`.

- Adding a data into channel: `c <- 5`.

- Reading from a channel: `var x <- c`.

- Creating a channel buffer: `var c = make(chan int, 5)`.

## Generics

- A way for GO to implement reusing code for varying datatypes.

- To overload a function such that it can accept different datatypes:

    - `func sum[T int | float32](slice []T) T {...}`.
    - Here, the function can accept a slice of either int or float and accordingly return either a int or a float.