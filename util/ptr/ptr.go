package ptr

func Bool(value bool) *bool {
    return &value
}

func Uint64(value uint64) *uint64 {
    return &value
}

func Int64(value int64) *int64 {
    return &value
}

//Converts a string to pointer form.
//Helper function to avoid
//myValue := "myValue"
//someField = &myValue.
//Can now do: someField = ptr.String("myValue")
func String(value string) *string {
    return &value
}
