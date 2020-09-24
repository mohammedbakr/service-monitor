package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/healthcheck/urlcheck"
	"github.com/rs/cors"
)

/*type Data struct {
	Tf     bool   `json:"tf"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}*/
/*type ResponseObj struct {
	Response []Item `json:"response"`
}*/

type Item struct {
	Typetf bool   `json:"state"`
	Num    int    `json:"code"`
	Def    string `json:"status"`
}
type Items []Item

func main() {
	go func() {
		for {

			a, b, c := urlcheck.Urlcheck("https://www.yahoo.com/games")
			item := Item{a, b, c}
			js, err := json.Marshal(item)
			if err != nil {
				panic(err)

			}
			fmt.Println(string(js))
			/*	f, err := os.Create("a.txt")
				if err != nil {
					defer f.Close()
				}
				w := bufio.NewWriter(f)
				n4, err := w.WriteString(js)
				if err != nil {
					panic(err)
				}
				fmt.Printf("wrote %d bytes\n", n4)*/
			jsonFile, err := os.OpenFile("title.txt",
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			//	jsonFile, err := os.Create("./data.txt")
			jsonFile.Write(js)
			jsonFile.WriteString("\n")
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/configurations", urlcheck.Handlercheck)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":10000", handler)

}

//json_byte := []byte(`{"response":[{"t_int":, "t_bool": true,  "t_null_or_string": null}, {"t_int":2, "t_bool": false, "t_null_or_string": "string1"}]}`)

/*	data_json := ResponseObj{}
		if err := json.Unmarshal(json_byte, &data_json); err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", data_json)

	}
	//go func() {

	//	mux := http.NewServeMux()
	//	mux.HandleFunc("/", urlcheck.Handlercheck)
	//handler := cors.Default().Handler(mux)
	//	time.Sleep(5 * time.Second)
	//	handler := cors.Default().Handler(mux)
	//	log.Println(http.ListenAndServe("localhost:6060", handler))

	//}()
	//handler := cors.Default().Handler(mux)
	//log.Println(http.ListenAndServe("localhost:6060", handler))
	//time.Sleep(5 * time.Second)

}

/*
type Data struct {
	Tf     bool   `json:"tf"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var stream = health.NewStream()

func vastPlayer(w http.ResponseWriter, r *http.Request) {
	var s []Data
	tf, code, status := Urlcheck("https://www.google.com")
	thedata := []Data{
		{tf, code, status},
	}
	s = append(s, thedata...)
	json.NewEncoder(w).Encode(s)
	time.Sleep(5 * time.Second)
}
func Urlcheck(url string) (bool, int, string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println(true, resp.StatusCode, http.StatusText(resp.StatusCode))
		return true, resp.StatusCode, http.StatusText(resp.StatusCode)

	}
	fmt.Println(false, resp.StatusCode, http.StatusText(resp.StatusCode))
	//time.Sleep(5 * time.Second)
	return false, resp.StatusCode, http.StatusText(resp.StatusCode)

	//time.Sleep(5 * time.Second)

}

func main() {
	// Log to stdout!
	stream.AddSink(&health.WriterSink{os.Stdout})
	// Make sink and add it to stream
	sink := health.NewJsonPollingSink(time.Second*5, time.Second*20)
	stream.AddSink(sink)
	// Start the HTTP server! This will expose metrics via a JSON API.
	adr := "127.0.0.1:5001"
	sink.StartServer(adr)

	http.HandleFunc("/api/getVastPlayer", vastPlayer)
	log.Println("Listening...")
	panic(http.ListenAndServe(":2001", nil))
}
func doSomething(s string) {
	fmt.Println("doing something", s)
}

func startPolling1() {
	for {
		time.Sleep(2 * time.Second)
		go doSomething("from polling 1")
	}
}

func startPolling2() {
	for {
		<-time.After(2 * time.Second)
		go doSomething("from polling 2")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	go startPolling1()
	go startPolling2()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

/*	mux := http.NewServeMux()
	mux.HandleFunc("/check", urlcheck.Handlercheck)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":1000", handler)
	//	}

}*/
/*
var (
	timeSumsMu sync.RWMutex
	timeSums   int64
)

func main() {
	// Start the goroutine that will sum the current time
	// once per second.
	go runDataLoop()
	// Create a handler that will read-lock the mutext and
	// write the summed time to the client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeSumsMu.RLock()
		defer timeSumsMu.RUnlock()
		fmt.Fprint(w, timeSums)
	})
	http.ListenAndServe(":8080", nil)
}

func runDataLoop() {
	for {
		// Within an infinite loop, lock the mutex and
		// increment our value, then sleep for 1 second until
		// the next time we need to get a value.
		timeSumsMu.Lock()
		timeSums += time.Now().Unix()
		timeSumsMu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
*/
