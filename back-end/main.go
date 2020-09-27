package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"time"

	"github.com/service-monitor/back-end/timeresponse"
	"github.com/rs/cors"
)

type Item struct {
	Url       string    `json:"url"`
	Timeresp  Duration  `json:"timeresponse"`
	Num       int       `json:"code"`
	Timestamp time.Time `json:"time"`
	//Def       string    `json:"status"`
	Def string `json:"status"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

/*var urls = []string{
	eachline,
}*/

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type Items []Item

var x string

var dataa = []int{0}
var a string
var b int
var c string

func main() {

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

			for i, eachline := range txtlines {
				var urls = []string{
					eachline,
				}
				//fmt.Println(eachline)
				//chh := make(chan string)
				results := asyncHttpGets(urls)
				//	dataa = append(dataa, []string{<-ch})
				for _, result := range results {
					fmt.Printf("%s status: %s\n", result.url, result.response.Status)
					a = result.url
					b = result.response.StatusCode
					c = result.response.Status

					/*start := time.Now()
					// create a channel of strings:
					timestamp := make(chan string)
					statuscode := make(chan string)
					urls := []string{
						eachline,
					}

					// Call `fetch` in a new goroutine for each URL in `urls`:
					for _, url := range urls {
						go fetch(url, timestamp)
						go fetchcode(url, statuscode)
					}

					// Receive and print each string sent to the `ch` channel from `fetch`:
					for range urls {
						fmt.Println(<-timestamp)
						fmt.Println(<-statuscode)
					}

					//	a = <-timestamp
					//	b = <-statuscode

					// Print the total seconds spent in `main`
					// The individual request response times reported by `fetch` equal a sum greater than the total
					// seconds spent in `main`, thus illustrating that the `fetch` requests occurred concurrently.
					fmt.Printf("Total time: %.2fs\n", time.Since(start).Seconds())
					*/
					//go urlcheck.Urlcheck(eachline)

					timeresp := Duration{timeresponse.Getresptime(eachline)}

					var arrDetails []Item
					//	arrDetails = append(arrDetails, []{<-chh})
					for _, v := range dataa {

						arrDetails = append(arrDetails, Item{
							//Count:    fmt.Sprintf(b, v),
							Url: eachline,
							//Timeresp: Duration{timeresponse.Getresptime(eachline)},
							Timeresp: timeresp,
							Num:      b,
							//Def:      fmt.Sprintf(b, v),
							Timestamp: time.Now(),
							Def:       fmt.Sprintf(c, v),
						})

					}

					s = append(s, arrDetails...)
					if len(s) == i*11 {
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
			//fmt.Println(string(data))
			x = string(data)
			fmt.Println(string(data))

			defer file.Close()
			time.Sleep(5 * time.Second)

		}

	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/api", Handlercheck)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":10000", handler)

}

func Handlercheck(w http.ResponseWriter, r *http.Request) {
	//bodyBytes, _ := ioutil.ReadAll(r.x)
	//w.Write(bodyBytes)
	fmt.Fprintf(w, x, html.EscapeString(r.URL.Path))
}
func fetch(url string, timestamp chan<- string) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		// Send an error to the `ch` channel if one is encountered:

		timestamp <- fmt.Sprint(err)
		return
	}

	//secs := time.Since(start).Seconds()
	// Send a summary string to the `ch` channel containing the URL, its request response time, and its HTTP status code
	timestamp <- fmt.Sprintf("%s", start)

}
func fetchcode(url string, statuscode chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// Send an error to the `ch` channel if one is encountered:

		statuscode <- fmt.Sprint(err)
		return
	}
	//secs := time.Since(start).Seconds()
	// Send a summary string to the `ch` channel containing the URL, its request response time, and its HTTP status code

	statuscode <- fmt.Sprintf("%d", resp.StatusCode)
}
func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)

			resp, _ := http.Get(url)

			/*if err != nil {
				chh <- fmt.Sprintf("url: %s, err: %s ", url, err)
			} else {
				chh <- fmt.Sprintf("url: %s, status: %s ", url, resp.Status) // put response into a channel
				defer resp.Body.Close()

			}*/

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
