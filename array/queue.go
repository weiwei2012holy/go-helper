package array

// Pop array_pop 尾部出栈
func Pop[V any](s *[]V) (v V, empty bool) {
    if len(*s) == 0 {
        empty = true
        return
    }
    v = (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return
}

// Shift array_shift 头部出栈
func Shift[V any](s *[]V) (v V, empty bool) {
    if len(*s) == 0 {
        empty = true
        return
    }
    v = (*s)[0]
    *s = (*s)[1:]
    return
}

// Push array_push 尾部入栈
func Push[V any](s *[]V, elements ...V) int {
    *s = append(*s, elements...)
    return len(*s)
}

// Unshift array_unshift 头部入栈
func Unshift[V any](s *[]V, elements ...V) int {
    *s = append(elements, *s...)
    return len(*s)
}
