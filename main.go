package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func urlAccessible(url string) bool {
	client := http.Client{
		Timeout: time.Second * 10, // Maximum of 10 secs
	}
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	domain := strings.ReplaceAll(input, "[", "")
	domain = strings.ReplaceAll(domain, "]", "")
	domain = strings.TrimSpace(domain) // remove newline character

	_, err := net.LookupHost(domain)
	if err != nil {
		fmt.Println("DNS record does not exist.")
	} else {
		httpsURL := "https://" + domain
		httpURL := "http://" + domain
		if urlAccessible(httpsURL) {
			fmt.Println(httpsURL)
		} else if urlAccessible(httpURL) {
			fmt.Println(httpURL)
		} else {
			fmt.Println("No url.")
		}
	}
}
