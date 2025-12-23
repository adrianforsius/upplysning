package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Person struct {
	Name     string
	DOB      int
	Street   string
	PostCode string
}

var (
	nameOne   = flag.String("nameOne", "", "help message for flag n")
	genderOne = flag.String("genderOne", "", "help message for flag n")
	nameTwo   = flag.String("nameTwo", "", "help message for flag n")
	genderTwo = flag.String("genderTwo", "", "help message for flag n")
)

func main() {
	flag.Parse()

	male1 := "False"
	female1 := "True"
	if *genderOne == "m" {
		male1 = "True"
		female1 = "False"
	}

	male2 := "False"
	female2 := "True"
	if *genderTwo == "m" {
		male2 = "True"
		female2 = "False"
	}

	url := "https://www.upplysning.se/person/?f=%s&c=vaxholm&county=01&municipality=0187&malegender=%s&femalegender=%s&m=0&sl=detail"
	url1 := fmt.Sprintf(url, *nameOne, male1, female1)
	url2 := fmt.Sprintf(url, *nameTwo, male2, female2)
	// url1 := "https://omni.se"
	// url2 := "https://omni.se"
	// _ = male1
	// _ = male2
	// _ = female1
	// _ = female2

	// url := "https://www.wikipedia.org/"
	launcher := launcher.New().Bin("google-chrome-stable").Headless(false).MustLaunch()

	browser := rod.New().ControlURL(launcher).MustConnect()
	defer browser.Close()

	page := browser.MustPage(url1)

	el, err := page.Element("[href='/logga-in']")
	if err != nil {
		log.Fatal(err)
	}
	el.MustClick()
	page.MustWaitIdle()
	accept, err := page.Element("#onetrust-accept-btn-handler")
	if err != nil {
		log.Fatal(err)
	}
	accept.MustClick()

	user, err := page.Element("[placeholder='E-postadress']")
	if err != nil {
		log.Fatal(err)
	}

	password, err := page.Element("[placeholder='Lösenord']")
	if err != nil {
		log.Fatal(err)
	}
	user.Input("adrianforsius@gmail.com")
	password.Input("Dx!meD9Twnw96fx")

	login, err := page.Element("[value='Logga in']")
	if err != nil {
		log.Fatal(err)
	}
	login.MustClick()
	page.MustElement("[href='/min-sida/']").MustWaitVisible()

	page = page.MustNavigate(url1)
	page.MustWaitIdle()
	page.MustElement(".search-button").MustClick()

	search1 := scrape(page)
	page = page.MustNavigate(url2)
	page.MustElement(".search-button").MustClick()

	page.MustWaitIdle()
	search2 := scrape(page)
	for _, one := range search1 {
		for _, two := range search2 {
			if one.Street == two.Street {
				fmt.Println(one, " ", two)
			}
		}
	}
}

func scrape(page *rod.Page) []Person {
	var persons []Person

	for {
		searchResult, err := page.Elements(".search-result-item")
		if err != nil {
			log.Fatal(err)
		}
		for _, s := range searchResult {
			// el = s.MustElement("div")
			el := s.MustElement("img")
			attr, err := el.Attribute("alt")
			if err != nil {
				log.Println(err)
			}

			txt := s.MustText()

			// fmt.Println(txt)
			parts := strings.Split(txt, "\n")

			dob := 0
			if attr != nil {
				ageParts := strings.Split(*attr, "-")
				dob, err = strconv.Atoi(ageParts[0])
				if err != nil {
					log.Println(err)
				}
			}

			persons = append(persons, Person{
				Name:   strings.TrimSpace(parts[0]),
				DOB:    dob,
				Street: strings.TrimSpace(parts[1]),
			})
		}

		page.MustWaitIdle()
		if !page.MustHas("[rel='next']") {
			break
		}

		el, err := page.Element("[rel='next']")
		if err != nil {
			break
		}
		// fmt.Println(persons)
		el.MustClick()
	}
	return persons
}
