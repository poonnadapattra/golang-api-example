CREATE SEQUENCE collections_id_seq;
CREATE TABLE collections (
    id int NOT NULL PRIMARY KEY DEFAULT nextval('collections_id_seq'),
    name varchar,
    type varchar,
    metadata json,
    created_date TIMESTAMP DEFAULT NOW()
)
ALTER SEQUENCE collections_id_seq RESTART WITH 1;


CREATE SEQUENCE groups_id_seq;
CREATE TABLE groups(
    id int NOT NULL PRIMARY KEY DEFAULT nextval('groups_id_seq'),
    name varchar,
    type varchar,
    metadata json,
    collection_id int REFERENCES collections(id),
    created_date TIMESTAMP DEFAULT NOW()
)
ALTER SEQUENCE groups_id_seq RESTART WITH 1;

CREATE SEQUENCE items_id_seq;
CREATE TABLE items(
    id int NOT NULL PRIMARY KEY DEFAULT nextval('items_id_seq'),
    name varchar,
    created_date TIMESTAMP DEFAULT NOW()
)


CREATE TABLE groups_items(
    group_id int REFERENCES groups(id),
    items_id int REFERENCES items(id)
)



ALTER SEQUENCE seq RESTART WITH 1;
UPDATE t SET idcolumn=nextval('seq')