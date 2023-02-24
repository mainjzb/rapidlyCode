package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"golang.org/x/net/html"
)

var HeroSkillUrl = "https://maplestory.fandom.com/wiki/Hero/Skills"

var skillUrl = []struct {
	class string
	url   string
}{
	{"Hero", "https://maplestory.fandom.com/wiki/Hero/Skills"},
	{"Paladin", "https://maplestory.fandom.com/wiki/Paladin/Skills"},
	{"Dark Knight", "https://maplestory.fandom.com/wiki/Dark_Knight/Skills"},
	{"Magician (Fire, Poison)", "https://maplestory.fandom.com/wiki/Magician_(Fire,_Poison)/Skills"},
	{"Magician_(Ice, Lightning)", "https://maplestory.fandom.com/wiki/Magician_(Ice,_Lightning)/Skills"},
	{"Bishop", "https://maplestory.fandom.com/wiki/Bishop/Skills"},
	{"Bowmaster", "https://maplestory.fandom.com/wiki/Bowmaster/Skills"},
	{"Marksman", "https://maplestory.fandom.com/wiki/Marksman/Skills"},
	{"Pathfinder", "https://maplestory.fandom.com/wiki/Pathfinder/Skills"},
	{"Night Lord", "https://maplestory.fandom.com/wiki/Night_Lord/Skills"},
	{"Shadower", "https://maplestory.fandom.com/wiki/Shadower/Skills"},
	{"Dual Blade", "https://maplestory.fandom.com/wiki/Dual_Blade/Skills"},
	{"Buccaneer", "https://maplestory.fandom.com/wiki/Buccaneer/Skills"},
	{"Corsair", "https://maplestory.fandom.com/wiki/Corsair/Skills"},
	{"Cannoneer", "https://maplestory.fandom.com/wiki/Cannoneer/Skills"},
	{"Jett", "https://maplestory.fandom.com/wiki/Jett/Skills"},
	{"Dawn Warrior", "https://maplestory.fandom.com/wiki/Dawn_Warrior/Skills"},
	{"Blaze Wizard", "https://maplestory.fandom.com/wiki/Blaze_Wizard/Skills"},
	{"Wind Archer", "https://maplestory.fandom.com/wiki/Wind_Archer/Skills"},
	{"Night Walker", "https://maplestory.fandom.com/wiki/Night_Walker/Skills"},
	{"Thunder Breaker", "https://maplestory.fandom.com/wiki/Thunder_Breaker/Skills"},
	{"Mihile", "https://maplestory.fandom.com/wiki/Mihile/Skills"},
	{"Aran", "https://maplestory.fandom.com/wiki/Aran/Skills"},
	{"Evan", "https://maplestory.fandom.com/wiki/Evan/Skills"},
	{"Mercedes", "https://maplestory.fandom.com/wiki/Mercedes/Skills"},
	{"Phantom", "https://maplestory.fandom.com/wiki/Phantom/Skills"},
	{"Luminous", "https://maplestory.fandom.com/wiki/Luminous/Skills"},
	{"Shade", "https://maplestory.fandom.com/wiki/Shade/Skills"},
	{"Blaster", "https://maplestory.fandom.com/wiki/Blaster/Skills"},
	{"Battle Mage", "https://maplestory.fandom.com/wiki/Battle_Mage/Skills"},
	{"Wild Hunter", "https://maplestory.fandom.com/wiki/Wild_Hunter/Skills"},
	{"Mechanic", "https://maplestory.fandom.com/wiki/Mechanic/Skills"},
	{"Demon Slayer", "https://maplestory.fandom.com/wiki/Demon_Slayer/Skills"},
	{"Demon Avenger", "https://maplestory.fandom.com/wiki/Demon_Avenger/Skills"},
	{"Xenon", "https://maplestory.fandom.com/wiki/Xenon/Skills"},
	{"Kaiser", "https://maplestory.fandom.com/wiki/Kaiser/Skills"},
	{"Kain", "https://maplestory.fandom.com/wiki/Kain/Skills"},
	{"Cadena", "https://maplestory.fandom.com/wiki/Cadena/Skills"},
	{"Angelic Buster", "https://maplestory.fandom.com/wiki/Angelic_Buster/Skills"},
	{"Adele", "https://maplestory.fandom.com/wiki/Adele/Skills"},
	{"Illium", "https://maplestory.fandom.com/wiki/Illium/Skills"},
	//	{"Khali", "https://maplestory.fandom.com/wiki/Khali/Skills"},
	{"Ark", "https://maplestory.fandom.com/wiki/Ark/Skills"},
	{"Hayato", "https://maplestory.fandom.com/wiki/Hayato/Skills"},
	{"Kanna", "https://maplestory.fandom.com/wiki/Kanna/Skills"},
	{"Lara", "https://maplestory.fandom.com/wiki/Lara/Skills"},
	{"Hoyoung", "https://maplestory.fandom.com/wiki/Hoyoung/Skills"},
	{"Beast Tamer", "https://maplestory.fandom.com/wiki/Beast_Tamer/Skills"},
	{"Zero", "https://maplestory.fandom.com/wiki/Zero/Skills"},
	{"Kinesis", "https://maplestory.fandom.com/wiki/Kinesis/Skills"},

	// {"Pink Bean", "https://maplestory.fandom.com/wiki/Pink_Bean/Skills"},
	// {"Yeti", "https://maplestory.fandom.com/wiki/Yeti/Skills"},
}

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

