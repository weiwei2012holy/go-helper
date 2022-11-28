package string

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestStrRandom(t *testing.T) {

    for i := 1; i < 100; i++ {
        res := StrRandom(i)
        assert.Equal(t, i, len(res))
        //        t.Log(res)
    }

}

func TestBase64(t *testing.T) {
    src := Any2String(time.Now())
    t.Log(src)

    b := Base64Encode([]byte(src))
    t.Log(b)

    db, err := Base64Decode(b)
    assert.NoError(t, err)
    assert.Equal(t, src, string(db))
}

func TestMD5(t *testing.T) {
    type args struct {
        str string
    }
    tests := []struct {
        name string
        args args
        want string
    }{

        {
            name: "",
            args: args{
                str: "123456",
            },
            want: "e10adc3949ba59abbe56e057f20f883e",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := MD5(tt.args.str); got != tt.want {
                t.Errorf("MD5() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestExplode(t *testing.T) {
    str := "a,b,c,d,e"
    r := Explode(",", str)
    assert.Equal(t, 5, len(r))

    str2 := ""
    r2 := Explode(",", str2)
    assert.Equal(t, 0, len(r2))
    t.Log(r2)
}

func TestImplode(t *testing.T) {
    type args struct {
        src       []string
        separator string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "a1",
            args: args{
                src:       []string{"a", "b", "c", "d", "e"},
                separator: ",",
            },
            want: "a,b,c,d,e",
        },
        {
            name: "a2",
            args: args{
                src:       []string{"a", "b", "c", "d", "e"},
                separator: "",
            },
            want: "abcde",
        },
        {
            name: "a3",
            args: args{
                src:       []string{},
                separator: "",
            },
            want: "",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, Implode(tt.args.src, tt.args.separator), "Implode(%v, %v)", tt.args.src, tt.args.separator)
        })
    }
}

func TestContains(t *testing.T) {
    type args struct {
        src     string
        needles []string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{

        {
            name: "c1",
            args: args{
                src:     "abcdefg",
                needles: []string{"a", "c", "e"},
            },
            want: true,
        },
        {
            name: "c1",
            args: args{
                src:     "abcdefg",
                needles: []string{"aaa"},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, Contains(tt.args.src, tt.args.needles...), "Contains(%v, %v)", tt.args.src, tt.args.needles)
        })
    }
}

func TestContainsAll(t *testing.T) {
    type args struct {
        src     string
        needles []string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "c1",
            args: args{
                src:     "abcdefg",
                needles: []string{"a", "b", "c"},
            },
            want: true,
        },
        {
            name: "c1",
            args: args{
                src:     "abcdefg",
                needles: []string{"a", "a1", "c"},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, ContainsAll(tt.args.src, tt.args.needles...), "ContainsAll(%v, %v)", tt.args.src, tt.args.needles)
        })
    }
}

func TestEndsWith(t *testing.T) {
    type args struct {
        src     string
        needles []string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "a1",
            args: args{
                src:     "abcde",
                needles: []string{"a", "e"},
            },
            want: true,
        },
        {
            name: "a2",
            args: args{
                src:     "abcde",
                needles: []string{"a", "e1"},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, EndsWith(tt.args.src, tt.args.needles...), "EndsWith(%v, %v)", tt.args.src, tt.args.needles)
        })
    }
}

func TestStartsWith(t *testing.T) {
    type args struct {
        src     string
        needles []string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{

        {
            name: "a1",
            args: args{
                src:     "abcde",
                needles: []string{"a", "b"},
            },
            want: true,
        },
        {
            name: "a2",
            args: args{
                src:     "abcde",
                needles: []string{"b1", "b"},
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, StartsWith(tt.args.src, tt.args.needles...), "StartsWith(%v, %v)", tt.args.src, tt.args.needles)
        })
    }
}
