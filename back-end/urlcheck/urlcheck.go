package urlcheck

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//Data is
type Data struct {
	//Tf     bool   `json:"tf"`
	Code   int    `json:"code"`
	Status string `json:"status"`
	Time   string `json:"time"`
}
type Teststruct struct {
	Test string
}

var url string

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

//Handlercheck is
func Handlercheck(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("app.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
		json.NewEncoder(w).Encode(eachline)
	}
	decoder := json.NewDecoder(r.Body)

	var t Teststruct
	err = decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	//fmt.Println(t.Test)
	//for {
	//	Urlcheck(t.Test)
	//	}
	/*	f, err := os.Create("testt.txt")
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
		}*/
} /*
	file, err := os.Open("/path/to/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(s)
	w.Write(scanner.Text())

}

/*
func Handlercheck(w http.ResponseWriter, r *http.Request) {
	var s []Data
	tf, code, status := Urlcheck("https://www.google.com")
	thedata := []Data{
		{tf, code, status},
	}
	s = append(s, thedata...)
	json.NewEncoder(w).Encode(s)
	//	time.Sleep(5 * time.Second)

}
*/
//Urlcheck is
func Urlcheck(url string) (int, string, time.Time) {
	now := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	if resp.StatusCode == 200 {
		fmt.Println(true, resp.StatusCode, http.StatusText(resp.StatusCode), now)
		return resp.StatusCode, http.StatusText(resp.StatusCode), now

	}
	fmt.Println(false, resp.StatusCode, http.StatusText(resp.StatusCode), now)
	//time.Sleep(5 * time.Second)
	return resp.StatusCode, http.StatusText(resp.StatusCode), now

	//time.Sleep(5 * time.Second)

}
