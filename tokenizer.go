package fuzzer

import (
	"bufio"
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strings"
)

func (f *fuzzer) splitFn(data []byte, eof bool) (int, []byte, error) {
	inToken := false
	returnToken := false
	returnTokenData := false
	offset := 0
	index := 0
	depth := 0
	var char byte
	tokenOpen := tknStart[0]
	tokenClose := tknEnd[0]
	for index, char = range data {
		if returnTokenData {
			return index, data[offset:index], nil // return the token data
		}

		if char == tokenOpen {
			if !inToken {
				inToken = true
				offset = index
				if index > 0 {
					returnToken = true
				}
			}
			depth++
		} else if char == tokenClose {
			if inToken && depth == 1 {
				returnTokenData = true
			}
			depth--
		}

		if returnToken {
			return index, data[:index], nil // return the token
		}
	}

	if index > 0 {
		if inToken {
			v := data[:index+1]
			return index, v, nil // return the token
		}

		return index, data[:index], nil // we get here if the string does not end in a token
	}

	return 0, nil, nil // there was no data
}

func (f *fuzzer) tokenize(input string) []interface{} {
	input = removeNonPrintable(input)
	if input[len(input)-1:] != " " {
		input += " " // fixes a case where the end of the string is not correctly returned
	}
	buf := make([]byte, len(input))
	s := bufio.NewScanner(strings.NewReader(input))
	s.Buffer(buf, len(input))

	s.Split(f.splitFn)
	tokens := []interface{}{}
	for s.Scan() {
		t := s.Text()
		l := len(t) - 1
		iv := t[:1] == tknStart && t[l:] == tknEnd && strings.Count(t, tknStart) == 1
		if iv {
			if l >= 2 && t[1:2] == tknTypeStr {
				if t[1:4] == tknTypeStrL {
					tokens = append(tokens, tStrLower{Length: intFromString(t[4:l], 0)})
					continue
				}

				if t[1:4] == tknTypeStrU {
					tokens = append(tokens, tStrUpper{Length: intFromString(t[4:l], 0)})
					continue
				}

				if t[1:3] == tknTypeStrR {
					tokens = append(tokens, tStr{Length: intFromString(t[3:l], 0)})
					continue
				}
			}

			if l >= 2 && t[1:2] == tknTypeAlpha {
				if t[1:4] == tknTypeAlphaL {
					tokens = append(tokens, tAlphaLower{Length: intFromString(t[4:l], 0)})
					continue
				}

				if t[1:4] == tknTypeAlphaU {
					tokens = append(tokens, tAlphaUpper{Length: intFromString(t[4:l], 0)})
					continue
				}

				if t[1:3] == tknTypeAlphaR {
					tokens = append(tokens, tAlpha{Length: intFromString(t[3:l], 0)})
					continue
				}
			}

			if l >= 2 && t[1:2] == tknTypeHash {
				if l >= 6 && t[1:6] == tknTypeUUID {
					tokens = append(tokens, tUUID(""))
					continue
				}

				if reRandHash.MatchString(t) {
					tokens = append(tokens, tHash{Length: intFromString(t[2:l], 0)})
					continue
				}
			}

			if l >= 4 && t[1:3] == tknTypeInt {
				tokens = append(tokens, tInt{Length: intFromString(t[3:l], 0)})
				continue
			}

			if l >= 4 && t[1:3] == tknTypeFloat {
				e := strings.SplitN(t[3:l], ".", 2)
				tokens = append(tokens, tFloat{LengthA: intFromString(e[0], 0), LengthB: intFromString(e[1], 0)})
				continue
			}

			if strings.ContainsAny(t, "1234567890.-") {
				if reRandInt.MatchString(t) {
					v := strings.SplitN(t[1:l], ":", 2)
					tokens = append(tokens, tIntRange{Min: intFromString(v[0], 0), Max: intFromString(v[1], 0)})
					continue
				}

				if reIntRange.MatchString(t) {
					v := strings.SplitN(t[1:l], "..", 2)
					tokens = append(tokens, tIntList{Min: intFromString(v[0], 0), Max: intFromString(v[1], 0)})
					continue
				}

				if reRandFloat.MatchString(t) {
					v := strings.SplitN(t[1:l], ":", 2)
					tokens = append(tokens, tFloatRange{Min: floatFromString(v[0], 0), Max: floatFromString(v[1], 0)})
					continue
				}

			}

			if t[1:2] == tknTypeMem {
				tokens = append(tokens, tMemory{Data: t[2:l]})
				continue
			}

			if t[1:2] == tknTypeFn {
				e := strings.SplitN(t[2:l], ":", 2)
				name, data := e[0], e[1]
				if fn, ok := f.fns[name]; ok {
					tokens = append(tokens, tUserFn{Data: data, Fn: fn})
					continue
				}

				tokens = append(tokens, tUnknown(t))
				continue
			}

			if l >= 6 {
				if t[1:5] == tknTypeEncBase64 {
					tokens = append(tokens, tAny(base64.URLEncoding.EncodeToString([]byte(t[5:l]))))
					continue
				}

				if t[1:5] == tknTypeEncURL {
					tokens = append(tokens, tAny(url.QueryEscape(t[5:l])))
					continue
				}

				if t[1:5] == tknTypeEncHex {
					hexEncoded := hex.EncodeToString([]byte(t[5:l]))
					tokens = append(tokens, tAny(hexEncoded))
				}

				if t[1:5] == tknTypeEncBase32 {
					tokens = append(tokens, tAny(base32.StdEncoding.EncodeToString([]byte(t[5:l]))))
					continue
				}

				if t[1:5] == tknTypeEncBin {
					tokens = append(tokens, tAny(binaryEncode(t[5:l])))
					continue
				}

				if t[1:5] == tknTypeEncBase85 {
					base85Encoded := make([]byte, ascii85.MaxEncodedLen(len(t[5:l])))
					base85Len := ascii85.Encode(base85Encoded, []byte(t[5:l]))
					tokens = append(tokens, tAny(string(base85Encoded[:base85Len])))
					continue
				}

			}

			if reRandStr.MatchString(t) {
				tokens = append(tokens, tStrFromList{Strings: strings.Split(t[1:l], ",")})
				continue
			}

			tokens = append(tokens, tUnknown(t))
			continue
		}
		tokens = append(tokens, tAny(t))
	}
	return tokens
}
