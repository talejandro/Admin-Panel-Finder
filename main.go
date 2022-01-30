package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	inicio := time.Now()

	site := "https://itu.uncuyo.edu.ar/" //change here

	linkfile(site)

	timetotal := time.Since(inicio)

	fmt.Println(timetotal)
}

func linkfile(site string) {

	LFile, err := os.Open("link.txt")

	if err != nil {
		log.Fatal(err)
	}
	fileScan := bufio.NewScanner(LFile)
	cont := 1
	for fileScan.Scan() {
		page := site + fileScan.Text()
		cont = cont + 1
		u, err := http.Get(page)

		if err != nil {
			log.Fatal(err)
		}
		switch u.StatusCode {
		case 200:
			fmt.Println(page, "-OK---", cont)
		case 403:
			fmt.Println(page, "-Forbidden---", cont)
		case 301:
			fmt.Println(page, "-Moved Permanently---", cont)
		default:
		}
	}
	LFile.Close()
}
