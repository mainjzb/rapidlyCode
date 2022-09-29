package rule

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Rule struct {
	TerminologyInsert map[string][]dest // 大写专用名词规则 （插入） map[source]dest
	//	Index             map[string]int    // 插入位置 map[source]index
}

type dest struct {
	Index   int
	Content string
}

func ReadRule(path string) *Rule {
	rules := make(map[string][]dest)

	entry, err := os.ReadDir(path)
	if err != nil {
		panic("readdir err: " + err.Error())
	}

	// read every rule file
	for _, e := range entry {
		if !e.IsDir() {
			file, err := os.Open(filepath.Join(path, e.Name()))
			if err != nil {
				panic("open file err: " + err.Error())
			}

			// read each line
			scanner := bufio.NewScanner(file)
			lineNum := 0
			for scanner.Scan() {
				lineNum++
				content := strings.TrimSpace(scanner.Text())
				if len(content) == 0 {
					continue
				}
				// skip # comment
				if strings.HasPrefix(content, "#") {
					continue
				}

				left, right, err := splitTwoEqualSign(content)
				if err != nil {
					log.Panic(err.Error()+" : line ", lineNum, "content: ", content)
				}

				// parse right
				rightList := strings.Split(right, ";")

				// parse left and create `dest`
				d := make([]dest, 0, 1)
				for i := 0; ; i++ {
					index := strings.Index(left, "{}")
					if index == -1 {
						break
					}
					if i > len(d) {
						panic("left {} can't find more right contents")
					}

					left = left[0:index] + left[index+2:]
					d = append(d, dest{
						Index:   index,
						Content: rightList[i],
					})
				}

				rules[left] = d
			}

			if err := scanner.Err(); err != nil {
				panic("scanner err: " + err.Error())
			}
		}
	}

	rule := Rule{
		TerminologyInsert: rules,
	}

	return &rule
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
