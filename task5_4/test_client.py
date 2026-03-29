import pytest
import asyncio
import websockets
import json

@pytest.mark.asyncio
async def test_real_connection():
    uri = "ws://localhost:8080/chat"
    async with websockets.connect(uri) as ws:
        message = {"user": "PyTest", "content": "Integration Check"}
        await ws.send(json.dumps(message))
        
        response = await ws.recv()
        data = json.loads(response)
        
        assert data["user"] == "PyTest"
        assert data["content"] == "Integration Check"

@pytest.mark.asyncio
async def test_multi_client_broadcast():
    uri = "ws://localhost:8080/chat"
    async with websockets.connect(uri) as client1, \
               websockets.connect(uri) as client2:
        
        msg = {"user": "Alice", "content": "Hi Bob"}
        await client1.send(json.dumps(msg))
        
        resp1 = await client1.recv()
        resp2 = await client2.recv()
        
        assert json.loads(resp1) == msg
        assert json.loads(resp2) == msg