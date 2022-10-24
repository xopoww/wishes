CREATE TABLE Users (
    user_id     INTEGER PRIMARY KEY,
    user_name   TEXT    NOT NULL UNIQUE,

    pwd_hash    TEXT    NOT NULL
);