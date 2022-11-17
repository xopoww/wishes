import typing as ty

from .user import User
from .client import Client
import pytest

@pytest.fixture
def list_data():
    return {
        "title": "list",
        "items": [
            {"title": "foo"},
            {"title": "bar", "desc": "with description"}
        ],
        "access": "private"
    }

class TestList:

    def assert_list_eq(self, want, got, check_items=False):
        assert want["title"] == got["title"]
        assert want["access"] == got["access"]
        if check_items:
            assert want["items"] == got["items"]

    def test_auth(self, client: Client, make_user: ty.Callable[[],User], list_data):
        u = make_user()
        u.must_register(client)

        resp = client.get(f"/lists")
        assert resp.status_code == 401

        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 401

    def test_post(self, client: Client, make_user: ty.Callable[[],User], list_data):
        u = make_user()
        u.must_register(client)
        u.must_login(client)

        resp = client.get(f"/lists")
        assert resp.status_code == 200
        assert resp.json() == []

        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.get(f"/lists")
        assert resp.status_code == 200
        assert resp.json() == [lid]

        list_data["access"] = "wrong"
        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 422
    
    def test_get(self, client: Client, make_user: ty.Callable[[],User], list_data):
        u = make_user()
        u.must_register(client)
        u.must_login(client)

        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.get(f"/lists/{lid}")
        assert resp.status_code == 200
        self.assert_list_eq(list_data, resp.json())

        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        body = resp.json()
        assert body["items"] == list_data["items"]
        assert body.get("rev") is not None

        resp = client.get(f"/lists/{lid + 50}")
        assert resp.status_code == 404

        resp = client.get(f"/lists/{lid + 50}/items")
        assert resp.status_code == 404

        resp = client.get("/lists/john/items")
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
            ],
            "access": "private"
        }
        resp = client.post("/lists", json=data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        rev = resp.json()["rev"]

        data["title"] = "edited list"
        data["access"] = "public"
        data["items"][0]["desc"] = "now with description"
        data["items"][1]["title"] = "edited bar"
        del data["items"][2]["desc"]
        data["items"].append({"title": "quux", "desc": "new item"})
        resp = client.patch(f"/lists/{lid}", json=data)
        assert resp.status_code == 204

        resp = client.get(f"/lists/{lid}")
        assert resp.status_code == 200
        got = resp.json()
        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        got |= resp.json()
        self.assert_list_eq(data, got, check_items=True)
        assert got["rev"] == rev + 1

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
    
    def test_delete(self, client: Client, make_user: ty.Callable[[],User], list_data):
        u1 = make_user()
        u1.must_register(client)
        u1.must_login(client)

        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 204
        
        resp = client.get("/lists")
        assert resp.status_code == 200
        assert resp.json() == []

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 404

        resp = client.delete(f"/lists/john")
        assert resp.status_code == 422

        resp = client.post("/lists", json=list_data)
        assert resp.status_code == 201
        lid = resp.json()["id"]

        u2 = make_user()
        u2.must_register(client)
        u2.must_login(client)

        resp = client.delete(f"/lists/{lid}")
        assert resp.status_code == 403
    
    def test_access(self, client: Client, make_user: ty.Callable[[],User], list_data):
        u1 = make_user()
        u1.must_register(client)
        u1.must_login(client)

        accesses = ["private", "public", "link"]
        lids = {}
        tokens = {}
        for access in accesses:
            data = list_data.copy()
            data["access"] = access
            
            resp = client.post("/lists", json=data)
            assert resp.status_code == 201
            lid = resp.json()["id"]

            resp = client.get(f"/lists/{lid}/token")
            assert resp.status_code == 200
            token = resp.json()["token"]

            lids[access] = lid
            tokens[access] = token
        
        # access by owner -- always ok
        for access in accesses:
            lid = lids[access]
            token = tokens[access]
            for url in [f"/lists/{lid}", f"/lists/{lid}/items"]:
                resp = client.get(url)
                assert resp.status_code == 200

                resp = client.get(url, params={"accessToken": token})
                assert resp.status_code == (200 if access != "private" else 403)

        # all lists are visible
        resp = client.get("/lists", params={"UserID": u1.id})
        assert resp.status_code == 200
        body = resp.json()
        for lid in lids.values():
            assert lid in body
        
            
        u2 = make_user()
        u2.must_register(client)
        u2.must_login(client)

        # only public list is visible
        resp = client.get("/lists", params={"UserID": u1.id})
        assert resp.status_code == 200
        assert resp.json() == [lids["public"]]

        # private list -- always forbidden
        for url in [f"/lists/{lids['private']}", f"/lists/{lids['private']}/items"]:
            for params in [{}, {"accessToken": tokens['private']}]:
                resp = client.get(url, params=params)
                assert resp.status_code == 403
        
        # public list -- always ok
        for url in [f"/lists/{lids['public']}", f"/lists/{lids['public']}/items"]:
            for params in [{}, {"accessToken": tokens['public']}]:
                resp = client.get(url, params=params)
                assert resp.status_code == 200
        
        # link list -- ok only with access token
        for url in [f"/lists/{lids['link']}", f"/lists/{lids['link']}/items"]:
            for params, code in zip([{}, {"accessToken": tokens['link']}], [403, 200]):
                resp = client.get(url, params=params)
                assert resp.status_code == code
        
        # wrong token -- forbidden
        for url in [f"/lists/{lids['link']}", f"/lists/{lids['link']}/items"]:
            resp = client.get(url, params={"accessToken": tokens["public"]})
            assert resp.status_code == 403
        
        for lid in lids.values():
            resp = client.get(f"/lists/{lid}/token")
            assert resp.status_code == 403
    

        


        
        
