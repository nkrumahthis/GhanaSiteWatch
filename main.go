package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://ghanaweb.com",
		"https://passport.mfa.gov.gh",
		"https://ghs.gov.gh/covid19",
		"https://google.com",
		"https://myjoyonline.com",
	}

	start := time.Now()
	for _, site := range sites {
		hit(site)
	}
	fmt.Printf("\nTime elapsed: %v", time.Since(start))
}

func hit(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Printf("%s is down\n", site)
	} else {
		fmt.Printf("%s is up\n", site)
	}
}
