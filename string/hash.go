package string

import (
    "crypto/md5"
    "fmt"
    "io"
)

// MD5 md5
func MD5(str string) string {
    h := md5.New()
    _, _ = io.WriteString(h, str)
    return fmt.Sprintf("%x", h.Sum(nil))
}
