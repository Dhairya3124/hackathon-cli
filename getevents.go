package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "log"

	"github.com/gocolly/colly"
	//    "strings"
)
type Event struct {
    Name string `json:"name"`
    Location string `json:"location"`
    ModeofParticpation string `json:"mode_of_participation"`
    StartDate string `json:"start_date"`
    EndDate string `json:"end_date"`
    Link string `json:"link"`

} 

func main() {
    allEvents := make([]Event, 0)
    // Instantiate default collector
    c := colly.NewCollector()

    // On every a element which has href attribute call callback
    c.OnHTML(".event-wrapper", func(e *colly.HTMLElement) {
        title:= e.ChildText("h3")

        fmt.Println(title)
        // Timeline:= e.ChildText("p")
        // fmt.Println(Timeline)
        link:= e.ChildAttr("a[href]","href")
        fmt.Println(link)
        location:= e.ChildText("div.event-location > span:nth-child(1)")
        fmt.Println(location)
        location2:= e.ChildText("div.event-location > span:nth-child(2)")
        fmt.Println(location2)
        mode_of_participation:= e.ChildText("div.event-hybrid-notes > span")
        fmt.Println(mode_of_participation)
        start_date:= e.ChildAttr("div.inner>meta:nth-child(5)", "content")
        fmt.Println(start_date)
        end_date:= e.ChildAttr("div.inner>meta:nth-child(6)", "content")
        fmt.Println(end_date)

        Events:= Event{
            Name: title,
            Location: location + "," + location2,
            ModeofParticpation: mode_of_participation,
            StartDate: start_date,
            EndDate: end_date,
            Link: link,
        }
        allEvents = append(allEvents, Events)
        writeJSON(allEvents)


	// 	mystring:=e.ChildText("a")
    //     split:=strings.Split(mystring," ")
    //   //  split2:=strings.Split(split[1]," ")
    //     //fmt.Println(split[0])
    //     //fmt.Println(split)
    //     for i:=0;i<len(split);i++{ 
    //         fmt.Println(strings.TrimSpace(split[i]))
    //     }

	})

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    // Start scraping on https://hackerspaces.org
    c.Visit("https://mlh.io/seasons/2021/events")
}
func writeJSON(data []Event){
    newFile, err:= json.MarshalIndent(data, "", " ")
    if err != nil {
        log.Println("File couldn't be made.!")
        return
    }
    _ = ioutil.WriteFile("events.json", newFile, 0644)

}