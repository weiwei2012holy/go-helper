package array

import (
    "testing"
)

func TestGroupBy(t *testing.T) {
    s := []testStruct{
        {
            Age:  12,
            Name: "aaa",
        },
        {
            Age:  12,
            Name: "bbb",
        },
        {
            Age:  22,
            Name: "ccc",
        },
        {
            Age:  22,
            Name: "ddd",
        },
    }

    r1 := GroupBy(s, func(v testStruct) int {
        return v.Age
    })
    t.Log(r1)

    r2 := GroupBy(s, func(v testStruct) string {
        return v.Name
    })
    t.Log(r2)

    r3 := GroupBy(s, func(v testStruct) string {
        return "asd"
    })
    t.Log(r3)

}
