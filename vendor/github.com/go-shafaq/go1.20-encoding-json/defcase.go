package json

var defCase = func(pkgPath, name string) string {
	return name
}

func DefCase(f func(tag string) func(pkgPath, name string) string) {
	defCase = f("json")
}
