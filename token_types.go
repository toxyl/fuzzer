package fuzzer

type tokenTypeLen struct {
	Length int
}

type tokenTypeDualLen struct {
	LengthA int
	LengthB int
}

type tokenTypeStr struct {
	Data string
}

type tokenTypeStrs struct {
	Strings []string
}

type tokenTypeRange struct {
	Min int
	Max int
}

type tokenTypeRangeFloat struct {
	Min float64
	Max float64
}

type tokenTypeFn struct {
	Data string
	Fn   func(args ...string) string
}

type tokenTypeBasic string

type (
	tAlphaLower  tokenTypeLen
	tAlphaUpper  tokenTypeLen
	tAlpha       tokenTypeLen
	tHash        tokenTypeLen
	tFloat       tokenTypeDualLen
	tInt         tokenTypeLen
	tStrLower    tokenTypeLen
	tStrUpper    tokenTypeLen
	tStr         tokenTypeLen
	tIntList     tokenTypeRange
	tIntRange    tokenTypeRange
	tFloatRange  tokenTypeRangeFloat
	tMemory      tokenTypeStr
	tStrFromList tokenTypeStrs
	tUserFn      tokenTypeFn
	tUUID        tokenTypeBasic
	tAny         tokenTypeBasic
	tUnknown     tokenTypeBasic
)
