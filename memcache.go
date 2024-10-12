package fuzzer

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/toxyl/errors"
	"github.com/toxyl/flo"
)

type memCache struct {
	mu      *sync.Mutex
	content map[string][]string
}

func explode(str string) []string {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.ReplaceAll(str, "\r", "\n")
	res := []string{}
	// filter out empty lines and commented lines (first non-space char is #)
	for _, s := range strings.Split(str, "\n") {
		st := strings.TrimSpace(s)
		if st == "" || st[0] == '#' {
			continue
		}
		res = append(res, s)
	}
	return res
}

func randomLineFromCache(path string) (data string, found bool) {
	path = strings.Trim(path, "/")
	fnGetLine := func(data []string) string {
		l := len(data) - 1
		if l < 0 {
			return ""
		}
		return data[randomInt(0, l)]
	}

	// check if we have the data cached
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if d, ok := cache.content[path]; ok {
		return fnGetLine(d), true // return the cached data
	}

	// nope, data is not cached, good time to do so.
	cache.content[path] = []string{}

	// we have a few options what the path could refer to:

	// a) a directory
	if d, ok := importedFiles[path]; ok {
		// this is a directory, so we append all file entries to the cache
		for _, s := range d {
			cache.content[path] = append(cache.content[path], explode(s)...)
		}

		// there might be subdirectories, let's append their files too
		pp := path + "/" // we don't care about the depth, so anything prefixed with the path + slash must be a subdirectory
		for p, d := range importedFiles {
			if strings.HasPrefix(p, pp) {
				// this is a sub directory, append the data
				for _, s := range d {
					cache.content[path] = append(cache.content[path], explode(s)...)
				}
			}
		}
		return fnGetLine(cache.content[path]), true
	}

	// b) a file
	if strings.Count(path, "/") == 0 {
		if d, ok := importedFiles[""]; ok {
			if s, ok := d[path]; ok {
				cache.content[path] = append(cache.content[path], explode(s)...)
				return fnGetLine(cache.content[path]), true // we found the file
			}
		}
		return fnGetLine(cache.content[path]), false // we didn't find the file or section
	}

	// c) a directory + file
	if d, ok := importedFiles[filepath.Dir(path)]; ok {
		if s, ok := d[filepath.Base(path)]; ok {
			cache.content[path] = append(cache.content[path], explode(s)...)
			return fnGetLine(cache.content[path]), true // we found the file
		}
		return fnGetLine(cache.content[path]), false // we didn't find the file
	}

	return fnGetLine(cache.content[path]), false // we didn't find anything
}

func getRandomLine(path string) string {
	data, found := randomLineFromCache(path)
	if !found {
		return ""
	}
	return data
}

func importFromDir(dir string) error {
	if dir == "" {
		return errors.Newf("no import path given")
	}

	baseDir := flo.Dir(dir)
	if !baseDir.Exists() {
		return errors.Newf("import path not found")
	}
	baseDirPath := strings.TrimPrefix(strings.ToLower(baseDir.Path()), "/")

	parsePath := func(path, basePath string) (section, data string) {
		path = strings.TrimPrefix(path, basePath)
		data = filepath.Base(path)
		section = filepath.Dir(path)
		section = strings.TrimPrefix(section, "/")
		section = strings.TrimPrefix(section, ".")
		return
	}

	baseDir.Each(func(f *flo.FileObj) {
		section, data := parsePath(f.Path(), baseDir.Path())
		section = strings.TrimPrefix(strings.TrimPrefix(section, baseDirPath), "/")
		if _, ok := importedFiles[section]; !ok {
			importedFiles[section] = map[string]string{}
		}
		importedFiles[section][data] = f.AsString()
	}, nil)

	return nil
}
