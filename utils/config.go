package utils

func WithDefault(m map[string]interface{}, k string, d interface{}) interface{} {
	if v, ok := m[k]; ok {
		return v
	}
	return d
}
