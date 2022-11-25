package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		name := r.FormValue("name")
		adress := r.FormValue("adress")
		lname := r.FormValue("lname")
		bday := r.FormValue("birthday")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", adress)
		fmt.Fprintf(w, "Lastname = %s\n", lname)
		fmt.Fprintf(w, "Bday = %v\n", bday)
		if name != "" && adress != "" && bday != "" && lname != "" {

		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func dice(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dice" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Not an allowed operation")
	case "POST":
		fmt.Fprintf(w, "%04d", rand.Intn(1000))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func currentTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Not an allowed operation")
	case "POST":
		fmt.Fprintf(w, "%02dh%02d", time.Now().Hour(), time.Now().Minute())
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func dices(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dices" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Not an allowed operation")
	case "POST":
		dices := [8]int{2, 4, 6, 8, 10, 12, 20, 100}
		var randint int
		for i := 0; i < 15; i++ {
			randint = rand.Intn(7)
			if randint == 7 {
				fmt.Fprintf(w, "%03d", rand.Intn(100))
			}
			if dices[randint] < 21 && dices[randint] > 9 {
				fmt.Fprintf(w, "%02d", rand.Intn(dices[randint]))
			} else {
				fmt.Fprintf(w, "%03d", rand.Intn(dices[randint]))
			}
			switch randint {
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", currentTime)
	http.HandleFunc("/dice", dice)
	http.HandleFunc("/dices", dices)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":9000", nil)

}
