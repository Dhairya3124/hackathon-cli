package hackathonslist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
    "strconv"
	"time"

	"github.com/common-nighthawk/go-figure"
    "github.com/briandowns/spinner"
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
var year,month,day = time.Now().Date()

func ListofEvents() {
    fmt.Print("\033[H\033[2J") // clears the scren ,source: https://student.cs.uwaterloo.ca/~cs452/terminal.html
    fmt.Println("\033[?25l")   // hide cursor , source: https://student.cs.uwaterloo.ca/~cs452/terminal.html
    banner:= figure.NewFigure("Hackathons List", "", true)
    banner.Print()
    fmt.Println("\n")

    allEvents := make([]Event, 0)
    spinner := spinner.New(spinner.CharSets[27], 100*time.Millisecond)
    spinner.Start()
    spinner.Suffix = "Loading..."
    time.Sleep(4 * time.Second) 
    spinner.Stop()
    fmt.Println("\n")
    // Instantiate default collector
    c := colly.NewCollector()

    // On every a element which has href attribute call callback
    c.OnHTML(".event-wrapper", func(e *colly.HTMLElement) {
        title:= e.ChildText("h3")

        //fmt.Println(title)
        // Timeline:= e.ChildText("p")
        // fmt.Println(Timeline)
        link:= e.ChildAttr("a[href]","href")
        //fmt.Println(link)
        location:= e.ChildText("div.event-location > span:nth-child(1)")
        //fmt.Println(location)
        location2:= e.ChildText("div.event-location > span:nth-child(2)")
        //fmt.Println(location2)
        mode_of_participation:= e.ChildText("div.event-hybrid-notes > span")
        //fmt.Println(mode_of_participation)
        start_date:= e.ChildAttr("div.inner>meta:nth-child(5)", "content")
        //fmt.Println(start_date)
        end_date:= e.ChildAttr("div.inner>meta:nth-child(6)", "content")
        //fmt.Println(end_date)

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
        showJSON()
        // _,yo,_:= "2020-07-03"
        // fmt.Println(yo.Month())

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
        //fmt.Println("Visiting", r.URL.String())
    })

    // Start scraping on https://hackerspaces.org
    c.Visit("https://mlh.io/seasons/2022/events")
}
func writeJSON(data []Event){
    newFile, err:= json.MarshalIndent(data, "", " ")
    if err != nil {
        log.Println("File couldn't be made.!")
        return
    }
    _ = ioutil.WriteFile("events.json", newFile, 0644)

}
func showJSON(){
    currentMonth:= strconv.Itoa(int(month))
 


    jsonFile, err := ioutil.ReadFile("events.json")
    if err != nil {
        log.Println("File couldn't be read.!")
        return
    }
    var data []Event
    err = json.Unmarshal(jsonFile, &data)
    if err != nil {
        log.Println("File couldn't be unmarshalled.!")
        return
    }

    for i:=0;i<len(data);i++{
        if len(data[i].StartDate) >= 7{
            checkmonth,_:=strconv.Atoi(data[i].StartDate[5:7])
            currentMonth,_:=strconv.Atoi(currentMonth)
            if checkmonth == currentMonth{
        
            fmt.Println("Title:",data[i].Name)
            fmt.Println("Location:",data[i].Location)
            fmt.Println("Mode of Particpation:",data[i].ModeofParticpation)
            fmt.Println("Start Date:",data[i].StartDate)
            fmt.Println("End Date:",data[i].EndDate)
            fmt.Println("Registration Link",data[i].Link)
            fmt.Println("\n")    
        }
          
        
    }
        
         
        // fmt.Println(data[i].StartDate)
        // fmt.Println(data[i].Name)
        // fmt.Println(data[i].Link)
    }
 //   fmt.Println(currentMonth)
}