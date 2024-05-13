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
