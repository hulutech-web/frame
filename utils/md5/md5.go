package md5

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func CreateMd5(params string) string {
	w := md5.New()
	_, _ = io.WriteString(w, params)
	return fmt.Sprintf("%x", w.Sum(nil))
}

//先base64，然后MD5
func Base64Md5(params string) string {
	return CreateMd5(base64.StdEncoding.EncodeToString([]byte(params)))
}
