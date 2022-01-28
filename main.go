func linkfile(site string) {

	LFile, err := os.Open("link.txt")

	if err != nil {
		log.Fatal(err)
	}
	fileScan := bufio.NewScanner(LFile)
	cont := 1
	for fileScan.Scan() {
		page := site + "/" + fileScan.Text()
		cont = cont + 1
		u, err := http.Get(page)

		if err != nil {
			log.Fatal(err)
		}
		if u.StatusCode == 200 {
			fmt.Println(page, "-OK---", cont)
		} else if u.StatusCode == 403 {
			fmt.Println(page, "-Forbidden---", cont)
		} else if u.StatusCode == 301 {
			fmt.Println(page, "-Moved Permanently---", cont)
		}
	}

	LFile.Close()
}
func main() {
	inicio := time.Now()

	site := "https://ejemplo.com"

	linkfile(site)

	timetotal := time.Since(inicio)

	fmt.Println(timetotal)
