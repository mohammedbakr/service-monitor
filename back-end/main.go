package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
	"github.com/service-monitor/back-end/timeresponse"
)

type Person struct {
	ID  int    `json:"id,omitempty"`
	Url string `json:"firstname,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

type Item struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Timeresp  Duration  `json:"timeresponse"`
	Num       int       `json:"code"`
	Timestamp time.Time `json:"time"`
	Def string `json:"status"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type Items []Item

var x string
var ii string
var dataa = [1]int{}
var a string
var b int
var c string
var y string
var i int
var ss int
var data []Item
var arrDetails []Item

func f() {
	go func() {
		var s []Item
		for {
			file, err := os.Open("urls.txt")
			if err != nil {
				panic(err)
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var txtlines []string

			for scanner.Scan() {
				txtlines = append(txtlines, scanner.Text())
			}

			file.Close()
			ss = 0
			for _, eachline := range txtlines {
				ss++
				var urls = []string{
					eachline,
				}
				results := asyncHttpGets(urls)
				for _, result := range results {
					fmt.Printf("%s status: %s\n", result.url, result.response.Status)
					a = result.url
					b = result.response.StatusCode
					c = result.response.Status

					timeresp := Duration{timeresponse.Getresptime(eachline)}
					file, _ := os.Open("urls.txt")
					fileScanner := bufio.NewScanner(file)
					lineCount := 0

					for fileScanner.Scan() {
						lineCount++

					}
					fmt.Println("number of lines:", lineCount)
					var arrDetails []Item
					for _, v := range dataa {
						arrDetails = append(arrDetails, Item{
							Id:  ss,
							Url: eachline,
							Timeresp: timeresp,
							Num:      b,
							Timestamp: time.Now(),
							Def:       fmt.Sprintf(c, v),
						})
					}
					file, _ = os.Open("urls.txt")
					fileScanner = bufio.NewScanner(file)
					lineCount = 0

					for fileScanner.Scan() {
						lineCount++

					}
					fmt.Println("number of lines:", lineCount)
					s = append(s, arrDetails...)
					if len(s) == (lineCount*1)+1 {
						first := s[0]
						fmt.Println(first)

						s = s[1:]
					}
				}
			}
			data, err := json.Marshal(s)
			if err != nil {
				panic(err)
			}

			x = string(data)
			fmt.Println(string(data))
			time.Sleep(5 * time.Second)

		}

	}()

}
func main() {
	go f()
		mux := http.NewServeMux()
		mux.HandleFunc("/api", Handlercheck)

		handler := cors.Default().Handler(mux)
		http.ListenAndServe(":10000", handler)
}

func Handlercheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x, html.EscapeString(r.URL.Path))
}
func fetch(url string, timestamp chan<- string) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		timestamp <- fmt.Sprint(err)
		return
	}

	timestamp <- fmt.Sprintf("%s", start)

}
func fetchcode(url string, statuscode chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		statuscode <- fmt.Sprint(err)
		return
	}

	statuscode <- fmt.Sprintf("%d", resp.StatusCode)
}
func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls))
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)

			resp, _ := http.Get(url)
			ch <- &HttpResponse{url, resp, nil}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		default:
			fmt.Printf(".")
			time.Sleep(5e7)
		}
	}
	return responses

}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
