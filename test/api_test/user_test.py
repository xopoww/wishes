from datetime import datetime, timedelta
import typing as ty
import pytest
import jwt

from .user import User
from .client import Client


class TestAuth:

    def test_register(self, client: Client, make_user: ty.Callable[[],User]):
        u = make_user()
        
        u.must_register(client)
        assert u.id is not None
        
        resp = u.register(client)
        assert resp.status_code == 200
        assert resp.json()["ok"] == False


    def test_login(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u1.must_login(client)
        assert client.token is not None
        client.token = None

        u2 = make_user()
        resp = u2.login(client)
        assert resp.status_code == 200
        assert resp.json()["ok"] == False

        u3 = User(u1.username, u1.password + "wrong")
        resp = u3.login(client)
        assert resp.status_code == 200
        assert resp.json()["ok"] == False
    
    @pytest.mark.parametrize("payload,secret", [
        ({
            "iat": datetime.now(),
            "nbf": datetime.now() - timedelta(seconds=5),
            "exp": datetime.now() + timedelta(hours=1),
            "sub": "test",
            "alg": "HS256"
        }, "wrong-secret"),
        ({
            "iat": datetime.now(),
            "nbf": datetime.now() - timedelta(seconds=5),
            "exp": datetime.now() + timedelta(hours=1),
            "sub": "test",
            "alg": "none",
        }, "")
    ])
    def test_token(self, client: Client, payload, secret):
        client.token = jwt.encode(payload=payload, key=secret)
        resp = client.get("/users/1")
        assert resp.status_code == 401


class TestUser:

    def test_get(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u2 = make_user()
        u2.must_register(client)
        
        for id in [u1.id, u2.id]:
            resp = client.get(f"/users/{id}")
            assert resp.status_code == 401
        
        u2.must_login(client)
        for u in [u1, u2]:
            resp = client.get(f"/users/{u.id}")
            assert resp.status_code == 200
            body = resp.json()
            assert body["username"] == u.username
            assert body.get("fname") == ""
            assert body.get("lname") == ""
        
        resp = client.get(f"/users/{u2.id + 50}")
        assert resp.status_code == 404

        resp = client.get("/users")
        assert resp.status_code == 405

        resp = client.get("/users/john")
        assert resp.status_code == 422
    
    def test_patch(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u2 = make_user()
        u2.must_register(client)

        info = {
            "fname": "John",
            "lname": "Doe"
        }
        
        resp = client.patch(f"/users/{u1.id}", json=info)
        assert resp.status_code == 401

        u1.must_login(client)
        
        resp = client.patch(f"/users/{u1.id}", json=info)
        assert resp.status_code == 204
        resp = client.get(f"/users/{u1.id}")
        assert resp.status_code == 200
        body = resp.json()
        assert body["fname"] == info["fname"]
        assert body["lname"] == info["lname"]

        resp = client.patch(f"/users/{u2.id}", json=info)
        assert resp.status_code == 403

        resp = client.patch("/users", json=info)
        assert resp.status_code == 405

        resp = client.patch("/users/john", json=info)
        assert resp.status_code == 422

        resp = client.patch(f"/users/{u1.id}", json={"wrong": "fields", "in": "body"})
        assert resp.status_code == 422
