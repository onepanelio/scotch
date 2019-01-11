package str

import (
    "strings"
    "unicode"
)

func StripWhitespace(value string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return -1
        }
        return r
    }, value)
}
