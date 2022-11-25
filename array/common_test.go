package array

import (
    "fmt"
    "github.com/shopspring/decimal"
    "github.com/stretchr/testify/assert"
    string2 "go-helper/string"
    "math"
    "math/rand"
    "testing"
    "time"
)

type testStruct struct {
    Age  int
    Name string
}

func TestArr(t *testing.T) {
    mS := map[string]string{}
    mInt8 := map[int8]int8{}
    mInt16 := map[int16]int16{}
    mInt32 := map[int32]int32{}
    mInt64 := map[int64]int64{}
    mInt := map[int]int{}

    mUint8 := map[uint8]uint8{}
    mUint16 := map[uint16]uint16{}
    mUint32 := map[uint32]uint32{}
    mUint64 := map[uint64]uint64{}

    mFloat32 := map[float32]float32{}
    mFloat64 := map[float64]float64{}

    complex64Map := map[complex64]complex64{}
    complex128Map := map[complex128]complex128{}
    size := 10

    for i := 1; i <= size; i++ {
        mS["key-"+string2.Any2String(i)] = "value-" + string2.Any2String(i)
        mInt8[int8(i)] = int8(i)
        mInt16[int16(i)] = int16(i)
        mInt32[int32(i)] = int32(i)
        mInt64[int64(i)] = int64(i)

        mUint8[uint8(i)] = uint8(i)
        mUint16[uint16(i)] = uint16(i)
        mUint32[uint32(i)] = uint32(i)
        mUint64[uint64(i)] = uint64(i)

        mFloat32[float32(i)+float32(i/10)] = float32(i) + float32(i/10)
        mFloat64[float64(i)+float64(i/10)] = float64(i) + float64(i/10)

        complex64Map[complex(float32(i), float32(i+1))] = complex(float32(i), float32(i+1))
        complex128Map[complex(float64(i), float64(i+1))] = complex(float64(i), float64(i+1))

        mInt[i] = i
    }

    t.Run("Keys", func(t *testing.T) {

        fmt.Println(Keys(mS))
        fmt.Println(Keys(mFloat64))
        fmt.Println(Keys(complex128Map))
    })
    t.Run("Values", func(t *testing.T) {
        fmt.Println(Values(mS))
        fmt.Println(Values(mInt8))
    })

}

func TestUnique(t *testing.T) {
    s := make([]int, 0)
    s2 := make([]float64, 0)
    s3 := make([]string, 0)
    for i := 0; i < 10; i++ {
        s = append(s, 1)
        s2 = append(s2, 1.23)
        s3 = append(s3, "123")
    }
    su := Unique(s)
    su2 := Unique(s2)
    su3 := Unique(s3)
    assert.Equal(t, len(su), 1)
    assert.Equal(t, len(su2), 1)
    assert.Equal(t, len(su3), 1)
}

func TestChan(t *testing.T) {
    size := 5
    minSize := 1

    t.Run("Pop", func(t *testing.T) {
        c := getSliceInt(minSize, minSize)
        c2 := getSliceInt(size, size)
        v1, e := Pop(&c)
        assert.Equal(t, minSize, v1)
        assert.True(t, !e)
        _, e = Pop(&c)
        assert.True(t, e)

        v2, _ := Pop(&c2)
        assert.Equal(t, size, v2)
        assert.Equal(t, len(c2), size-1)
    })

    t.Run("Shift", func(t *testing.T) {
        c := getSliceInt(minSize, minSize)
        c2 := getSliceInt(size, size)
        v1, e := Shift(&c)
        assert.Equal(t, minSize, v1)
        assert.True(t, !e)
        _, e = Shift(&c)
        assert.True(t, e)

        v2, _ := Shift(&c2)
        assert.Equal(t, 1, v2)
        assert.Equal(t, len(c2), size-1)

    })

}

func TestShuffle(t *testing.T) {
    s := getSliceInt(5, 5)

    Shuffle(&s)
    fmt.Println(s)
}

func TestSum(t *testing.T) {
    s := getSliceInt(5, 5)
    v := Sum(s)
    fmt.Println(s, v)

    s2 := getSliceFloat(5)
    v2 := Sum(s2)
    fmt.Println(s2, v2)
}

func TestComplex(t *testing.T) {
    v := complex(1.23, 2.34)
    v2 := complex(2, 2)
    fmt.Println(v + v2)

    fmt.Println(math.Ceil(0.121))

    len := func() {
        fmt.Println("len")
    }

    len()
}

