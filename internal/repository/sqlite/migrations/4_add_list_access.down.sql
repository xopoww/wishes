PRAGMA foreign_keys=off;

CREATE TABLE IF NOT EXISTS _Lists_new( 
    id          INTEGER PRIMARY KEY,
    
    title       TEXT NOT NULL,
    owner_id    INTEGER,
    FOREIGN KEY(owner_id) REFERENCES Users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO _Lists_new(id, title, owner_id)
SELECT id, title, owner_id
FROM Lists;

DROP TABLE Lists;

ALTER TABLE _Lists_new RENAME TO Lists;

PRAGMA foreign_keys=on;

DROP TABLE ListAccessEnum;