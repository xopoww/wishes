import typing as ty

from .user import User
from .client import Client

class TestList:

    def test_auth(self, client: Client, make_user: ty.Callable[[],User]):
        u = make_user()
        u.must_register(client)

        resp = client.get(f"/users/{u.id}/lists")
        assert resp.status_code == 401

        resp = client.post("/lists", json={
            "title": "list",
            "items": [
                {"title": "foo"},
                {"title": "bar", "desc": "with description"}
            ]
        })
        assert resp.status_code == 401

    def test_post(self, client: Client, make_user: ty.Callable[[],User]):
        u = make_user()
        u.must_register(client)
        u.must_login(client)

        resp = client.get(f"/users/{u.id}/lists")
        assert resp.status_code == 200
        assert resp.json() == []

        resp = client.post("/lists", json={
            "title": "list",
            "items": [
                {"title": "foo"},
                {"title": "bar", "desc": "with description"}
            ]
        })
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.get(f"/users/{u.id}/lists")
        assert resp.status_code == 200
        assert resp.json() == [lid]
    
    def test_get(self, client: Client, make_user: ty.Callable[[],User]):
        u = make_user()
        u.must_register(client)
        u.must_login(client)

        data = {
            "title": "list",
            "items": [
                {"title": "foo"},
                {"title": "bar", "desc": "with description"}
            ]
        }
        resp = client.post("/lists", json=data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.get(f"/lists/{lid}")
        assert resp.status_code == 200
        assert resp.json() == data

        resp = client.get(f"/lists/{lid + 50}")
        assert resp.status_code == 404

        resp = client.get("/lists")
        assert resp.status_code == 405

        resp = client.get("/lists/john")
        assert resp.status_code == 422
    
    def test_patch(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u1.must_login(client)

        data = {
            "title": "list",
            "items": [
                {"title": "foo"},
                {"title": "bar", "desc": "with description"},
                {"title": "baz", "desc": "also with description"}
            ]
        }
        resp = client.post("/lists", json=data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        data["title"] = "edited list"
        data["items"][0]["desc"] = "now with description"
        data["items"][1]["title"] = "edited bar"
        del data["items"][2]["desc"]
        data["items"].append({"title": "quux", "desc": "new item"})
        resp = client.patch(f"/lists/{lid}", json=data)
        assert resp.status_code == 204

        resp = client.get(f"/lists/{lid}")
        assert resp.status_code == 200
        assert resp.json() == data

        resp = client.patch(f"/lists/{lid + 50}", json=data)
        assert resp.status_code == 404

        resp = client.patch("/lists", json=data)
        assert resp.status_code == 405

        resp = client.patch("/lists/john", json=data)
        assert resp.status_code == 422

        resp = client.patch(f"/lists/{lid}", json={"wrong": "fields", "in": "body"})
        assert resp.status_code == 422

        u2 = make_user()
        u2.must_register(client)
        u2.must_login(client)

        resp = client.patch(f"/lists/{lid}", json=data)
        assert resp.status_code == 403
    
    def test_delete(self, client: Client, make_user: ty.Callable[[],User]):
        u1 = make_user()
        u1.must_register(client)
        u1.must_login(client)

        data = {
            "title": "list",
            "items": [
                {"title": "foo"},
                {"title": "bar", "desc": "with description"},
                {"title": "baz", "desc": "also with description"}
            ]
        }
        resp = client.post("/lists", json=data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 204
        
        resp = client.get(f"/users/{u1.id}/lists")
        assert resp.status_code == 200
        assert resp.json() == []

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 404

        resp = client.delete(f"/lists/john")
        assert resp.status_code == 422

        resp = client.post("/lists", json=data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        u2 = make_user()
        u2.must_register(client)
        u2.must_login(client)

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 403
