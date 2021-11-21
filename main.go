package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	//"strings"
)

func main() {

	fName:="data.csv"
	file,err:=os.Create(fName)
	if err!=nil {
		log.Fatalf("Failed to c reate file: %q",err)
		return
	}

	defer file.Close()
	writer:=csv.NewWriter(file)
	defer writer.Flush()


	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("#ctl00_MPane_m_198_10561_ctnr_m_198_10561_Panel1>table>tbody>tr", func(e *colly.HTMLElement) {
		var liste_stats []string
		var takim string
		var statistics string
		e.ForEach("td", func(_ int, elem *colly.HTMLElement) {

			takim = elem.ChildText("a")

				statistics= elem.ChildText("span")


			liste_stats=append(liste_stats,takim )
			liste_stats=append(liste_stats, statistics)

			// Print link
		})
			fmt.Println(liste_stats)
			writer.Write(liste_stats)

			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
			c.Visit(e.Request.AbsoluteURL(takim))

	})

	counter:=0
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		if counter != 1 {
			fmt.Println("Visiting", r.URL.String())
			counter++
		}

	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.tff.org/default.aspx?pageID=198")
}