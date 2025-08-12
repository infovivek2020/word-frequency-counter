package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "missing 'text' query parameter", http.StatusBadRequest)
		return
	}

	// Normalize: lowercase, remove punctuation
	re := regexp.MustCompile(`[^\w\s]`)
	cleanText := re.ReplaceAllString(strings.ToLower(text), "")

	// Split into words
	words := strings.Fields(cleanText)

	// Count frequencies
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	// Convert to slice for sorting
	var result []WordCount
	for word, count := range freq {
		result = append(result, WordCount{Word: word, Count: count})
	}

	// Sort: by count desc, then alphabetically asc
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count == result[j].Count {
			return result[i].Word < result[j].Word
		}
		return result[i].Count > result[j].Count
	})

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/count", countHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