func TestAppend(t *testing.T) {

    type A struct {
        v []int
    }
    a := A{}

    a.v = append(a.v, 1)
    fmt.Println(a.v)
}

func TestChunk(t *testing.T) {

    size := 11
    s := getSliceInt(size, size)

    c1 := 2
    r1 := Chunk(s, c1)
    assert.Equal(t, size/c1+1, len(r1))
    r2 := Chunk(s, -1)
    assert.Equal(t, 1, len(r2))
    fmt.Println(r1, r2)

    s3 := make([]int, 0)
    r3 := Chunk(s3, 1)
    assert.Equal(t, 0, len(r3))
}

func getSliceInt(size int, max int) []int {
    s := make([]int, 0, size)
    for i := max - size + 1; i <= max; i++ {
        s = append(s, i)
    }
    return s
}

func getSliceStruct(size int, max int) []testStruct {
    s := make([]testStruct, 0, size)
    for i := max - size + 1; i <= max; i++ {
        s = append(s, testStruct{
            Age:  i,
            Name: fmt.Sprintf("name-%d", i),
        })
    }
    return s
}

func getSliceFloat(size int) []float64 {
    s := make([]float64, 0, size)

    rand.Seed(time.Now().Unix())

    for i := 0; i < size; i++ {
        v := decimal.NewFromFloat32(rand.Float32()).Add(decimal.NewFromInt(int64(i))).Round(2).InexactFloat64()
        s = append(s, v)
    }
    return s
}

func TestWhere(t *testing.T) {

    t.Run("sliceInt", func(t *testing.T) {
        s := getSliceInt(5, 5)
        fmt.Println(s)
        res := Where(s, func(v int) bool {
            return v > 2
        })
        assert.Equal(t, 3, len(res))
        fmt.Print(res)
    })

    t.Run("sliceStruct", func(t *testing.T) {
        s := getSliceStruct(5, 5)
        fmt.Println(s)
        res := Where(s, func(v testStruct) bool {
            return v.Age > 2
        })
        assert.Equal(t, 3, len(res))
        fmt.Print(res)
    })

}

func TestFirst(t *testing.T) {
    t.Run("sliceInt", func(t *testing.T) {
        s := getSliceInt(5, 5)
        fmt.Println(s)
        res := First(s, func(v int) bool {
            return v == 2
        })
        assert.Equal(t, 2, res)
        fmt.Print(res)
    })
    t.Run("sliceStruct", func(t *testing.T) {
        s := getSliceStruct(5, 5)
        res := First(s, func(v testStruct) bool {
            return v.Age == 2
        })
        assert.Equal(t, 2, res.Age)
        fmt.Print(res)
    })

}

func TestColumn(t *testing.T) {
    s := getSliceStruct(5, 5)
    t.Run("column-1", func(t *testing.T) {
        r1 := Column(s, func(v testStruct) int {
            return v.Age
        })
        t.Log(r1)
    })
    t.Run("column-2", func(t *testing.T) {
        r2 := Column(s, func(v testStruct) string {
            return v.Name
        })
        t.Log(r2)
    })
    t.Run("k-v column", func(t *testing.T) {
        r3 := ColumnWithKey(s, func(v testStruct) (string, int) {
            return v.Name, v.Age
        })
        t.Log(r3)
    })
    t.Run("k-v slice", func(t *testing.T) {
        r4 := ColumnWithKey(s, func(v testStruct) (string, testStruct) {
            return v.Name, v
        })
        t.Log(r4)
    })

}

func TestRand(t *testing.T) {
    s := getSliceStruct(5, 5)
    r := Rand(s, 2)
    t.Log(r)

    r2 := Rand(s, 10)
    t.Log(r2)
}

func TestRange(t *testing.T) {

    s := Range(1, 10, 1)
    assert.Equal(t, 10, len(s))
    t.Log(s)
    s2 := Range(1, 11, 3)
    assert.Equal(t, 4, len(s2))
    t.Log(s2)
}

func TestCombine(t *testing.T) {
    k := []string{"a", "b", "c", "d", "e"}
    v := []int{1, 2, 3, 4, 5}
    r, err := Combine(k, v)
    assert.NoError(t, err)
    t.Log(r)
    r2, err := Combine(v, k)
    assert.NoError(t, err)
    t.Log(r2)
    k = append(k, "test")
    _, err = Combine(k, v)
    assert.Error(t, err)
}
