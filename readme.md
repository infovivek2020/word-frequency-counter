# Word Frequency Counter API (Go)

A simple HTTP API written in Go that counts how many times each word appears in a given input text.  
The API is **case-insensitive**, ignores punctuation, and returns results sorted by:

1. **Highest frequency first**
2. **Alphabetically for ties**

---

## Features
- Accepts text input via query parameter (`GET /count?text=...`)
- Removes punctuation (e.g., commas, periods, etc.)
- Converts text to lowercase before counting
- Returns results in JSON format
- Sorts by frequency (descending) and alphabetically (ascending) when counts match

---

## Example

**Request:**
```bash
curl "http://localhost:8080/count?text=Go+is+fun+and+go+is+easy"

**Response:**

[
  {"word": "go", "count": 2},
  {"word": "is", "count": 2},
  {"word": "and", "count": 1},
  {"word": "easy", "count": 1},
  {"word": "fun", "count": 1}
]
