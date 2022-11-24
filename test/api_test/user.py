import requests
from .client import Client

class User:
    def __init__(self, username: str, password: str):
        self.username = username
        self.password = password
        self.id: int|None = None
    
    def register(self, c: Client) -> requests.Response:
        return c.post("/auth/register", json={"username": self.username, "password": self.password})
    
    def must_register(self, c: Client):
        resp = self.register(c)
        assert resp.status_code == 200
        body = resp.json()
        assert body["ok"], body.get("error") or "error"
        self.id = body["user"]["id"]

    def login(self, c: Client):
        return c.post("/auth/login", json={"username": self.username, "password": self.password})

    def must_login(self, c: Client):
        resp = self.login(c)
        assert resp.status_code == 200
        body = resp.json()
        assert body["ok"] == True
        c.token = body["token"]