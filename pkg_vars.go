package fuzzer

import (
	"regexp"
	"sync"
)

var (
	reRandHash    = regexp.MustCompile(`#\d+`)
	reRandInt     = regexp.MustCompile(`\[\-{0,1}\d+:\-{0,1}\d+\]`)
	reIntRange    = regexp.MustCompile(`\[\-{0,1}\d+\.\.\-{0,1}\d+\]`)
	reRandFloat   = regexp.MustCompile(`\[\-{0,1}\d+\.\d+:\-{0,1}\d+\.\d+\]`)
	reRandStr     = regexp.MustCompile(`(.+?,)+([^,]+)`)
	importedFiles = map[string]map[string]string{}
	cache         = &memCache{
		mu:      &sync.Mutex{},
		content: map[string][]string{},
	}
	fz *fuzzer
)
