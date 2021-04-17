package funcs

func KeyExists(key string, m map[string]string) bool {

	_, ok := m[key]
	return ok
}
