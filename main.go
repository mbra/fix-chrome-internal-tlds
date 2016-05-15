package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var urlPattern string = "^(https?:\\/\\/)?([\\da-zA-Z\\.-]+)(\\.[\\da-zA-Z\\.]*)*\\.(%s)(:[2-9]{1,5})?([\\/\\w \\.-]*)*\\/?$"
var searchEngine string = ""
var internalDomains []string

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %s\n", err)
		return
	}
	s := r.Form.Get("s")

	matched, err := regexp.Match(fmt.Sprintf(urlPattern, strings.Join(internalDomains, "|")), []byte(s))
	if err != nil {
		log.Printf("Error parsing form: %s\n", err)
		return
	}

	if matched {
		log.Printf("Search is internal URL: %s\n", s)
		http.Redirect(w, r, fmt.Sprintf("http://%s/", s), http.StatusFound)
	} else {
		log.Printf("Search is probably a real search string: %s", s)
		http.Redirect(w, r, fmt.Sprintf(searchEngine, s), http.StatusFound)
	}
}

func main() {
	var domains, listen = "", ""
	flag.StringVar(&searchEngine, "searchstring", "https://www.google.com/search?q=%s", "The searchengine string you use")
	flag.StringVar(&domains, "domains", "", "A comma seperated list of domain names")
	flag.StringVar(&listen, "listen", ":8080", "The IP and port to listen on")
	flag.Parse()

	internalDomains = strings.Split(domains, ",")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
	}
}
