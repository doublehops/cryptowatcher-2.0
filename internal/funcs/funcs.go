package funcs

func KeyExists(key string, m map[string]interface{}) bool {

	_, ok := m[key]
	return ok
}
