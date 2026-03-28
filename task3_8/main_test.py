import pytest
import httpx

BASE_URL = "http://localhost:8000"

@pytest.mark.asyncio
async def test_fastapi_swagger_ui_exists():
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{BASE_URL}/docs")
        assert response.status_code == 200
        assert "swagger-ui" in response.text.lower()

@pytest.mark.asyncio
async def test_fastapi_openapi_json_structure():
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{BASE_URL}/openapi.json")
        assert response.status_code == 200
        spec = response.json()
        
        assert spec["info"]["title"] == "API Gateway Service"
        assert "/fetch-status" in spec["paths"]
        assert "/send-data" in spec["paths"]
        assert spec["paths"]["/send-data"]["post"]["responses"]["200"]

@pytest.mark.asyncio
async def test_fastapi_docs_negative_not_found():
    async with httpx.AsyncClient() as client:
        response = await client.get(f"{BASE_URL}/documentation")
        assert response.status_code == 404