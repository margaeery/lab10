package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const (
	gatewayGoURL = "http://localhost:8000/api/v1/items"
	directGoURL  = "http://localhost:8002/items"
)

func TestGo_Direct_GetItems_Status200(t *testing.T) {
	resp, err := http.Get(directGoURL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %v", resp.StatusCode)
	}
}

func TestGo_Gateway_GetItems_ArrayCheck(t *testing.T) {
	resp, err := http.Get(gatewayGoURL)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	var items []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if len(items) < 2 {
		t.Errorf("Expected at least 2 items, got %d", len(items))
	}
}

func TestGo_Gateway_CreateItem_Success(t *testing.T) {
	payload := map[string]interface{}{
		"id":    "3",
		"name":  "Keyboard",
		"price": 2500.0,
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(gatewayGoURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201, got %d", resp.StatusCode)
	}
}

func TestGo_Direct_CreateItem_NegativePrice(t *testing.T) {
	payload := map[string]interface{}{
		"id":    "4",
		"name":  "Free Item",
		"price": -1.0,
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(directGoURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != 422 {
		t.Errorf("Expected 422 for negative price, got %d", resp.StatusCode)
	}
}

func TestGo_Gateway_CreateItem_EmptyName(t *testing.T) {
	payload := map[string]interface{}{
		"id":    "5",
		"name":  "",
		"price": 100.0,
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(gatewayGoURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 (logic allows empty name), got %d", resp.StatusCode)
	}
}