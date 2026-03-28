import pytest
import httpx
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_fetch_status_integration():
    response = client.get("/fetch-status")
    assert response.status_code == 200
    assert "data" in response.json()
    assert response.json()["data"]["status"] == "ok"

def test_fetch_info_integration():
    response = client.get("/fetch-info")
    assert response.status_code == 200
    assert response.json()["service"] == "go-gin"

def test_send_data_integration():
    payload = {"test_key": "test_value"}
    response = client.post("/send-data", json=payload)
    assert response.status_code == 200
    assert response.json()["message"] == "data received"

def test_not_found_integration():
    response = client.get("/fetch-non-existent")
    assert response.status_code == 404

def test_go_unavailable_integration():
    response = client.get("/fetch-status")
    if response.status_code == 503:
        assert response.json()["detail"] == "Go service unavailable"