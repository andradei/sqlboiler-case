CREATE TABLE IF NOT EXISTS persons (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL
);

CREATE TABLE IF NOT EXISTS things (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    bought_on TIMESTAMP
);

CREATE TABLE IF NOT EXISTS persons_things (
    person_id INT NOT NULL REFERENCES persons(id),
    thing_id INT NOT NULL REFERENCES things(id),
    PRIMARY KEY (person_id, thing_id)
);

INSERT INTO persons (id, name, age) VALUES
(1, 'Spiderman', 20),
(2, 'Ironman', 40),
(3, 'Doctor Strange', 44);

INSERT INTO things (id, name, bought_on) VALUES
(1, 'mask', now()),
(2, 'armor', NULL),
(3, 'cape', NULL);

INSERT INTO persons_things (person_id, thing_id) VALUES
(1, 1),
(2, 1),
(2, 2),
(3, 3);
