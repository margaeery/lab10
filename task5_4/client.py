import asyncio
import websockets
import json
import sys

async def listen(ws):
    while True:
        try:
            raw_data = await ws.recv()
            data = json.loads(raw_data)
            sys.stdout.write("\r" + " " * 50 + "\r")
            print(f"[{data['user']}]: {data['content']}")
            sys.stdout.write("Вы: ")
            sys.stdout.flush()
        except websockets.ConnectionClosed:
            break

async def send(ws, user):
    while True:
        text = await asyncio.get_event_loop().run_in_executor(None, sys.stdin.readline)
        text = text.strip()
        if text:
            message = json.dumps({"user": user, "content": text})
            await ws.send(message)

async def start():
    name = input("Введите ваше имя: ")
    uri = "ws://localhost:8080/chat"
    async with websockets.connect(uri) as ws:
        print(f"--- Подключено к чату как {name} ---")
        sys.stdout.write("Вы: ")
        sys.stdout.flush()
        await asyncio.gather(listen(ws), send(ws, name))

asyncio.run(start())