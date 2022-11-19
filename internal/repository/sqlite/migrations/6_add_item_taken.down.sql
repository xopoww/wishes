PRAGMA foreign_keys=off;

CREATE TABLE IF NOT EXISTS _Items_new( 
    id      INTEGER PRIMARY KEY,

    title   TEXT NOT NULL,
    desc    TEXT NOT NULL DEFAULT "",

    list_id INTEGER,

    FOREIGN KEY(list_id) REFERENCES Lists(id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO _Items_new(id, title, desc, list_id)
SELECT id, title, desc, list_id FROM Items;

DROP TABLE Items;

ALTER TABLE _Items_new RENAME TO Items;

PRAGMA foreign_keys=on;