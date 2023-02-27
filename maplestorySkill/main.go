package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
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
	DeleteAll()

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

	s := doc.Find(".mw-parser-output").First()
	skills := make([]Skill, 0, 100)

	var jobLevel string
	s.Find(".wikitable,.mw-headline").Each(func(i int, n *goquery.Selection) {
		if class == "Evan" && i <= 2 {
			return
		}

		if n.Is("span") {
			jobLevel = strings.Trim(n.Text(), " \r\n")
			fmt.Println("jobLevel: " + jobLevel)
		}
		if n.Is("table") {
			Trs := n.Children().Filter("tbody").Children().Filter("tr")
			var Line1, line2, line3, line4 *goquery.Selection
			var requireLevel = 0

			if jobLevel == "Enhancements" {
				return
			}

			if jobLevel == "Hyper Skills" {
				requireLevel, err = strconv.Atoi(strings.Trim(Trs.Eq(1).ChildrenFiltered("th,td").Last().Text(), " \n\r"))
				if err != nil {
					panic(err)
				}
				Line1 = Trs.Eq(0).ChildrenFiltered("th,td")
				line2 = Trs.Eq(2).ChildrenFiltered("th,td")
				line3 = Trs.Eq(3).ChildrenFiltered("th,td")
				line4 = Trs.Eq(Trs.Length() - 1).ChildrenFiltered("th,td")
			} else {
				Line1 = Trs.Eq(0).ChildrenFiltered("th,td")
				line2 = Trs.Eq(1).ChildrenFiltered("th,td")
				line3 = Trs.Eq(2).ChildrenFiltered("th,td")
				line4 = Trs.Eq(Trs.Length() - 1).ChildrenFiltered("th,td")
			}

			icons := make([]string, 0, 5)
			for node := Line1.Eq(1).Children().First(); node.Nodes != nil; node = node.Next() {
				if node.Find("img").Length() > 0 {
					icons = append(icons, node.AttrOr("href", "error"))
				}
			}

			maxLevel, err := strconv.Atoi(strings.Trim(line2.Eq(1).Text(), " \r\n"))
			if err != nil {
				panic(err)
			}

			skill := Skill{
				JobAdvancement: jobLevel,
				Name:           strings.Trim(Line1.Last().Children().Last().Text(), " \r\n"),
				Type:           strings.Trim(Line1.First().Text(), " \r\n"),
				Icons:          icons,
				Detail:         strings.Trim(Line1.Last().Children().Last().AttrOr("href", "error"), " \r\n"),
				MaxLevel:       maxLevel,
				Description:    strings.Trim(line3.Last().Text(), " \r\n"),
				RequireLevel:   requireLevel,

				MechanicsLevel:  strings.Trim(line4.Eq(-2).Text(), " \r\n"),
				MechanicsDetail: strings.Trim(line4.Last().Text(), " \r\n"),
			}
			AddSkill(skill)
			AddClassSkill(class, skill)
			skills = append(skills, skill)
		}
	})
}
