package fuzzer

import (
	"fmt"
	"strings"
)

func (t *tAlphaLower) parse() string {
	return randomString(charsetAlphaLower, t.Length)
}

func (t *tAlphaUpper) parse() string {
	return randomString(charsetAlphaUpper, t.Length)
}

func (t *tAlpha) parse() string {
	return randomString(charsetAlphaRand, t.Length)
}

func (t *tMemory) parse() string {
	return getRandomLine(t.Data)
}

func (t *tHash) parse() string {
	return randomString(charsetHex, t.Length)
}

func (t *tIntList) parse() string {
	a := make([]string, t.Max-t.Min+1)
	for i := range a {
		a[i] = fmt.Sprint(t.Min + i)
	}
	return strings.Join(a, ",")
}

func (t *tIntRange) parse() string {
	return fmt.Sprint(randomInt(t.Min, t.Max))
}

func (t *tInt) parse() string {
	return randomString(charsetInt, t.Length)
}

func (t *tFloat) parse() string {
	return randomString(charsetInt, t.LengthA) + "." + randomString(charsetInt, t.LengthB)
}

func (t *tFloatRange) parse() string {
	return fmt.Sprint(randomFloat(t.Min, t.Max))
}

func (t *tStrFromList) parse() string {
	return randomStringFromList(t.Strings...)
}

func (t *tStrLower) parse() string {
	return randomString(charsetStringLower, t.Length)
}

func (t *tStrUpper) parse() string {
	return randomString(charsetStringUpper, t.Length)
}

func (t *tStr) parse() string {
	return randomString(charsetStringRand, t.Length)
}

func (t *tUserFn) parse() string {
	return t.Fn(strings.Split(t.Data, ";")...)
}

func (t *tUUID) parse() string {
	return randomString(charsetHex, 8) +
		"-" + randomString(charsetHex, 4) +
		"-" + randomString(charsetHex, 4) +
		"-" + randomString(charsetHex, 12)
}
