CREATE TABLE Lists (
    id          INTEGER PRIMARY KEY,
    
    title       TEXT NOT NULL,
    owner_id    INTEGER,
    FOREIGN KEY(owner_id) REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE Items (
    id      INTEGER PRIMARY KEY,

    title   TEXT NOT NULL,
    desc    TEXT NOT NULL DEFAULT "",

    list_id INTEGER,
    FOREIGN KEY(list_id) REFERENCES Lists(id) ON DELETE CASCADE ON UPDATE CASCADE
);