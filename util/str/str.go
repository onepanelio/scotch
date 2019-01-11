package str

import (
    "strings"
    "unicode"
)

//Converts a string to pointer form.
//Helper function to avoid
//myValue := "myValue"
//someField = &myValue.
//Can now do: someField = PStr("myValue")
func P(value string) *string {
    return &value
}

func StripWhitespace(value string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, value)
}
