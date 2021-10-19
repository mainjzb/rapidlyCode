package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	fileName := "rule.md"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now().Sub(startTime) )
	fmt.Printf("%x", h.Sum(nil))

	ff, err := os.Open("file_info",)
	defer ff.Close()
	b:= []byte(fileName+"="+fmt.Sprintf("%x", h.Sum(nil)))
	ff.Write(b)

}