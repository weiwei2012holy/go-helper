package string

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

//Any2String 转字符串
func Any2String(v any) string {
    return fmt.Sprintf("%v", v)
}

//StrRandom 生成随机字符串
func StrRandom(size int) string {
    if size <= 0 {
        size = 16
    }
    str := ""
    for size > 0 {
        rand.Seed(time.Now().UnixNano())
        token := make([]byte, size)
        rand.Read(token)
        s := StrReplace(Base64Encode(token), []string{"=", "/", "+"}, "")
        if len(s) >= size {
            str += s[:size]
            break
        } else {
            str += s
            size -= len(s)
        }
    }
    return str
}

//StrReplace 字符串查找替换
func StrReplace(src string, search []string, replace string) string {
    for i := 0; i < len(search); i++ {
        src = strings.ReplaceAll(src, search[i], replace)
    }
    return src
}

//StrRepeat str_repeat
func StrRepeat(input string, multiplier int) string {
    return strings.Repeat(input, multiplier)
}

// Explode explode — 使用一个字符串分割另一个字符串
func Explode(separator, str string) []string {
    if str == "" {
        return nil
    }
    return strings.Split(str, separator)
}

// Implode implode — 用字符串连接数组元素
func Implode(src []string, separator string) string {
    return strings.Join(src, separator)
}

// Trim trim — 去除字符串首尾处的空白字符（或者其他字符）
func Trim(src string, characters ...string) string {
    for _, char := range characters {
        src = strings.Trim(src, char)
    }
    return src
}

//Contains 验证字符串包含部分子字符串
func Contains(src string, needles ...string) bool {
    for i := 0; i < len(needles); i++ {
        if strings.Contains(src, needles[i]) {
            return true
        }
    }
    return false
}

//ContainsAll 验证字符串包含全部子字符串
func ContainsAll(src string, needles ...string) bool {
    for i := 0; i < len(needles); i++ {
        if !strings.Contains(src, needles[i]) {
            return false
        }
    }
    return true
}

//EndsWith 验证字符串包含特定字符结尾
func EndsWith(src string, needles ...string) bool {
    for i := 0; i < len(needles); i++ {
        if needles[i] == "" || strings.HasSuffix(src, needles[i]) {
            return true
        }
    }
    return false
}

// StartsWith 验证字符串包含特定字符开始
func StartsWith(src string, needles ...string) bool {
    for i := 0; i < len(needles); i++ {
        if needles[i] == "" || strings.HasPrefix(src, needles[i]) {
            return true
        }
    }
    return false
}

// Strpos strpos
func Strpos(src, needle string) int {
    return strings.Index(src, needle)
}

// MbStrpos 查找字符串在另一个字符串中首次出现的位置
func MbStrpos(src, needle string) int {
    idx := strings.Index(src, needle)
    if idx <= 0 {
        return idx
    }
    reader := strings.NewReader(src)
    newIdx := 0
    for {
        _, size, err := reader.ReadRune()
        if err != nil {
            return -1
        }
        idx -= size
        newIdx++
        if idx <= 0 {
            return newIdx
        }
    }
}

// Strlen  strlen 获取字符串长度
func Strlen(src string) int {
    return len(src)
}

// MbStrlen mb_strlen 获取字符串长度
func MbStrlen(src string) int {
    return len([]rune(src))
}

func substr(src string, offset, length int) string {
    return src[offset : offset+length]
}

// Before 截取第一个子串 N 前的字符串
func Before(src, needle string) string {
    if needle == "" {
        return src
    }
    idx := strings.Index(src, needle)
    if idx <= 0 {
        return ""
    }
    return src[:idx]
}

// BeforeLast 截取最后一个子串 N 前的字符串
func BeforeLast(src, needle string) string {
    if needle == "" {
        return src
    }
    idx := strings.LastIndex(src, needle)
    if idx <= 0 {
        return ""
    }
    return src[:idx]
}

// After 截取第一个子串 N 后的字符串
func After(src, needle string) string {
    if needle == "" {
        return src
    }
    idx := strings.Index(src, needle)
    if idx == -1 {
        return ""
    }
    return src[idx+len(needle):]
}

// AfterLast  截取最后一个子串 N 后的字符串
func AfterLast(src, needle string) string {
    if needle == "" {
        return src
    }
    idx := strings.LastIndex(src, needle)
    if idx == -1 {
        return ""
    }
    return src[idx+len(needle):]
}
