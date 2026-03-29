from fastapi import FastAPI
from pydantic import BaseModel, Field

app = FastAPI(
    title="Analytics Service API",
    description="Сервис для проверки метрик",
    version="1.0"
)

class Stats(BaseModel):
    users: int = Field(..., example=10)
    revenue: float = Field(..., example=500.5)

@app.get("/health", tags=["system"], summary="Проверка связи")
async def health_check():
    return {"status": "ok", "service": "analytics"}

@app.post("/check", tags=["analytics"], summary="Проверка данных", status_code=201)
async def check_data(data: Stats):
    return {
        "status": "success",
        "message": "Данные успешно получены",
        "captured_revenue": data.revenue
    }