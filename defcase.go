package sonic

import (
	json "github.com/go-shafaq/go1.20-encoding-json"
)

func DefCase(f func(tag string) func(pkgPath, name string) string) {
	json.DefCase(f)
}
