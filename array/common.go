package array

import (
    "errors"
    "go-helper/lib"
    "math"
    "math/rand"
    "time"
)

// Keys array_keys 返回MAP键名
func Keys[K comparable, V any](m map[K]V) []K {
    var res = make([]K, 0)
    for k := range m {
        res = append(res, k)
    }
    return res
}

// Values array_values 返回MAP中所有的值
func Values[K comparable, V any](m map[K]V) []V {
    var res = make([]V, 0)
    for _, v := range m {
        res = append(res, v)
    }
    return res
}

// Unique unique 去重
func Unique[V comparable](s []V) []V {
    tmp := make(map[V]struct{})
    res := make([]V, 0)
    for i := 0; i < len(s); i++ {
        _, e := tmp[s[i]]
        if e {
            continue
        }
        tmp[s[i]] = struct{}{}
        res = append(res, s[i])
    }
    return res
}

// Combine array_combine 两个数组合并
func Combine[K comparable, V any](sk []K, sv []V) (map[K]V, error) {
    if len(sk) != len(sv) {
        return nil, errors.New("combine slice numbers not equal")
    }
    if len(sv) == 0 || len(sv) == 0 {
        return nil, nil
    }
    var data = make(map[K]V, len(sk))
    for i, k := range sk {
        data[k] = sv[i]
    }
    return data, nil
}

// Shuffle 打乱数组
func Shuffle[V any](s *[]V) {
    if len(*s) <= 1 {
        return
    }
    rand.Seed(time.Now().Unix())
    rand.Shuffle(len(*s), func(i, j int) {
        (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
    })
    return
}

// Sum array_sum
func Sum[V lib.Integers | lib.Float](s []V) V {
    var sum V
    for _, v := range s {
        sum += v
    }
    return sum
}

// Chunk array_chunk 将一个数组分割成多个
func Chunk[V any](s []V, size int) [][]V {
    length := len(s)
    if length == 0 {
        return nil
    }
    if size <= 0 {
        size = length
    }
    var chunks int
    chunks = int(math.Ceil(float64(length) / float64(size)))
    res := make([][]V, 0, chunks)
    for i := 0; i < chunks; i++ {
        end := (i + 1) * size
        if end > length {
            end = length
        }
        res = append(res, s[i*size:end])
    }
    return res
}

// CountValues array_count_values 统计数组中所有的值
func CountValues[V comparable](s []V) map[V]int {
    res := make(map[V]int)
    for _, v := range s {
        res[v]++
    }
    return res
}

//Column array_column
func Column[V any, C any](s []V, column func(V) C) (res []C) {
    for i := 0; i < len(s); i++ {
        res = append(res, column(s[i]))
    }
    return
}

//ColumnWithKey array_column with key support
func ColumnWithKey[V any, K comparable, C any](s []V, kc func(V) (K, C)) (res map[K]C) {
    res = make(map[K]C, len(s))
    for i := 0; i < len(s); i++ {
        k, c := kc(s[i])
        res[k] = c
    }
    return
}

//Rand array_rand 从slice中随机取出一个或多个随机键
func Rand[V any](src []V, num int) (res []int) {
    if num <= 0 || len(src) == 0 {
        return
    }
    rand.Seed(time.Now().UnixNano())
    res = rand.Perm(len(src))
    if len(res) <= num {
        return
    }
    res = res[:num]
    return
}

//Range 根据范围创建数组，包含指定的元素
func Range[V lib.Integers](start V, end V, step V) (res []V) {
    if step <= 0 {
        step = 1
    }
    for start <= end {
        res = append(res, start)
        start += step
    }
    return

}
