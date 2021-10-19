package main

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	_ "io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	compareHash()
}

func compareHash() {
	// todo:遍历所有文件名字
	var files = []string{
		"rule.md",
	}
	oldHash := ReadFileHash()

	flag := false
	for _, fileName := range files {
		newHash, err := calcFileSha(fileName)
		if err != nil {
			return
		}
		if oldHash[fileName] != newHash {
			flag = true
			// todo: check rule
			oldHash[fileName] = newHash
		}
	}
	if flag {
		SaveFileHash(oldHash)
	}
}

func SaveFileHash(hash map[string]string) {
	f, err := os.OpenFile("hashInfo", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for left, right := range hash {
		f.WriteString(left + "==" + right)
	}
}

func calcFileSha(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func ReadFileHash() map[string]string {
	file, err := os.Open("hashInfo")
	if errors.Is(err, os.ErrNotExist) {
		return make(map[string]string)
	}

	if err != nil {
		log.Fatal(errors.Is(err, os.ErrNotExist))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	fileHash := make(map[string]string)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		content := strings.TrimSpace(scanner.Text())
		if len(content) == 0 {
			continue
		}
		left, right, err := splitTwoEqualSign(content)
		if err != nil {
			log.Panic(err.Error()+" : line ", lineNum)
		}
		fileHash[left] = right
	}

	if err := scanner.Err(); err != nil {
		return nil
	}
	return fileHash
}

func splitTwoEqualSign(content string) (string, string, error) {
	result := strings.Split(content, "==")
	if len(result) < 2 {
		return "", "", errors.New("split EqualSign error: don't exist ==")
	}
	if result[0] == "" {
		return "", "", errors.New("split EqualSign error: left value is null")
	}

	return result[0], result[1], nil
}

func checkRule() {
	// todo:前句覆盖后句? 如何做到
	// todo:<span> 不应该出现

}

func replace() {
	// todo: 默认模式
	// 自带\bcontent\b 效果
	//
}
