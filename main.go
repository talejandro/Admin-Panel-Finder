package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func linkfile(site string) {

	LFile, err := os.Open("link.txt")

	if err != nil {
		log.Fatal(err)
	}
	fileScan := bufio.NewScanner(LFile)

	for fileScan.Scan() {
		page := site + "/" + fileScan.Text()

		u, err := http.Get(page)

		if err != nil {
			log.Fatal(err)
		}
		if u.StatusCode != 404 {
			fmt.Println(page)
		}
	}

	LFile.Close()

}

func main() {

	site := "https:exaple.com"
	linkfile(site)

}
