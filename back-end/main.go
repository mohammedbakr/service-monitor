/*package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"



	"github.com/rs/cors"
)

type FruitBasket struct {
	Url string
}

//url = nil
var url []byte
var resp = []byte(`{
	"result": [
	 {
	  "id": "ID 1"
	 },
	 {
	  "id": "ID 2"
	 }
	]
   }`)

func main() {
	//	url = nil
	//	decodeurl.X(url)
	mux := http.NewServeMux()
	/*	mux.HandleFunc("/url/check", func(w http.ResponseWriter, r *http.Request) {
		decodeurl.X()
	})

	type ResultStruct struct {
		result []map[string]string
	}
	var jsonData ResultStruct
	err := json.NewDecoder(resp.Body).Decode(&jsonData)
	//var jsonData ResultStruct
	//err = json.Unmarshal(respBytes, &jsonData)

	mux.HandleFunc("/api/url", func(w http.ResponseWriter, r *http.Request) {

		jsonData := []byte(Url)
		var basket FruitBasket
		err := json.Unmarshal(jsonData, &basket)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(basket.Url)
		//fmt.Println(basket.Created)

		f, err := os.Create("test.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString(basket.Url)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	//decodeurl.X()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/url", decodeurl.X)

	//cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":1000", handler)

}*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"os"
)

//Teststruct is
type Teststruct struct {
	Test string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t Teststruct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Test)
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(t.Test)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	/*http.HandleFunc("/", parseGhPost)
	http.ListenAndServe(":8080", nil)*/

	mux := http.NewServeMux()
	mux.HandleFunc("/api/configurations", parseGhPost)

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":10000", handler)
}
