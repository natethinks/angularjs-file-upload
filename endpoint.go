package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("no file found in request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	io.Copy(&Buf, file)
	contents := Buf.String()

	log.Println(contents)

	Buf.Reset()

	w.WriteHeader(http.StatusOK)
	return
}
