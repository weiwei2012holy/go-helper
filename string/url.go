package string

import (
    "encoding/base64"
    url2 "net/url"
)

//Base64Encode  base64_encode
func Base64Encode(src []byte) string {
    return base64.StdEncoding.EncodeToString(src)
}

//Base64Decode base64_decode
func Base64Decode(src string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(src)
}

//ParseUrl parse_url 解析 URL，返回其组成部分
func ParseUrl(url string) (*url2.URL, error) {
    return url2.Parse(url)
}

// HttpBuildQuery http_build_query — 生成 URL-encode 之后的请求字符串
func HttpBuildQuery(queryData url2.Values) string {
    return queryData.Encode()
}

//UrlDecode urldecode — 解码已编码的 URL 字符串
func UrlDecode(str string) (string, error) {
    return url2.QueryUnescape(str)
}

// UrlEncode urlencode — 编码 URL 字符串
func UrlEncode(str string) string {
    return url2.QueryEscape(str)
}
