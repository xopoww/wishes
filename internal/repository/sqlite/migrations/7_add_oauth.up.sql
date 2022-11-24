CREATE TABLE OAuth (
    provider    TEXT NOT NULL,
    external_id TEXT NOT NULL,
    user_id     INTEGER NOT NULL,

    FOREIGN KEY(user_id) REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE
    UNIQUE(provider, external_id) ON CONFLICT ABORT
);