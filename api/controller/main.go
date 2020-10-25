package controller

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

// This is the root directory of uploaded files
var base = "/home/mehrdadep/example"

func Upload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}

	defer src.Close()
	n := fmt.Sprintf("%d - %s", time.Now().UTC().Unix(), file.Filename)
	dst := fmt.Sprintf("%s/%s", base, n)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}

	defer out.Close()

	_, err = io.Copy(out, src)
	return n, err
}
