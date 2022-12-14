package string

import (
    "strings"
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
        {
            name: "a2",
            args: args{
                src:     "abcde??????",
                needles: []string{"??????", "???"},
            },
            want: true,
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
        {
            name: "a3",
            args: args{
                src:     "????????????",
                needles: []string{"??????"},
            },
            want: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, StartsWith(tt.args.src, tt.args.needles...), "StartsWith(%v, %v)", tt.args.src, tt.args.needles)
        })
    }
}

func TestStrpos(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "s1",
            args: args{
                src:    "abcde",
                needle: "a",
            },
            want: 0,
        },
        {
            name: "s2",
            args: args{
                src:    "abcde",
                needle: "abc",
            },
            want: 0,
        },
        {
            name: "s3",
            args: args{
                src:    "abcde",
                needle: "aaa",
            },
            want: -1,
        },
        {
            name: "s4",
            args: args{
                src:    "abababab",
                needle: "a",
            },
            want: 0,
        },
        {
            name: "5",
            args: args{
                src:    "????????????",
                needle: "???",
            },
            want: 1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, Strpos(tt.args.src, tt.args.needle), "Strpos(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}
func TestStringSplit(t *testing.T) {
    s := "a,b,c,d,e,f"
    r1 := strings.Split(s, ",") //???????????????
    t.Log(r1, len(r1))          //[a b c d e f] 6

    r2 := strings.SplitN(s, ",", 2) //???????????????-????????? N ???
    t.Log(r2, len(r2))              //[a b,c,d,e,f] 2

    r3 := strings.SplitAfter(s, ",") //???????????????-???????????????????????????
    t.Log(r3, len(r3))               //[a, b, c, d, e, f] 6

    r4 := strings.SplitAfterN(s, ",", 2)
    t.Log(r4, len(r4)) //[a, b,c,d,e,f] 2
}

func TestMbStrpos(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "s1",
            args: args{
                src:    "a1????????????",
                needle: "???",
            },
            want: 3,
        },
        {
            name: "s2",
            args: args{
                src:    "a1????????????",
                needle: "a1",
            },
            want: 0,
        },
        {
            name: "s3",
            args: args{
                src:    "????????????",
                needle: "??????",
            },
            want: 2,
        },
        {
            name: "s3",
            args: args{
                src:    "????????????",
                needle: "a",
            },
            want: -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, MbStrpos(tt.args.src, tt.args.needle), "MbStrpos(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}

func TestMbStrlen(t *testing.T) {
    type args struct {
        src string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "a1",
            args: args{
                src: "????????????",
            },
            want: 4,
        },
        {
            name: "a1",
            args: args{
                src: "a1????????????",
            },
            want: 6,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, MbStrlen(tt.args.src), "MbStrlen(%v)", tt.args.src)
        })
    }
}

func TestBefore(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "b1",
            args: args{
                src:    "abcde",
                needle: "c",
            },
            want: "ab",
        },
        {
            name: "b2",
            args: args{
                src:    "abc????????????",
                needle: "???",
            },
            want: "abc???",
        },
        {
            name: "b2",
            args: args{
                src:    "????????????",
                needle: "???",
            },
            want: "",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, Before(tt.args.src, tt.args.needle), "Before(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}

func TestAfter(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "b1",
            args: args{
                src:    "abcde",
                needle: "c",
            },
            want: "de",
        },
        {
            name: "b2",
            args: args{
                src:    "abc????????????",
                needle: "???",
            },
            want: "??????",
        },
        {
            name: "b3",
            args: args{
                src:    "????????????",
                needle: "???",
            },
            want: "",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, After(tt.args.src, tt.args.needle), "After(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}

func TestBeforeLast(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "b1",
            args: args{
                src:    "abcabc",
                needle: "b",
            },
            want: "abca",
        },
        {
            name: "b2",
            args: args{
                src:    "????????????????????????",
                needle: "??????",
            },
            want: "??????????????????",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, BeforeLast(tt.args.src, tt.args.needle), "BeforeLast(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}

func TestAfterLast(t *testing.T) {
    type args struct {
        src    string
        needle string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "b1",
            args: args{
                src:    "abcabc",
                needle: "b",
            },
            want: "c",
        },
        {
            name: "b2",
            args: args{
                src:    "????????????????????????",
                needle: "??????",
            },
            want: "??????",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, AfterLast(tt.args.src, tt.args.needle), "AfterLast(%v, %v)", tt.args.src, tt.args.needle)
        })
    }
}

func TestSubstr(t *testing.T) {
    s := "abc"
    for i := 0; i <= 4; i++ {
        r := Substr(s, 1, i)
        t.Log(r)
    }
}

func TestStrlen(t *testing.T) {
    s1 := "abc"
    s2 := "??????"
    s3 := "??????abc"
    assert.Equal(t, 3, Strlen(s1))
    assert.Equal(t, 3, MbStrlen(s1))
    assert.NotEqual(t, 2, Strlen(s2))
    assert.Equal(t, 2, MbStrlen(s2))
    assert.NotEqual(t, 5, Strlen(s3))
    assert.Equal(t, 5, MbStrlen(s3))
}
