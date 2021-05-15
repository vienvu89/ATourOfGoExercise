func WordCount(s string) map[string]int {
	m := make(map[string] int)

	list :=  strings.FieldsFunc(s, unicode.IsSpace)
	for i := 0; i < len(list); i++ {
		a, found := m[list[i]]
		if found {
			m[list[i]] = a + 1
		} else {
			m[list[i]] = 1
		}
	}
	return m
}