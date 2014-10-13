package cfg

var cfg = make(map[string]string)

func Set(key string, val string) {
    cfg[key] = val
}

func Get(key string) string {
    return cfg[key]
}