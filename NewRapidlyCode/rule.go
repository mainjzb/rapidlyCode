package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Rule struct {
	rules   map[string]string
	order   []string
	current int
}

func (r *Rule) Next() (string, string) {
	if r.current == -1 {
		r.current += 1
		return "", ""
	}
	key := r.order[r.current]
	value := r.rules[key]
	r.current += 1
	if r.current >= len(r.order) {
		r.current = -1
	}
	return key, value
}

func ReadRuleFromDir(path string) *Rule {
	rules := make(map[string]string)

	entry, err := os.ReadDir(path)
	if err != nil {
		panic("readdir err: " + err.Error())
	}

	for _, e := range entry {
		if !e.IsDir() {
			file, err := os.Open(filepath.Join(path, e.Name()))
			if err != nil {
				panic("open file err: " + err.Error())
			}
			scanner := bufio.NewScanner(file)
			lineNum := 0

			for scanner.Scan() {
				lineNum++
				content := strings.TrimSpace(scanner.Text())
				if len(content) == 0 {
					continue
				}
				if strings.HasPrefix(content, "#") {
					continue
				}
				left, right, err := splitTwoEqualSign(content)
				if err != nil {
					log.Panic(err.Error()+" : line ", lineNum)
				}
				rules[left] = right
			}

			if err := scanner.Err(); err != nil {
				panic("scanner err: " + err.Error())
			}
		}
	}

	keys := make([]string, 0, len(rules))
	for k, _ := range rules {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	rule := Rule{
		rules: rules,
		order: keys,
	}

	return &rule
}

func ReadRule(path string) (*Rule, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	rules := make(map[string]string)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		content := strings.TrimSpace(scanner.Text())
		if len(content) == 0 {
			continue
		}
		if strings.HasPrefix(content, "#") {
			continue
		}
		left, right, err := splitTwoEqualSign(content)
		if err != nil {
			log.Panic(err.Error()+" : line ", lineNum)
		}
		rules[left] = right
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	keys := make([]string, 0, len(rules))
	for k, _ := range rules {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	rule := Rule{
		rules: rules,
		order: keys,
	}

	return &rule, nil
}
