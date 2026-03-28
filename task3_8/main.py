import httpx
from fastapi import FastAPI, HTTPException

app = FastAPI(
    title="API Gateway Service",
    description="Интеграционный сервис для связи с Go-бэкендом",
    version="1.0.0"
)

GO_SERVICE_URL = "http://localhost:8080"

async def make_request(method, path, json=None):
    async with httpx.AsyncClient() as client:
        try:
            url = f"{GO_SERVICE_URL}{path}"
            response = await client.request(method, url, json=json)
            response.raise_for_status()
            return response.json()
        except httpx.HTTPStatusError as e:
            raise HTTPException(
                status_code=e.response.status_code,
                detail=f"Go service error: {e.response.text}"
            )
        except httpx.RequestError:
            raise HTTPException(
                status_code=503,
                detail="Go service unavailable"
            )

@app.get("/fetch-status", tags=["Go Proxy"])
async def get_go_status():
    data = await make_request("GET", "/status")
    return {"source": "go-service", "data": data}

@app.get("/fetch-info", tags=["Go Proxy"])
async def get_go_info():
    return await make_request("GET", "/info")

@app.post("/send-data", tags=["Go Proxy"])
async def send_to_go(payload: dict):
    return await make_request("POST", "/data", json=payload)