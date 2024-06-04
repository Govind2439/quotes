package main

import (
	"fmt"
	"io"
	"net/http"
)

func getQuote(w http.ResponseWriter, r *http.Request) {
	url := "https://quotes15.p.rapidapi.com/quotes/random/?language_code=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "ee661ee060mshd9ae3e417ebd7f1p1089c5jsn2fa1b576c16f")
	req.Header.Add("x-rapidapi-host", "quotes15.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch quote", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Write(body)
}

func main() {
	http.HandleFunc("/quote", getQuote)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
