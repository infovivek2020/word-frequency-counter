package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		expectedStatus int
		expectedResult []WordCount
	}{
		{
			name:           "Normal input",
			query:          "Go+is+fun+and+go+is+easy",
			expectedStatus: http.StatusOK,
			expectedResult: []WordCount{
				{"go", 2},
				{"is", 2},
				{"and", 1},
				{"easy", 1},
				{"fun", 1},
			},
		},
		{
			name:           "Empty input",
			query:          "",
			expectedStatus: http.StatusBadRequest,
			expectedResult: nil,
		},
		{
			name:           "Case insensitive with punctuation",
			query:          "Go, go! IS? is.",
			expectedStatus: http.StatusOK,
			expectedResult: []WordCount{
				{"go", 2},
				{"is", 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/count?text="+tt.query, nil)
			rec := httptest.NewRecorder()

			countHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var got []WordCount
				if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to parse response: %v", err)
				}

				if len(got) != len(tt.expectedResult) {
					t.Fatalf("expected %d results, got %d", len(tt.expectedResult), len(got))
				}

				for i := range got {
					if got[i] != tt.expectedResult[i] {
						t.Errorf("mismatch at index %d: got %v, expected %v", i, got[i], tt.expectedResult[i])
					}
				}
			}
		})
	}
}
