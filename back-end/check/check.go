package check

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" //database

	"github.com/gorilla/mux"
	"github.com/k8-proxy/service-monitor/back-end/timeresponse"
)

//Post is struct of database
type Post struct {
	Id     int
	TheUrl string
	Time   time.Duration
}

//Item is struct of returned variables
type Item struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Timeresp  Duration  `json:"timeresponse"`
	Num       int       `json:"code"`
	Timestamp time.Time `json:"time"`
	//Def       string    `json:"status"`
	Def string `json:"status"`
}

//Duration is data type for time
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

//ErrorCheck is
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// F is function of healthcheck
func F() {
	go func() {

		for {

			ss = 0
			db, err := sql.Open("mysql", "root@(127.0.0.1:3306)/healthcheck")

			// if there is an error opening the connection, handle it
			if err != nil {
				panic(err.Error())
			}
			defer db.Close()

			rows, e := db.Query("select * from url_info")

			var post = Post{}

			for rows.Next() {
				e = rows.Scan(&post.Id, &post.TheUrl, &post.Time)
				ErrorCheck(e)
				fmt.Println(post.TheUrl)
				fmt.Println(post.Time)
				//}
				ss++
				if strings.Contains(post.TheUrl, "http://") || strings.Contains(post.TheUrl, "https://") {
					var urls = []string{
						post.TheUrl,
					}
					//chh := make(chan string)
					results := asyncHttpGets(urls)

					for _, result := range results {
						fmt.Printf("%s status: %s\n", result.url, result.response.Status)
						a = result.url
						b = result.response.StatusCode
						c = result.response.Status

						timeresp := Duration{timeresponse.Getresptime(post.TheUrl)}

						var arrDetails []Item
						//	arrDetails = append(arrDetails, []{<-chh})
						//	counting := 0
						for _, v := range dataa {
							//	for i, line := range lines {
							//	fmt.Println(i, line)
							onearrDetails = append(onearrDetails, Item{

								Id:        ss,
								Url:       post.TheUrl,
								Timeresp:  timeresp,
								Num:       b,
								Timestamp: time.Now(),
								Def:       fmt.Sprintf(c, v),
							})
							//						counting++
							//}
						}

						/*	if len(onearrDetails) == ss {
								first := onearrDetails[0]
								fmt.Println(first)

								onearrDetails = onearrDetails[1:]
							}
						*/
						for _, v := range dataa {

							arrDetails = append(arrDetails, Item{

								Id:        post.Id,
								Url:       post.TheUrl,
								Timeresp:  timeresp,
								Num:       b,
								Timestamp: time.Now(),
								Def:       fmt.Sprintf(c, v),
							})

							//}
						}

						s = append(s, arrDetails...)

						//sa = append(sa, arrDetails...)

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
			s = []Item{}

			dat, err := json.Marshal(sa)
			if err != nil {
				panic(err)
			}

			//fmt.Println(string(data))
			xx = string(dat)
			//fmt.Println(string(data))

			//defer file.Close()
			time.Sleep(post.Time * time.Second)

			//defer file.Close()
			//time.Sleep(eachline.Time * time.Second)
		}

	}()

}
func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		ErrorCheck(err)
	}
	return count
}

var itemm Item

//Handlercheck to get all urls and its responses
func Handlercheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, x, html.EscapeString(r.URL.Path))

}

//GetItemEndpoint is to get a url and its responses
func GetItemEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	var coun int
	for i := 0; i < 1; i++ {
		for _, itemm = range onearrDetails {

			if itemm.Id == id {
				profiless := []Item{
					{itemm.Id, itemm.Url, itemm.Timeresp, itemm.Num, itemm.Timestamp, itemm.Def},
				}

				am = append(am, profiless...)

			}

		}
	}
	coun++
	if coun < 6 {
		json.NewEncoder(w).Encode(am)
		am = []Item{}
		time.Sleep(5 * time.Second)
	}

}

//InsertUser is post method
func InsertUser(response http.ResponseWriter, request *http.Request) {

	var userDetails Post
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()
	if err != nil {
		panic(err)

	} else {
		isInserted := insertUserInDB(userDetails)
		if isInserted {

			fmt.Println("done")
		} else {
			//return
		}
	}
}
func insertUserInDB(userDetails Post) bool {
	stmt, err := db.Prepare("INSERT into url_info SET url=?,time=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.TheUrl, userDetails.Time)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

//UpdateUser in the database regarde it's id
func UpdateUser(response http.ResponseWriter, request *http.Request) {

	userID := mux.Vars(request)["id"]
	if userID == "" {

		panic(err)
	} else {

		var userDetailss Post
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&userDetailss)
		defer request.Body.Close()

		if err != nil {
			panic(err)
		}

		isUpdated := updateUserInDB(userDetailss, userID)
		if isUpdated {

			fmt.Println("updated")
		} else {

			panic(err)
		}
	}
}

func updateUserInDB(userDetailss Post, userID string) bool {
	stmt, err := db.Prepare("UPDATE url_info SET url=?,time=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetailss.TheUrl, userDetailss.Time, userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

var db *sql.DB
var err error

//connect to database and write values
func ConnectDatabse() {
	db, err = sql.Open("mysql", "root@(127.0.0.1:3306)/healthcheck")
	fmt.Println("Database connected.")
	insert, err := db.Query("INSERT INTO url_info (url,time) VALUES ('http://www.google.com',5),('http://www.golang.org',5),('http://www.yahoo.com/games',5),('http://www.example.com',5)")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	//be careful deferring Queries if you are using transactions
	defer insert.Close()
}

//DeleteUser is to delete record from database
func DeleteUser(response http.ResponseWriter, request *http.Request) {

	userID := mux.Vars(request)["id"]
	if userID == "" {

		panic(err)
	} else {
		isdeleted := deleteUserFromDB(userID)
		if isdeleted {

			fmt.Println("deleted")

		} else {
			panic(err)

		}
	}
}
func deleteUserFromDB(userID string) bool {
	stmt, err := db.Prepare("DELETE FROM url_info WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
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
