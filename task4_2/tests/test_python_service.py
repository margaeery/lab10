import requests
import pytest

DIRECT_URL = "http://localhost:8001"
GATEWAY_URL = "http://localhost:8000/analytics"

def test_py_direct_health():
    r = requests.get(f"{DIRECT_URL}/health")
    assert r.status_code == 200
    assert r.json()["status"] == "ok"

def test_py_gateway_health():
    r = requests.get(f"{GATEWAY_URL}/health")
    assert r.status_code == 200
    assert "service" in r.json()

def test_py_gateway_check_valid_data():
    payload = {"users": 150, "revenue": 4500.75}
    r = requests.post(f"{GATEWAY_URL}/check", json=payload)
    assert r.status_code == 201
    assert r.json()["captured_revenue"] == 4500.75

def test_py_direct_check_wrong_type():
    payload = {"users": 10.5, "revenue": 100}
    r = requests.post(f"{DIRECT_URL}/check", json=payload)
    assert r.status_code == 422

def test_py_gateway_check_missing_field():
    payload = {"users": 100}
    r = requests.post(f"{GATEWAY_URL}/check", json=payload)
    assert r.status_code == 422

def test_py_direct_check_extra_fields():
    payload = {"users": 10, "revenue": 100, "extra": "data"}
    r = requests.post(f"{DIRECT_URL}/check", json=payload)
    assert r.status_code == 201

def test_py_gateway_health_content_type():
    r = requests.get(f"{GATEWAY_URL}/health")
    assert "application/json" in r.headers["Content-Type"]