func main() {
	read("Evan", "https://maplestory.fandom.com/wiki/Evan/Skills")
	return
	for _, url := range skillUrl {
		read(url.class, url.url)
	}
}

func read(class, url string) {

	// req.DevMode()                     // Treat the package name as a Client, enable development mode
	resp := req.MustGet(url) // Treat the package name as a Request, send GET request.
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
		if class == "Evan" && i <= 2 {
			return
		}

		var buf bytes.Buffer
		html.Render(&buf, n.Nodes[0])
		// fmt.Println(i, buf.String())
		// fmt.Println()
		// fmt.Println()

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
			var t Table
			err := xml.Unmarshal(buf.Bytes(), &t)
			if err != nil {
				panic(err)
			}

			icons := make([]string, 0, 5)
			detail := ""
			for _, aTag := range t.Tbody.Tr[0].Th[1].A {
				if aTag.Img.Src != "" {
					icons = append(icons, aTag.Href)
				} else {
					detail = aTag.Href
				}
			}

			if jobLevel == "Hyper Skills" {
				maxLevel, err := strconv.Atoi(strings.Trim(t.Tbody.Tr[2].Td.Text, " \r\n"))
				if err != nil {
					panic(err)
				}
				requireLevel, err := strconv.Atoi(strings.Trim(t.Tbody.Tr[1].Td.Text, " \r\n"))
				if err != nil {
					panic(err)
				}

				skill := Skill{
					JobAdvancement: jobLevel,
					Name:           strings.Trim(t.Tbody.Tr[0].Th[1].A[len(t.Tbody.Tr[0].Th[1].A)-1].Text, " \r\n"),
					Type:           strings.Trim(t.Tbody.Tr[0].Th[0].Text, " \r\n"),
					Icons:          icons,
					IconName:       strings.Trim(t.Tbody.Tr[0].Th[1].A[0].Img.DataImageKey, " \r\n"),
					Detail:         strings.Trim(detail, " \r\n"),
					RequireLevel:   requireLevel,
					MaxLevel:       maxLevel,
					Description:    strings.Trim(t.Tbody.Tr[3].Td.Text, " \r\n"),

					MechanicsLevel:  strings.Trim(t.Tbody.Tr[4].Th[1].Text, " \r\n"),
					MechanicsDetail: strings.Trim(t.Tbody.Tr[4].Td.Text, " \r\n"),
				}
				AddSkill(skill)
				AddClassSkill(class, skill)
				skills = append(skills, skill)

			} else if jobLevel == "Enhancements" {
			} else {
				tr := len(t.Tbody.Tr)
				maxLevel, err := strconv.Atoi(strings.Trim(t.Tbody.Tr[1].Td.Text, " \r\n"))
				if err != nil {
					panic(err)
				}

				skill := Skill{
					JobAdvancement: jobLevel,
					Name:           strings.Trim(t.Tbody.Tr[0].Th[1].A[len(t.Tbody.Tr[0].Th[1].A)-1].Text, " \r\n"),
					Type:           strings.Trim(t.Tbody.Tr[0].Th[0].Text, " \r\n"),
					Icons:          icons,
					IconName:       strings.Trim(t.Tbody.Tr[0].Th[1].A[0].Img.DataImageKey, " \r\n"),
					Detail:         strings.Trim(detail, " \r\n"),
					MaxLevel:       maxLevel,
					Description:    strings.Trim(t.Tbody.Tr[2].Td.Text, " \r\n"),

					MechanicsLevel:  strings.Trim(t.Tbody.Tr[tr-1].Th[5-tr].Text, " \r\n"),
					MechanicsDetail: strings.Trim(t.Tbody.Tr[tr-1].Td.Text, " \r\n"),
				}
				AddSkill(skill)
				AddClassSkill(class, skill)
				skills = append(skills, skill)

				fmt.Println(skill)

			}
		}
	})
}
