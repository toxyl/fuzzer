# Fuzzer

`fuzzer` is a versatile string fuzzing library for Golang that allows the generation of random data based on tokens embedded in strings. It supports a range of token types for generating random strings, numbers, UUIDs, and custom user-defined functions. Additionally, it can load data from a directory into memory for use in fuzzing.

## Features

- **Randomized Data Generation**: Generate strings, integers, floats, and UUIDs with simple tokens.
- **Custom Functions**: Define and use custom functions for generating specialized data.
- **In-Memory Data Cache**: Load data from files into memory and access them during fuzzing.
- **Data Encoding**: Encode data as Binary, Base32, Base64, Base85, Hex, or URL-encoded strings.

## Installation

```bash
go get github.com/toxyl/fuzzer
```

## Usage

### Initialization

Before using the fuzzer, initialize it with an optional data directory and a map of custom functions:

```go
import (
    "github.com/toxyl/fuzzer"
)

func main() {
    userFns := map[string]func(args ...string) string{
        "customFn": func(args ...string) string {
            return "custom output"
        },
    }
    fuzzer.Init("data_dir", userFns)
}
```

### Fuzzing a String

After initialization, use the `Fuzz` function to generate a randomized output based on the provided string:

```go
output := fuzzer.Fuzz("[sl:10] [su:5] [i:4] [f:6.2] [$customFn:arg1;arg2]")
```

## Available Tokens

The following tokens can be used within strings passed to `Fuzz()`:

- **UUIDs and Hashes**
  - `[UUID]` or `[#UUID]` - Generates a random UUID (e.g., `xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxx`).
  - `[#56]` - Generates a random 56-character hash.

- **Random Numbers**
  - `[i:6]` - Random integer with a 6-character length (zero-padded).
  - `[f:6.2]` - Random float with a 6-character integer part and a 2-character fractional part.
  - `[10:500]` - Random integer between 10 and 500 (inclusive, can include negative values).
  - `[0.5:5.5]` - Random float between 0.5 and 5.5 (inclusive, can include negative values).

- **Random Strings**
  - `[sl:6]` - Random lowercase string with 6 characters (a-z).
  - `[su:6]` - Random uppercase string with 6 characters (A-Z).
  - `[s:6]` - Random mixed-case string with 6 characters (a-z, A-Z).
  - `[al:6]` - Random lowercase alphanumeric string with 6 characters (a-z, 0-9).
  - `[au:6]` - Random uppercase alphanumeric string with 6 characters (A-Z, 0-9).
  - `[a:6]` - Random mixed-case alphanumeric string with 6 characters (a-z, A-Z, 0-9).

- **Lists and Ranges**
  - `[10..500]` - Generates a comma-separated list of all integers from 10 to 500 (inclusive).
  - `[a,b,c]` - Randomly selects a value from the given list.

- **In-Memory Cache**
  - `[:path]` - Reads a random line from the specified path in the in-memory cache. If the path is a directory, a random file from that directory will be used.

- **User Functions**
  - `[$fn:arg1;arg2;...;argN]` - Executes a user-defined function `fn` with the specified arguments.

- **Data Encoding**
  - `[bin:data]` - Encodes `data` as binary.
  - `[b32:data]` - Encodes `data` as Base32.
  - `[b64:data]` - Encodes `data` as Base64.
  - `[b85:data]` - Encodes `data` as Base85.
  - `[hex:data]` - Encodes `data` as hexadecimal.
  - `[url:data]` - Encodes `data` as URL (percent-encoding).

### Custom Functions

Use custom functions for more complex data generation:

```go
userFns := map[string]func(args ...string) string{
    "greet": func(args ...string) string {
        return "Hello, " + strings.Join(args, " ")
    },
}

fuzzer.Init("", userFns)

output := fuzzer.Fuzz("[$greet:John;Doe]")
fmt.Println(output) // Output: Hello, John Doe
```

## Examples

Run the test suite for examples that illustrate the use of different tokens and custom functions:
```bash
go test -v ./...
```

## License

This project is released into the public domain under the UNLICENSE.
