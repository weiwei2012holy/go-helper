package string

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestJsonEncode(t *testing.T) {
    type args struct {
        data any
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{{
        name: "encode1",
        args: args{
            data: map[string]int{
                "a": 1, "b": 2, "c": 3,
            },
        },
        want:    `{"a":1,"b":2,"c":3}`,
        wantErr: false,
    },
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := JsonEncode(tt.args.data)
            t.Log(got)
            if (err != nil) != tt.wantErr {
                t.Errorf("JsonEncode() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("JsonEncode() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestJsonDecode(t *testing.T) {
    data := map[string]int{}
    err := JsonDecode(`{"a":1,"b":2,"c":3}`, &data)
    assert.NoError(t, err)
    t.Log(data)
}
