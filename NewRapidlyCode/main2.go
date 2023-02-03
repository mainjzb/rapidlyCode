package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"RapidlyCode/rule"
	"github.com/PuerkitoBio/goquery"
	"github.com/dghubble/trie"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

var r *rule.Rule

func test() {
	t := trie.NewPathTrie()
	t.Put("cat one", 0)
	t.Put("cat two", 2)
	fmt.Println(t)
}

func main() {
	t := trie.NewPathTrieWithConfig(&trie.PathTrieConfig{
		Segmenter: func(path string, start int) (segment string, next int) {
			if len(path) == 0 || start < 0 || start > len(path)-1 {
				return "", -1
			}
			end := strings.IndexRune(path[start+1:], ' ') // next '/' after 0th rune
			if end == -1 {
				return path[start:], -1
			}
			return path[start : start+end+1], start + end + 2
		},
	})

	t.Put("cat one", 0)
	t.Put("cat two", 2)
	v := t.Get("cat two")

	walker := func(key string, value interface{}) error {
		// if(key )
		return nil
	}
	t.Walk(walker)
	fmt.Println(v)

	// read rule
	r = rule.ReadRule("rulefile")
	// rule, err := ReadRule("rule/rule.md")
	// if err != nil {
	// 	panic("read rule error:" + err.Error())
	// }

	doc := ReadSourceFile("a.html")
	html3(doc)
	h, err := doc.Html()
	if err != nil {
		log.Fatalln("save html err: ", err.Error())
	}
	file, _ := os.Create("b.html")
	defer file.Close()
	file.WriteString(h)
	fmt.Println("done")
	// Html2(source, r)
	// Insert(source, rule)
}

// func Insert(source string, rules *rule.Rule) []*html.Node {
// 	isChange := false
// 	findIndex := 0
// 	for {
// 		key, val := rules.Next()
// 		if key == "" && val == "" {
// 			break
// 		}
// 		r, index := getAddRuleIndex(key)
// 		for {
// 			find := strings.Index(source[findIndex:], r)
// 			if find == -1 {
// 				break
// 			}
//
// 			findIndex += find
// 			source = source[:findIndex+index] + val + source[findIndex+index:]
//
// 			isChange = true
// 			findIndex += index + len(val) + len(r)
//
// 			if findIndex > len(source) {
// 				panic(len(source))
// 			}
// 		}
// 	}
// 	if !isChange {
// 		return nil
// 	}
// 	node, err := html.ParseFragment(strings.NewReader(source), &html.Node{
// 		Type:     html.ElementNode,
// 		Data:     "body",
// 		DataAtom: atom.Body,
// 	})
// 	if err != nil {
// 		panic("insert function err: " + err.Error())
// 	}
// 	return node
// }

func getAddRuleIndex(rule string) (string, int) {
	i := strings.Index(rule, "{}")
	return rule[0:i] + rule[i+2:], i
}

func ReadSourceFile(path string) *goquery.Document {
	file, err := os.Open("a.html")
	if err != nil {
		log.Fatalln("can't open source html file err: " + err.Error())
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln("goqury Ducument Err: ", err.Error())
	}
	return doc
}

func Html(r io.Reader) {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			// ...
			return
		}
		// Process the current token.
		tagName, moreAttr := z.TagName()
		fmt.Println("Row:" + string(z.Raw()) + "|")
		fmt.Println(tt.String(), string(tagName), string(z.Text()))
		for moreAttr {
			var key, val []byte
			key, val, moreAttr = z.TagAttr()

			fmt.Print(string(key), ":", string(val), " ")
		}
		fmt.Println()
	}
}

func Html2(r io.Reader, rule *rule.Rule) {
	node, _ := html.Parse(r)

	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode {

		}
		if node.Type == html.TextNode && strings.Trim(node.Data, " \n") != "" {
			// fmt.Println(node.Data)
			//
			// nodes := Insert(node.Data, rule)
			// if nodes != nil {
			// 	for _, n := range nodes {
			// 		// 添加图片域名
			// 		for i := range n.Attr {
			// 			if n.Attr[i].Key == "src" {
			// 				n.Attr[i].Val = "https://bbs.gjfmxd.com/" + n.Attr[i].Val
			// 			}
			// 		}
			//
			// 		node.Parent.InsertBefore(n, node)
			// 	}
			// 	node.Parent.RemoveChild(node)
			// 	node = nodes[len(nodes)-1]
			// }
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(node)
	file, _ := os.Create("b.html")
	defer file.Close()
	html.Render(file, node)

	fmt.Println(node.Type)
	fmt.Println(node.Data)
	fmt.Println(node.Namespace)
	fmt.Println(*node)
}

func html3(doc *goquery.Document) {
	selection := doc.Children()
	for {
		if selection.First().Is("p") {
			break
		}
		selection = selection.Children()
	}

	for i := 0; i < len(selection.Nodes); i++ {
		s := selection.Eq(i)
		handle(s)
	}

}

func handle(s *goquery.Selection) {
	switch {
	case s.Is("p"):
		if len(s.Nodes) > 0 {
			c := s.Nodes[0].FirstChild
			for {
				if c.Type == html.TextNode {
					fmt.Println(c.Data)
					if strings.Contains(c.Data, "New events") {
						c.Data = "123457987985432132123"
					}
				}

				if c == s.Nodes[0].LastChild {
					break
				}
				c = c.NextSibling
			}
		}

		children := s.Children()
		fmt.Println(children)
	case s.Is("ul"):
		children := s.Children()
		for j := 0; j < len(children.Nodes); j++ {
			handle(children.Eq(j))
		}
	case s.Is("li"):
		children := s.Children()
		for j := 0; j < len(children.Nodes); j++ {
			handle(children.Eq(j))
		}
	case s.Is("strong"):
		// 搜索专用名词规则
		dests, ok := r.TerminologyInsert[s.Text()]
		if ok {
			// 插入图片
			source, err := s.Html()
			if err != nil {
				log.Error("can't get html err: " + err.Error())
			}
			newHtml := ""

			// insert content
			newHtml += source[:dests[0].Index] + dests[0].Content
			i := 1
			for ; i < len(dests); i++ {
				newHtml = source[dests[i-1].Index:dests[i].Index] + dests[i].Content
			}
			newHtml += source[dests[i-1].Index:]

			s.SetHtml(newHtml)
		}
	}
}
