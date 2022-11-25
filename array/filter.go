package array

// Where 过滤数组
func Where[V any](s []V, filter func(V) bool) (res []V) {
    for i := 0; i < len(s); i++ {
        if filter(s[i]) {
            res = append(res, s[i])
        }
    }
    return
}

// First 过滤数组获取其中一条记录
func First[V any](s []V, filter func(V) bool) (res V) {
    for i := 0; i < len(s); i++ {
        if filter(s[i]) {
            return s[i]
        }
    }
    return
}
