import requests
import os

class Client:
    def __init__(self):
        host = os.environ.get("WISHES_HOST")
        assert host is not None
        self.base_url = host
        self.token: str|None = None
    
    def request(self, method: str, path: str, **kwargs) -> requests.Response:
        if self.token is not None and not kwargs.get("wishes_no_auth", False):
            headers = kwargs.get("headers", {})
            headers = {"x-token": self.token} | headers
            kwargs["headers"] = headers
        return requests.request(method, self.base_url + path, **kwargs)
    
    def get(self, path: str, **kwargs) -> requests.Response:
        return self.request("get", path, **kwargs)

    def post(self, path: str, **kwargs) -> requests.Response:
        return self.request("post", path, **kwargs)

    def patch(self, path: str, **kwargs) -> requests.Response:
        return self.request("patch", path, **kwargs)

    def delete(self, path: str, **kwargs) -> requests.Response:
        return self.request("delete", path, **kwargs)
