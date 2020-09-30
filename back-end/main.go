package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/service-monitor/back-end/timeresponse"
)

type Item struct {
	Id        int       `json:"id"`
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

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

type Items []Item

var item []Item = arrDetails

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
				//for i= 1; i< len(txtlines);i++{
				ss++

				var urls = []string{
					eachline,
				}
				//chh := make(chan string)
				results := asyncHttpGets(urls)
				//	output = append(output, []string{<-chh})
				//for _, value := range output {
				//	fmt.Println(value)
				//}

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
					//	m := makeRange(1, lineCount+1)
					//a := string(i)
					var arrDetails []Item
					//	arrDetails = append(arrDetails, []{<-chh})
					for _, v := range dataa {
						//	for i, line := range lines {
						//	fmt.Println(i, line)
						arrDetails = append(arrDetails, Item{

							Id:        ss,
							Url:       eachline,
							Timeresp:  timeresp,
							Num:       b,
							Timestamp: time.Now(),
							Def:       fmt.Sprintf(c, v),
						})

						//}
					}
					/*lines, err := readLines("urls.txt")
					if err != nil {
						log.Fatalf("readLines: %s", err)
					}
					for i, line := range lines {
						fmt.Println(i, line)
						people = append(people, Person{ID: i})
					}*/
					file, _ = os.Open("urls.txt")
					fileScanner = bufio.NewScanner(file)
					lineCount = 0

					for fileScanner.Scan() {
						lineCount++

					}
					fmt.Println("number of lines:", lineCount)
					//m := makeRange(1, lineCount+1)

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

			//fmt.Println(string(data))
			x = string(data)
			fmt.Println(string(data))

			//defer file.Close()
			time.Sleep(10 * time.Second)

		}

	}()

}

func main() {
	go f()

	/*router := mux.NewRouter()

	router.HandleFunc("/api/stats", Handlercheck).Methods("GET")
	router.HandleFunc("/api/details/{id}", GetPersonEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))*/

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

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	for _, itemm := range item {
		if itemm.Id == id {
			data, err := json.Marshal(itemm)
			if err != nil {
				panic(err)
			}

			//json.NewEncoder(w).Encode(itemm)
			fmt.Fprintf(w, string(data), html.EscapeString(req.URL.Path))
			return
		}
	}
	fmt.Fprintf(w, x, html.EscapeString(req.URL.Path))
	//fmt.Fprintf(w, x, html.EscapeString(req.URL.Path))
	//json.NewEncoder(w).Encode(&Item{})
}

var output = make([][]string, 0)

func MakeRequest(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("url: %s, err: %s ", url, err)
	} //else {
	//	ch <- fmt.Sprintf("url: %s, status: %s ", url, resp.Status) // put response into a channel
	defer resp.Body.Close()
	//}
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			/*client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}*/
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

