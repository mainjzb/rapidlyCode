package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"golang.org/x/net/html"
)

var HeroSkillUrl = "https://maplestory.fandom.com/wiki/Hero/Skills"

type Table struct {
	XMLName xml.Name `xml:"table"`
	Text    string   `xml:",chardata"`
	Class   string   `xml:"class,attr"`
	Tbody   struct {
		Tr []struct {
			Th []struct {
				Text string `xml:",chardata"`
				A    []struct {
					Text  string `xml:",chardata"`
					Href  string `xml:"href,attr"`
					Title string `xml:"title,attr"`
					Img   struct {
						Text          string `xml:",chardata"`
						Alt           string `xml:"alt,attr"`
						Src           string `xml:"src,attr"`
						Decoding      string `xml:"decoding,attr"`
						Loading       string `xml:"loading,attr"`
						Width         string `xml:"width,attr"`
						Height        string `xml:"height,attr"`
						DataImageName string `xml:"data-image-name,attr"`
						DataImageKey  string `xml:"data-image-key,attr"`
					} `xml:"img"`
				} `xml:"a"`
			} `xml:"th"`
			Td struct {
				Text string `xml:",chardata"`
			} `xml:"td"`
		} `xml:"tr"`
	} `xml:"tbody"`
}

type Mechanic struct {
	Level  string
	Detail string
}

type Skill struct {
	JobAdvancement string

	Name         string
	Type         string
	Icon         string
	RequireLevel string
	IconName     string
	Detail       string
	MaxLevel     string
	Description  string

	MechanicsLevel  string
	MechanicsDetail string
}

func main() {

	// req.DevMode()                     // Treat the package name as a Client, enable development mode
	resp := req.MustGet(HeroSkillUrl) // Treat the package name as a Request, send GET request.
	// body := resp.String()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	// doc.Find("divc .mw-parser-output").Each(func(i int, s *goquery.Selection) {
	// 	// For each item found, get the title
	// 	title := s.Find("a").Text()
	// 	fmt.Printf("Review %d: %s\n", i, title)
	// })

	s := doc.Find(".mw-parser-output").First()
	skills := make([]Skill, 0, 100)

	var jobLevel string
	s.Find(".wikitable,.mw-headline").Each(func(i int, n *goquery.Selection) {
		// if n.Find(".mw-headline") != nil {
		// 	// fmt.Println(n.Find(".mw-headline").Text())
		// }
		// np := n.Parent()
		// NHTML, _ := n.Html()
		var buf bytes.Buffer
		html.Render(&buf, n.Nodes[0])
		fmt.Println(i, buf.String())
		fmt.Println()
		fmt.Println()

		if n.Nodes[0].Data == "span" {
			var htmlBuf bytes.Buffer
			html.Render(&htmlBuf, n.Nodes[0])

			if n.Nodes[0].FirstChild.Data == "a" {
				jobLevel = strings.Trim(n.Nodes[0].FirstChild.NextSibling.Data, " \r\n")
			} else {
				jobLevel = strings.Trim(n.Nodes[0].FirstChild.Data, " \r\n")
			}
			fmt.Println("jobLevel: " + jobLevel)
		}

		if n.Nodes[0].Data == "table" {
			if jobLevel == "Hyper Skills" {
				var t Table
				err := xml.Unmarshal(buf.Bytes(), &t)
				if err != nil {
					panic(err)
				}
				skill := Skill{
					JobAdvancement: jobLevel,
					Name:           strings.Trim(t.Tbody.Tr[0].Th[1].A[1].Text, " \r\n"),
					Type:           strings.Trim(t.Tbody.Tr[0].Th[0].Text, " \r\n"),
					Icon:           t.Tbody.Tr[0].Th[1].A[0].Img.Src,
					IconName:       t.Tbody.Tr[0].Th[1].A[0].Img.DataImageKey,
					Detail:         t.Tbody.Tr[0].Th[1].A[1].Href,
					RequireLevel:   strings.Trim(t.Tbody.Tr[1].Td.Text, " \r\n"),
					MaxLevel:       strings.Trim(t.Tbody.Tr[2].Td.Text, " \r\n"),
					Description:    t.Tbody.Tr[3].Td.Text,

					MechanicsLevel:  t.Tbody.Tr[4].Th[1].Text,
					MechanicsDetail: t.Tbody.Tr[4].Td.Text,
				}
				skills = append(skills, skill)

			} else if jobLevel == "Enhancements" {

			} else {
				var t Table
				err := xml.Unmarshal(buf.Bytes(), &t)
				if err != nil {
					panic(err)
				}
				tr := len(t.Tbody.Tr)
				skill := Skill{
					JobAdvancement: jobLevel,
					Name:           strings.Trim(t.Tbody.Tr[0].Th[1].A[len(t.Tbody.Tr[0].Th[1].A)-1].Text, " \r\n"),
					Type:           strings.Trim(t.Tbody.Tr[0].Th[0].Text, " \r\n"),
					Icon:           strings.Trim(t.Tbody.Tr[0].Th[1].A[0].Href, " \r\n"),
					IconName:       strings.Trim(t.Tbody.Tr[0].Th[1].A[0].Img.DataImageKey, " \r\n"),
					Detail:         strings.Trim(t.Tbody.Tr[0].Th[1].A[1].Href, " \r\n"),
					MaxLevel:       strings.Trim(t.Tbody.Tr[1].Td.Text, " \r\n"),
					Description:    strings.Trim(t.Tbody.Tr[2].Td.Text, " \r\n"),

					MechanicsLevel:  strings.Trim(t.Tbody.Tr[tr-1].Th[5-tr].Text, " \r\n"),
					MechanicsDetail: strings.Trim(t.Tbody.Tr[tr-1].Td.Text, " \r\n"),
				}

				skills = append(skills, skill)

				fmt.Println(skill)
			}
		}
	})
}
