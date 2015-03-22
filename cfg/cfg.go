package cfg

var cfg = make(map[string]interface{})

func Set(key string, val interface{}) {
	cfg[key] = val
}

func Get(key string) interface{} {
	return cfg[key]
}
