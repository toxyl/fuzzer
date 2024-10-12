package fuzzer

import (
	"strings"
	"testing"
	"time"
)

func TestFuzz(t *testing.T) {
	Init("testdata", map[string]func(args ...string) string{
		"foo":  func(args ...string) string { return strings.Join(args, " ") },
		"time": func(args ...string) string { return time.Now().Format(time.RFC3339Nano) },
	})
	tests := []struct {
		name string
		str  string
	}{
		{"str (lower)", "[sl:10]"},
		{"str (upper)", "[su:10]"},
		{"str (mix)", "[s:10]"},
		{"alpha (lower)", "[al:10]"},
		{"alpha (upper)", "[au:10]"},
		{"alpha (mix)", "[a:10]"},
		{"hex", "[#10]"},
		{"uuid", "[#UUID]"},
		{"range", "[1:5]"},
		{"range", "[-10:10]"},
		{"range", "[0.4:5.4]"},
		{"range", "[-5.0:5.0]"},
		{"list", "[1..5]"},
		{"list", "[-2..2]"},
		{"int", "[i:10]"},
		{"float", "[f:5.2]"},
		{"enc b32", "[b32:hello world & foo bar [0-10]]"},
		{"enc b64", "[b64:hello world & foo bar [0-10]]"},
		{"enc b85", "[b85:hello world & foo bar [0-10]]"},
		{"enc bin", "[bin:hello world & foo bar [0-10]]"},
		{"enc hex", "[hex:hello world & foo bar [0-10]]"},
		{"enc url", "[url:hello world & foo bar [0-10]]"},
		{"data", "[:a.txt] == [:b.txt] | [:test] == [:test/02]"},
		{"fn", "time is: [$time:]"},
		{"nesting", "hi there, [this is hardcoded,[$foo:it's;[a generated,an artificial];string]]"},
		{"nesting", "[[$foo:[:test/02];[:test/02];#[0:9]],[$foo:[:a.txt];[:b.txt]]], there are [0:10] new messages for you."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 10; i++ {
				t.Logf("%s | %s => %s\n", tt.name, tt.str, Fuzz(tt.str))
			}
		})
	}
}
