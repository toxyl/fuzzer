package fuzzer

import (
	"fmt"
	"strings"
)

type fuzzer struct {
	fns map[string]func(args ...string) string // contains user defined functions, during parsing the tag contents will be split on semicolons to generate the function args
}

func (f *fuzzer) process(str string) string {
	if str == "" {
		return str
	}

	tokens := f.tokenize(str)
	var parsed strings.Builder

	for _, t := range tokens {
		var result string
		switch tt := t.(type) {
		case tUserFn:
			result = f.process(tt.parse())
		case tMemory:
			result = f.process(tt.parse())
		case tFloat:
			result = tt.parse()
		case tFloatRange:
			result = tt.parse()
		case tInt:
			result = tt.parse()
		case tIntRange:
			result = tt.parse()
		case tIntList:
			result = tt.parse()
		case tHash:
			result = tt.parse()
		case tUUID:
			result = tt.parse()
		case tStrFromList:
			result = tt.parse()
		case tStrLower:
			result = tt.parse()
		case tStrUpper:
			result = tt.parse()
		case tStr:
			result = tt.parse()
		case tAlphaLower:
			result = tt.parse()
		case tAlphaUpper:
			result = tt.parse()
		case tAlpha:
			result = tt.parse()
		case tAny:
			result = fmt.Sprint(tt)
			for strings.Count(result, tknStart) > 0 {
				if strings.Count(result, tknStart) != strings.Count(result, tknEnd) {
					result = "" // Don't include unparsed tokens
					break
				}
				if strings.Count(result, tknStart) > 1 {
					// Solve nested tokens first
					result = "[" + f.process(result[1:len(result)-1]) + "]"
					continue
				}
				// Process the outermost token
				result = f.process(result)
			}
		case tUnknown:
			result = "" // Don't include unparsed tokens
		default:
			result = fmt.Sprint(tt)
		}
		parsed.WriteString(result)
	}

	return strings.Trim(parsed.String(), " \r\n")
}
