import typing as ty

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


class TestUser:

    def test_get(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u2 = make_user()
        u2.must_register(client)
        
        for id in [u1.id, u2.id]:
            resp = client.get("/user", params={"id": id})
            assert resp.status_code == 401
        
        u2.must_login(client)
        for u in [u1, u2]:
            resp = client.get("/user", params={"id": u.id})
            assert resp.status_code == 200
            body = resp.json()
            assert body["id"] == u.id
            assert body["username"] == u.username
            assert body.get("fname") is None
            assert body.get("lname") is None
        
        resp = client.get("/user", params={"id": u2.id + 50})
        assert resp.status_code == 404

        resp = client.get("/user")
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
        
        resp = client.patch("/user", json={"id": u1.id}|info)
        assert resp.status_code == 401

        u1.must_login(client)
        
        resp = client.patch("/user", json={"id": u1.id}|info)
        assert resp.status_code == 200
        resp = client.get("/user", params={"id": u1.id})
        assert resp.status_code == 200
        body = resp.json()
        assert body["fname"] == info["fname"]
        assert body["lname"] == info["lname"]

        resp = client.patch("/user", json={"id": u2.id}|info)
        assert resp.status_code == 403

        resp = client.patch("/user", json=info)
        assert resp.status_code == 422
