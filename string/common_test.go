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
