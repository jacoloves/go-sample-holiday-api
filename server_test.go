package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {

	h1 := holiday_2021()
	h2 := holiday_2022()
	h3 := holiday_2023()

	dHandler := &dateHandler{
		store: &datastore{
			m: map[string]Data{
				"2021": Data(h1),
				"2022": Data(h2),
				"2023": Data(h3),
			},
			RWMutex: &sync.RWMutex{},
		},
	}
	ts := httptest.NewServer(http.Handler(dHandler))
	defer ts.Close()
	// /holiday pass Test
	resp, err := http.Get(fmt.Sprintf("%s/holiday", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	// /holiday/ pass Test
	resp2, err := http.Get(fmt.Sprintf("%s/holiday/", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp2.StatusCode)
	}

	// /holiday/year/2022 pass Test
	resp3, err := http.Get(fmt.Sprintf("%s/holiday/year/2022", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp3.Body.Close()

	if resp3.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp3.StatusCode)
	}

	// /holiday/year/2021 pass Test
	resp4, err := http.Get(fmt.Sprintf("%s/holiday/year/2021", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp4.Body.Close()

	if resp4.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp4.StatusCode)
	}

	// /holiday/year/2023 pass Test
	resp5, err := http.Get(fmt.Sprintf("%s/holiday/year/2023", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp5.Body.Close()

	if resp5.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp5.StatusCode)
	}

	// /holiday/year/2020 not pass Test
	resp6, err := http.Get(fmt.Sprintf("%s/holiday/year/2020", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp6.Body.Close()

	if resp6.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %v", resp6.StatusCode)
	}

	// /holiday/year/2024 not pass Test
	resp7, err := http.Get(fmt.Sprintf("%s/holiday/year/2024", ts.URL))
	if err != nil {
		t.Fatalf("Excepted no error, got %v", err)
	}
	defer resp7.Body.Close()

	if resp7.StatusCode != http.StatusNotFound {
		t.Fatalf("Expected status code 404, got %v", resp7.StatusCode)
	}
}
