package fuzzer

import (
	"strings"
)

// Init initializes the fuzzer. If the given data dir not an empty string
// it will recursively load the files from the given directory into memory.
// You can supply additional functions to the fuzzer using the userFns param.
// These can then be used with the `[$fn:arg1;arg2;...;argN]` syntax.
func Init(dataDir string, userFns map[string]func(args ...string) string) {
	dataDir = strings.TrimSpace(dataDir)
	if dataDir != "" {
		importFromDir(dataDir)
	}
	fz = &fuzzer{
		fns: userFns,
	}
}

// Fuzz takes the input string and replaces all tokens
// with randomly generated data. Tokens can be nested.
//
// Available tokens:
//
//	[#UUID]    = random UUID (xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxx)
//	[#56]      = random 56-characters hash
//	[f:6.2]    = random float with a 6-character integer part and a 2-character fractional part (zero-padded)
//	[i:6]      = random 6-characters integer (zero-padded)
//	[sl:6]     = random 6-characters lowercase string (a-z)
//	[su:6]     = random 6-characters uppercase string (A-Z)
//	[s:6]      = random 6-characters mixed-case string (a-z, A-Z)
//	[al:6]     = random 6-characters lowercase alphanumeric string (a-z, 0-9)
//	[au:6]     = random 6-characters uppercase alphanumeric string (A-Z, 0-9)
//	[a:6]      = random 6-characters mixed-case alphanumeric string (a-z, A-Z, 0-9)
//	[0.5:5.5]  = random float64 value between 0.5 and 5.5 (inclusive, values can be negative)
//	[10:500]   = random value between 10 and 500 (inclusive, values can be negative)
//	[10..500]  = comma separated list with all ints from 10 to 500 (inclusive, values can be negative)
//	[a,b,c]    = random value from the list
//
// In-memory cache:
//
//	[:path]
//
// Reads a random line from the given "path" in the memory cache.
// If path is a directory a random file from that directory will be used.
//
// User functions:
//
//	[$fn:arg1;arg2;...;argN]
//
// Executes the user-supplied function "fn" with the arguments "arg1" through "argN"
// (ie. fn(arg1,arg2,...,argN)) which must return a new string.
//
// Data encoding:
//
//	[bin:data] = encodes 'data' as Binary
//	[b32:data] = encodes 'data' as Base32
//	[b64:data] = encodes 'data' as Base64
//	[b85:data] = encodes 'data' as Base85
//	[hex:data] = encodes 'data' as Hex
//	[url:data] = encodes 'data' as URL (percent-encoding)
func Fuzz(str string) string {
	if fz == nil {
		panic("You must first initialize the fuzzer using the Init() function!")
	}
	return fz.process(str)
}
