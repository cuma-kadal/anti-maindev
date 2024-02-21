package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
)

func executeDomain(w http.ResponseWriter, r *http.Request) {
	urlParam := r.URL.Query().Get("url")

	if isValidURL(urlParam) {
		cmd := exec.Command("./spurt","--url", urlParam)
		out, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(w, "Error executing command: %v", err)
			return
		}

		fmt.Fprintf(w, "%s", out)
	} else {
		fmt.Fprintf(w, "Invalid URL")
	}
}

func isValidURL(url string) bool {
	// Regex for a simple check, you may need to adjust it based on your requirements
	regex := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)
	return regex.MatchString(url)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/execute", executeDomain)
	http.ListenAndServe(":8080", nil)
}
