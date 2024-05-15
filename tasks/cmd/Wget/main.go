package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./my_wget <URL>")
		os.Exit(1)
	}

	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	downloadSite(url)
}

func downloadSite(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	baseURL := getBaseURL(url)
	baseURL = cleanFileName(baseURL)
	err = os.WriteFile(baseURL+".html", bodyBytes, 0644)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Site downloaded successfully.")
}

func getBaseURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) >= 3 {
		return parts[0] + "//" + parts[2]
	}
	return ""
}

func cleanFileName(fileName string) string {
	// Replace invalid characters with "-"
	reg := regexp.MustCompile(`[^\w\d\-]+`)
	return reg.ReplaceAllString(fileName, "-")
}
