package fuzzer

import (
	"bytes"
	"math/rand/v2"
	"strconv"
	"strings"
	"unicode"
)

func binaryEncode(input string) string {
	var result strings.Builder
	for _, b := range []byte(input) {
		result.WriteString(byteToBinaryString(b))
	}
	return result.String()
}

func byteToBinaryString(b byte) string {
	return strings.Join([]string{
		string('0' + (b>>7)&1),
		string('0' + (b>>6)&1),
		string('0' + (b>>5)&1),
		string('0' + (b>>4)&1),
		string('0' + (b>>3)&1),
		string('0' + (b>>2)&1),
		string('0' + (b>>1)&1),
		string('0' + b&1),
	}, "")
}

func randomInt(min, max int) int {
	n := max - min + 1
	if n <= 0 {
		return min
	}
	return rand.IntN(n) + min
}

func randomFloat(min, max float64) float64 {
	if min >= max {
		return min
	}
	return min + rand.Float64()*(max-min)
}

func floatFromString(str string, def float64) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return def
	}
	return i
}

func intFromString(str string, def int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return i
}

func randomStringFromList(strings ...string) string {
	if len(strings) <= 0 {
		return ""
	}
	return strings[randomInt(0, len(strings)-1)]
}

func randomString(chars string, length int) string {
	l := len(chars) - 1
	var b bytes.Buffer
	for i := 0; i < length; i++ {
		b.WriteByte(chars[randomInt(0, l)])
	}
	return b.String()
}

func removeNonPrintable(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
}
