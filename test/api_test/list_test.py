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
        body = resp.json()
        assert body.get("rev") is not None
        lid = body["id"]

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
        assert len(list_data["items"]) == len(body["items"])
        for want, got in zip(list_data["items"], body["items"]):
            assert want["title"] == got["title"]
            assert want.get("desc") == got.get("desc")
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
        body = resp.json()
        lid = body["id"]
        rev = body["rev"]

        # edit list - happy path

        data["title"] = "edited list"
        data["access"] = "public"
        resp = client.patch(f"/lists/{lid}", json=data)
        assert resp.status_code == 204

        resp = client.get(f"/lists/{lid}")
        assert resp.status_code == 200
        got = resp.json()
        self.assert_list_eq(data, got)

        # add list items - happy path

        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        items: list[dict] = resp.json()["items"]

        added_items = [
            {"title": "added"},
            {"title": "added with desc", "desc": "some description"}
        ]
        resp = client.post(f"/lists/{lid}/items", json={
            "rev": rev,
            "items": added_items
        })
        assert resp.status_code == 201
        assert resp.json()["rev"] == rev + 1
        rev += 1
        items.extend(added_items)

        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        body = resp.json()
        assert body["rev"] == rev
        assert len(items) == len(body["items"])
        for want, got in zip(items, body["items"]):
            assert want["title"] == got["title"]
            assert want.get("desc") == got.get("desc")
        items = body["items"]
        
        # delete list items - happy path

        delete_indices = [1, 3]
        delete_ids = [items[index]["id"] for index in delete_indices]
        resp = client.delete(f"/lists/{lid}/items", params={
            "rev": rev,
            "ids": ",".join(map(str, delete_ids))
        })
        assert resp.status_code == 200
        assert resp.json()["rev"] == rev + 1
        rev += 1
        items = [item for index, item in enumerate(items) if index not in delete_indices]

        resp = client.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        body = resp.json()
        assert body["rev"] == rev
        assert body["items"] == items

        # error handling

        resp = client.post(f"/lists/{lid}/items", json={
            "items": added_items
        })
        assert resp.status_code == 422

        resp = client.post(f"/lists/{lid}/items", json={
            "rev": rev-1,
            "items": added_items
        })
        assert resp.status_code == 409

        resp = client.delete(f"/lists/{lid}/items", params={
            "ids": delete_ids,
        })
        assert resp.status_code == 422

        resp = client.delete(f"/lists/{lid}/items", params={
            "rev": rev-1,
            "ids": delete_ids,
        })
        assert resp.status_code == 409

        u2 = make_user()
        u2.must_register(client)
        u2.must_login(client)

        resp = client.patch(f"/lists/{lid}", json=data)
        assert resp.status_code == 403

        resp = client.post(f"/lists/{lid}/items", json={
            "rev": rev,
            "items": added_items
        })
        assert resp.status_code == 403

        resp = client.delete(f"/lists/{lid}/items", params={
            "rev": rev,
            "ids": delete_ids,
        })
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
    
    def test_take(self, make_user: ty.Callable[[],User], list_data):
        users = [make_user() for _ in range(3)]
        clients = [Client() for _ in users]
        for u, c in zip(users, clients):
            u.must_register(c)
            u.must_login(c)
        u1, u2, u3 = users
        c1, c2, c3 = clients

        list_data["access"] = "public"
        resp = c1.post("/lists", json=list_data)
        assert resp.status_code == 201
        body = resp.json()
        lid = body["id"]

        resp = c2.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        body = resp.json()
        rev = body["rev"]
        iid = body["items"][0]["id"]
        assert body["items"][0].get("taken_by") is None

        resp = c2.post(f"/lists/{lid}/items/{iid}/taken_by", json={"rev": rev})
        assert resp.status_code == 204

        for c in [c2, c3]:
            resp = c.get(f"/lists/{lid}/items")
            assert resp.status_code == 200
            body = resp.json()
            assert body["rev"] == rev
            assert body["items"][0].get("taken_by") == u2.id
        
        resp = c1.get(f"/lists/{lid}/items")
        assert resp.status_code == 200
        body = resp.json()
        assert body["rev"] == rev
        assert body["items"][0].get("taken_by") is None

        for c in [c2, c3]:
            resp = c.post(f"/lists/{lid}/items/{iid}/taken_by", json={"rev": rev})
            assert resp.status_code == 409
            assert resp.json() == {"reason": "already taken", "taken_by": u2.id}
        
        resp = c1.post(f"/lists/{lid}/items/{iid}/taken_by", json={"rev": rev})
        assert resp.status_code == 403

        resp = c1.delete(f"/lists/{lid}/items/{iid}/taken_by", params={"rev": rev})
        assert resp.status_code == 403

        resp = c3.delete(f"/lists/{lid}/items/{iid}/taken_by", params={"rev": rev})
        assert resp.status_code == 409
        assert resp.json() == {"reason": "not taken"}

        resp = c2.delete(f"/lists/{lid}/items/{iid}/taken_by", params={"rev": rev})
        assert resp.status_code == 204

        for c in [c2, c3]:
            resp = c.get(f"/lists/{lid}/items")
            assert resp.status_code == 200
            assert resp.json()["items"][0].get("taken_by") is None
        
        resp = c2.delete(f"/lists/{lid}/items/{iid}/taken_by", params={"rev": rev})
        assert resp.status_code == 409
        assert resp.json() == {"reason": "not taken"}

        resp = c1.post(f"/lists/{lid}/items", json={"rev": rev, "items": [{"title": "new item"}]})
        assert resp.status_code == 201
        old_rev, rev = rev, resp.json()["rev"]

        for c in [c2, c3]:
            resp = c.post(f"/lists/{lid}/items/{iid}/taken_by", json={"rev": old_rev})
            assert resp.status_code == 409
            assert resp.json() == {"reason": "outdated revision"}
        
        resp = c2.post(f"/lists/{lid}/items/{iid}/taken_by", json={"rev": rev})
        assert resp.status_code == 204

        resp = c2.delete(f"/lists/{lid}/items/{iid}/taken_by", params={"rev": old_rev})
        assert resp.status_code == 409
        assert resp.json() == {"reason": "outdated revision"}






        


        
        
