package fuzzer

const (
	tknStart         = "["
	tknEnd           = "]"
	tknTypeStr       = "s"                  //
	tknTypeStrL      = tknTypeStr + "l:"    // [sl:5]
	tknTypeStrU      = tknTypeStr + "u:"    // [su:5]
	tknTypeStrR      = tknTypeStr + ":"     // [s:5]
	tknTypeAlpha     = "a"                  //
	tknTypeAlphaL    = tknTypeAlpha + "l:"  // [al:5]
	tknTypeAlphaU    = tknTypeAlpha + "u:"  // [au:5]
	tknTypeAlphaR    = tknTypeAlpha + ":"   // [a:5]
	tknTypeFloat     = "f:"                 // [f:5,2]
	tknTypeInt       = "i:"                 // [i:5]
	tknTypeHash      = "#"                  // [#5]
	tknTypeUUID      = tknTypeHash + "UUID" // [#UUID]
	tknTypeEncBase64 = "b64:"               // [b64:123]
	tknTypeEncBase32 = "b32:"               // [b32:123]
	tknTypeEncBase85 = "b85:"               // [b85:123]
	tknTypeEncBin    = "bin:"               // [bin:123]
	tknTypeEncURL    = "url:"               // [url:123]
	tknTypeEncHex    = "hex:"               // [hex:123]
	tknTypeMem       = ":"                  // [:dir/file]
	tknTypeFn        = "$"                  // [$fn:1234;5667;abc;3g]
)

const (
	charsetInt         = "0123456789"
	charsetStringLower = "abcdefghijklmnopqrstuvwxyz"
	charsetStringUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetHex         = "abcdef" + charsetInt
	charsetStringRand  = charsetStringLower + charsetStringUpper
	charsetAlphaLower  = charsetStringLower + charsetInt
	charsetAlphaUpper  = charsetStringUpper + charsetInt
	charsetAlphaRand   = charsetStringLower + charsetStringUpper + charsetInt
)
