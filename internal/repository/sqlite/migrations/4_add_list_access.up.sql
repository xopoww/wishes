CREATE TABLE ListAccessEnum (
    N INTEGER PRIMARY KEY,
    S TEXT NOT NULL
);

INSERT INTO ListAccessEnum(N, S) VALUES (0, "public");
INSERT INTO ListAccessEnum(N, S) VALUES (1, "link");
INSERT INTO ListAccessEnum(N, S) VALUES (2, "private");

PRAGMA foreign_keys=off;

CREATE TABLE IF NOT EXISTS _Lists_new( 
    id          INTEGER PRIMARY KEY,
    
    title       TEXT NOT NULL,
    owner_id    INTEGER,

    access      INTEGER NOT NULL,
    
    FOREIGN KEY(owner_id) REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY(access) REFERENCES ListAccessEnum(N) ON DELETE RESTRICT
);

INSERT INTO _Lists_new(id, title, owner_id, access)
SELECT id, title, owner_id, access FROM Lists JOIN (SELECT 2 as access);

DROP TABLE Lists;

ALTER TABLE _Lists_new RENAME TO Lists;

PRAGMA foreign_keys=on;