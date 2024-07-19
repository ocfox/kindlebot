package main

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"
)

type File struct {
	Name string
	Link string
}

type Attachment struct {
	Filename string
	MIME     string
	Data     []byte
}

func GetMIME(name string) string {
	if strings.HasSuffix(name, ".pdf") {
		return "application/pdf"
	}
	if strings.HasSuffix(name, ".epub") {
		return "application/epub+zip"
	}
	return "application/octet-stream"
}

func DownloadFile(file File) Attachment {
	resp, err := http.Get(file.Link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	println("downloaded file")
	if err != nil {
		panic(err)
	}

	encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encodedData, data)

	return Attachment{
		Filename: file.Name,
		MIME:     GetMIME(file.Name),
		Data:     encodedData,
	}
}
