package file

import (
    "bufio"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// FGetContents 一次性读取文件数据:file_get_contents()
func FGetContents(fileName string) ([]byte, error) {
    return ioutil.ReadFile(fileName)
}

// FPutContent 一次性写入数据到文件:file_put_content()，append=true表示追加写入
func FPutContent(fileName string, data []byte, append bool) error {
    var flag int
    if append {
        flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
    } else {
        flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
    }
    f, err := os.OpenFile(fileName, flag, 0666)
    if err != nil {
        return err
    }
    _, err = f.Write(data)
    if err1 := f.Close(); err1 != nil && err == nil {
        err = err1
    }
    return err
}

// Unlink 删除文件:unlink()
func Unlink(fileName string) error {
    return os.Remove(fileName)
}

func Touch(fileName string) (*os.File, error) {

    return os.Create(fileName)

}

// FOpen fopen(),flag参考源码go/libexec/src/os/file.go定义的O_XXX常量
func FOpen(fileName string, flag int) (*os.File, error) {
    return os.OpenFile(fileName, flag, 0666)
}

// FReadFile 读取文件-一步到位
func FReadFile(fileName string) ([]byte, error) {
    f, err := FOpen(fileName, os.O_RDONLY)
    if err != nil {
        return nil, err
    }
    defer FClose(f)
    return FRead(f)
}

// FRead 读取文件
func FRead(f *os.File) ([]byte, error) {
    size := 1024
    //必须要预定义长度
    b := make([]byte, size)
    chunks := make([]byte, 0)
    bf := bufio.NewReader(f)
    for {
        n, err := bf.Read(b)
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
        if n == 0 {
            break
        }
        //最后一次读取到的数据，可能小于预定义的大小
        if n < size {
            b = b[:n]
        }
        chunks = append(chunks, b...)
    }
    return chunks, nil
}

// FWrite 写入文件:fwrite()
func FWrite(f *os.File, b []byte) (int, error) {
    return f.Write(b)
}

func FPuts(f *os.File, b []byte) (int, error) {
    return FWrite(f, b)
}

// FClose fclose()
func FClose(f *os.File) error {
    return f.Close()
}

// FExists 判断文件是否存在：file_exists()
func FExists(fileName string) bool {
    _, err := os.Stat(fileName)
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return true
}

func DirName(fileName string) {

}

//获取当前目录
func __DIR__() (string, error) {
    return os.Getwd()
}

type FPathInfo struct {
    DirName   string
    BaseName  string
    Extension string
    FileName  string
    AbsPath   string
}

func PathInfo(fileName string) FPathInfo {
    var data FPathInfo
    data.AbsPath, _ = filepath.Abs(fileName)
    data.DirName = filepath.Dir(fileName)
    data.BaseName = filepath.Base(fileName)
    ext := filepath.Ext(fileName)
    if ext != "" {
        extIndex := strings.LastIndex(ext, ".")
        if extIndex != -1 {
            data.Extension = ext[extIndex+1:]
        }
    }

    return data
}

// IsDir is_dir()
func IsDir(fileName string) (bool, error) {
    fs, err := os.Stat(fileName)
    if err != nil {
        return false, err
    }
    return fs.IsDir(), nil
}

// IsFile is_file()
func IsFile(fileName string) (bool, error) {
    fs, err := os.Stat(fileName)
    if err != nil {
        return false, err
    }
    return fs.IsDir() == false, nil
}
