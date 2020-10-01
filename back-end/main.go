package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/service-monitor/back-end/timeresponse"
)

//Item is
type Item struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Timeresp  Duration  `json:"timeresponse"`
	Num       int       `json:"code"`
	Timestamp time.Time `json:"time"`
	//Def       string    `json:"status"`
	Def string `json:"status"`
}

//Duration is
type Duration struct {
	time.Duration
}

//MarshalJSON is
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

//HttpResponse is
type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

//Items is
type Items []Item

//var item []Item = arrDetails
type User struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Timeresp  Duration  `json:"timeresponse"`
	Num       int       `json:"code"`
	Timestamp time.Time `json:"time"`
	//Def       string    `json:"status"`
	Def string `json:"status"`
}

var x string
var xx string
var ii string
var dataa = [1]int{}
var a string
var b int
var c string
var y string
var i int
var ss int

//var data []Item
var arrDetails []Item
var onearrDetails []Item
var sa []Item
var s []Item

func f() {
	go func() {
		//var s []Item
		//	cc := 0
		for {
			file, err := os.Open("urls.txt")
			if err != nil {
				panic(err)
			}
			//if strings.Contains(), "http://") || strings.Contains(file, "http://") {
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var txtlines []string

			for scanner.Scan() {
				txtlines = append(txtlines, scanner.Text())
			}

			file.Close()

			ss = 0

			for _, eachline := range txtlines {
				//for _, eachline := range status {

				//for i= 1; i< len(txtlines);i++{
				ss++
				if strings.Contains(eachline, "http://") || strings.Contains(eachline, "https://") {
					var urls = []string{
						eachline,
					}
					//chh := make(chan string)
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
						//	m := makeRange(1, lineCount+1)
						//a := string(i)
						var arrDetails []Item
						//	arrDetails = append(arrDetails, []{<-chh})
						//	counting := 0
						for _, v := range dataa {
							//	for i, line := range lines {
							//	fmt.Println(i, line)
							onearrDetails = append(onearrDetails, Item{

								Id:        ss,
								Url:       eachline,
								Timeresp:  timeresp,
								Num:       b,
								Timestamp: time.Now(),
								Def:       fmt.Sprintf(c, v),
							})
							//						counting++
							//}
						}

						if len(onearrDetails) == ss {
							first := onearrDetails[0]
							fmt.Println(first)

							onearrDetails = onearrDetails[1:]
						}

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

						sa = append(sa, arrDetails...)

						if len(sa) == (lineCount*1)+1 {
							first := sa[0]
							fmt.Println(first)

							sa = sa[1:]
						}
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

			dat, err := json.Marshal(sa)
			if err != nil {
				panic(err)
			}

			//fmt.Println(string(data))
			xx = string(dat)
			//fmt.Println(string(data))

			//defer file.Close()
			time.Sleep(10 * time.Second)

			//defer file.Close()
			//time.Sleep(eachline.Time * time.Second)
		}

	}()

}

func main() {
	go f()

	router := mux.NewRouter()

	router.HandleFunc("/api/stats", Handlercheck).Methods("GET")
	router.HandleFunc("/api/urls/{id}", GetItemEndpoint).Methods("GET")
	//router.HandleFunc("/api/url/{id}", GetoneItemEndpoint).Methods("GET")
	router.HandleFunc("/api/urls", createUser).Methods("POST")
	router.HandleFunc("/api/urls/{id}", DeleteItemEndpoint).Methods("DELETE")
	//	router.HandleFunc("/api/url/{id}", CreateItemEndpoint).Methods("PUT")
	router.HandleFunc("/api/urls/{id}", updateItem).Methods("PUT")
	log.Fatal(http.ListenAndServe(":10000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	/*	router.Use(mux.CORSMethodMiddleware(router))

		http.ListenAndServe(":8080", router)*/
	//log.Fatal(http.ListenAndServe(":8080", router))

	/*mux := http.NewServeMux()
	mux.HandleFunc("/api/url", Handlercheck)
	mux.HandleFunc("/api/url/1", GetItemEndpoint)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":10000", handler)*/

}

var itemm Item

//Handlercheck is

func Handlercheck(w http.ResponseWriter, r *http.Request) {
	//bodyBytes, _ := ioutil.ReadAll(r.x)
	//w.Write(bodyBytes)

	fmt.Fprintf(w, x, html.EscapeString(r.URL.Path))

}
func GetoneItemEndpoint(w http.ResponseWriter, req *http.Request) {

	json.NewEncoder(w).Encode(onearrDetails)
}

//GetItemEndpoint is
func GetItemEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	//am = append(am, sa...)
	for _, itemm = range onearrDetails {
		//	am = append(am, sa...)
		if itemm.Id == id {
			profiless := []Item{
				{itemm.Id, itemm.Url, itemm.Timeresp, itemm.Num, itemm.Timestamp, itemm.Def},
				//s{},
				//s{},
			}

			am = append(am, profiless...)
			//return

		}

	}
	json.NewEncoder(w).Encode(am)
	time.Sleep(10 * time.Second)

}

type Teststruct struct {
	Test string
}

func createUser(w http.ResponseWriter, r *http.Request) {
	//	var u string
	decoder := json.NewDecoder(r.Body)

	var t Teststruct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}
	//	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//	}

	//	post, _ := strconv.Atoi(u)
	f, err := os.OpenFile("urls.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	file, _ := ioutil.ReadFile("urls.txt")
	s := strings.TrimSpace(string(file))
	fmt.Println(s)
	if strings.Contains(t.Test, "http://") {
		//	_ = ioutil.WriteFile("urls.txt", decoder, 0644)
		_, _ = f.WriteString("\n")
		if _, err := f.WriteString(t.Test); err != nil {
			log.Println(err)
		}
	}

}
func DeleteItemEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	//x := id - 1
	for _, item := range onearrDetails {
		if item.Id == id {
			//arrDetails = append(arrDetails[:index], arrDetails[index+1:]...)
			//break

			path := "urls.txt"
			x := id - 1
			removeLine(path, x)
			data, err := ioutil.ReadFile("urls.txt")
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			s := strings.TrimSpace(string(data))
			fmt.Println(s)

			return
		}
	}
	//json.NewEncoder(w).Encode(onearrDetails)
	///fmt.Fprintf(w, x, html.EscapeString(req.URL.Path))
}
func removeLine(path string, lineNumber int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if strings.Contains(string(file), "http://") {
		info, _ := os.Stat(path)
		mode := info.Mode()

		array := strings.Split(string(file), "\n")
		array = append(array[:lineNumber], array[lineNumber+1:]...)
		ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
	}
}

type Puturlstruct struct {
	Puturl string
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	//inputOrderID := params["orderId"]
	for _, order := range onearrDetails {
		if order.Id == id {

			path := "urls.txt"
			x := id - 1
			removeLine(path, x)
			decoder := json.NewDecoder(r.Body)

			var t Puturlstruct
			err := decoder.Decode(&t)

			if err != nil {
				panic(err)
			}

			f, err := os.OpenFile("urls.txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if strings.Contains(t.Puturl, "http://") || strings.Contains(t.Puturl, "https://") {
				//	_ = ioutil.WriteFile("urls.txt", decoder, 0644)
				//_, _ = f.WriteString(" \n")
				file, _ := ioutil.ReadFile("urls.txt")
				s := strings.TrimSpace(string(file))
				fmt.Println(s)
				_ = InsertStringToFile("urls.txt", "\n", x)
				err := InsertStringToFile("urls.txt", t.Puturl, x)
				file, _ = ioutil.ReadFile("urls.txt")
				s = strings.TrimSpace(string(file))
				fmt.Println(s)
				if err != nil {
					panic(err)
				}

				return
			}

		}
	}
}

func CreateItemEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var array Item
	//_ = json.NewDecoder(req.Body).Decode(&array)
	id, _ := strconv.Atoi(params["id"])
	array.Id = id
	onearrDetails = append(onearrDetails, array)
	//json.NewEncoder(w).Encode(onearrDetails)
	decoder := json.NewDecoder(req.Body)

	var t Teststruct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("urls.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if strings.Contains(t.Test, "http://") {
		//	_ = ioutil.WriteFile("urls.txt", decoder, 0644)
		_, _ = f.WriteString("\n")
		if _, err := f.WriteString(t.Test); err != nil {
			log.Println(err)
		}
	}

}
func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

var am []Item

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	counter := 0
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				//panic(err)
				//fmt.Println(err)
				return

			}
			// handle the error
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			req = req.WithContext(ctx)
			resp, err := client.Do(req)
			//resp, err := http.Get(url)

			if err != nil {
				//panic(err)
				fmt.Println(err)

			}
			counter++
			ch <- &HttpResponse{url, resp, nil}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Println("h", r)
			if r.response != nil {

				fmt.Printf("%s was fetched\n", r.url)

				responses = append(responses, r)
				//	if len(responses) == len(urls) {
				return responses

			}
			if counter == len(urls) {
				fmt.Println("a", len(urls), responses)
				return responses

			}
		//	}
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
