import hashlib
import random
import typing as ty

import pytest
from .client import Client
from .user import User

@pytest.fixture
def client() -> Client:
    return Client()

@pytest.fixture(scope="session")
def session_id() -> str:
    random.seed()
    return random.randbytes(2).hex()

@pytest.fixture(scope="function")
def make_user(request: pytest.FixtureRequest, session_id: str) -> ty.Callable[[],User]:
    caller_id = request.function.__name__
    if request.cls is not None:
        caller_id = f"{request.cls.__name__}.{caller_id}"
    caller_hash = hashlib.shake_128(caller_id.encode("ascii")).hexdigest(3)
    username_base = f"u.{session_id}.{caller_hash}"
    
    n_users = 0

    def _make_user(password="password") -> User:
        nonlocal n_users
        username = f"{username_base}.{n_users}"
        n_users += 1
        return User(username, password)

    return _make_user