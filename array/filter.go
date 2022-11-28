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

// GroupBy 对 slice 字段进行分组
func GroupBy[V any, K comparable](s []V, key func(V) K) map[K][]V {
    res := make(map[K][]V)
    for i := 0; i < len(s); i++ {
        k := key(s[i])
        d, e := res[k]
        if !e {
            d = make([]V, 0)
        }
        res[k] = append(d, s[i])
    }
    return res
}
