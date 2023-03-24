package crawler

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

func CrawlForEngines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]
    // Create new Colly Collector
	c:= colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.MaxDepth(100),
	)

	// Handle the Post call with the vin number
	url := "https://www.hollanderparts.com/"
	c.PostMultipart(url,map[string][]byte{
		"hdnVIN":[]byte(vin),
	})
    c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Something went very wrong:",err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited",r.Request.URL)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished",r.Request.URL)
	})
	fmt.Fprintf(w, "You used this vin: %s ", vin)
}
