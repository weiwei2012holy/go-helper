package file

import (
    "bytes"
    "fmt"
    "github.com/stretchr/testify/assert"
    "os"
    "strconv"
    "testing"
)

var dirName = "temp"
var fileName = "file.test.txt"
var filePath = dirName + "/" + fileName

func TestFileContent(t *testing.T) {
    t1 := "first line\n"
    t2 := "append line \n"
    _ = Unlink(filePath)
    t.Run("append", func(t *testing.T) {
        err := FPutContent(filePath, []byte(t1), true)
        assert.NoError(t, err)
        err2 := FPutContent(filePath, []byte(t2), true)
        assert.NoError(t, err2)
        content, err := FGetContents(filePath)
        assert.NoError(t, err)

        assert.Equal(t, t1+t2, (string)(content))
        fmt.Println((string)(content))
    })
    t.Run("not append", func(t *testing.T) {
        err := FPutContent(filePath, []byte(t1), false)
        assert.NoError(t, err)
        err2 := FPutContent(filePath, []byte(t2), false)
        assert.NoError(t, err2)
        content, err := FGetContents(filePath)
        assert.NoError(t, err)

        assert.Equal(t, t2, (string)(content))
        fmt.Println((string)(content))
    })

}

func TestUnlink(t *testing.T) {
    err := Unlink(filePath)
    assert.NoError(t, err)
}
func getStringData() bytes.Buffer {
    var builder bytes.Buffer

    for i := 0; i < 5; i++ {
        builder.WriteString("text-")
        builder.WriteString(strconv.Itoa(i))
        builder.WriteString("\n")
    }
    return builder
}

func TestDir(t *testing.T) {

    dir, err := __DIR__()
    assert.NoError(t, err)
    fmt.Println(dir)

}

func TestPathInfo(t *testing.T) {

    res := PathInfo(filePath)
    fmt.Println(res)
    dir, _ := __DIR__()

    res2 := PathInfo(dir + filePath)
    fmt.Println(res2)
}

func TestIsDir(t *testing.T) {
    dir, _ := __DIR__()
    t.Run("true", func(t *testing.T) {
        isDir, err := IsDir(dir)
        assert.NoError(t, err)
        assert.True(t, isDir)
    })
    t.Run("false", func(t *testing.T) {
        isDir, err := IsDir(dir + "/" + fileName)
        assert.NoError(t, err)
        assert.True(t, !isDir)
    })
}

func TestIsFile(t *testing.T) {
    dir, _ := __DIR__()
    t.Run("true", func(t *testing.T) {
        isDir, err := IsFile(dir + "/" + fileName)
        assert.NoError(t, err)
        assert.True(t, isDir)
    })
    t.Run("false", func(t *testing.T) {
        isDir, err := IsFile(dir)
        assert.NoError(t, err)
        assert.True(t, !isDir)
    })

}

func TestFileReadWrit(t *testing.T) {

    _ = Unlink(filePath)
    f, err := FOpen(filePath, os.O_WRONLY|os.O_CREATE)
    assert.NoError(t, err)

    for i := 10; i <= 20; i++ {
        n, err := FWrite(f, []byte(strconv.Itoa(i)))
        assert.NoError(t, err)
        assert.True(t, n > 0)
    }
    err = FClose(f)
    assert.NoError(t, err)

    f, err = FOpen(filePath, os.O_RDONLY)
    assert.NoError(t, err)

    b, err := FRead(f)
    assert.NoError(t, err)
    fmt.Println(string(b))
    err = FClose(f)
    assert.NoError(t, err)

    c := 'b'
    fmt.Println(c)

}
