package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	year := flag.Int("year", 2020, "year of the event")
	day := flag.Int("day", 1, "day")
	session := flag.String("session", "", "the value of your session cookie (required)")
	output := flag.String("output", "", "path to target output file (required)")

	flag.Parse()

	if *session == "" || *output == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	url := fmt.Sprint("https://adventofcode.com/", *year, "/day/", *day, "/input")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: *session})

	client := &http.Client{}

	resp, err := client.Do(req)

	if err == nil && resp != nil {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)

		file, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		length, err := file.WriteString(bodyString)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(length, " bytes written to file")

	} else {
		log.Fatal(err)
	}
}
