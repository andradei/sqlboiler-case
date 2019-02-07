CREATE TABLE IF NOT EXISTS persons (
    id SERIAL UNIQUE PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL
);

CREATE TABLE IF NOT EXISTS things (
    id SERIAL UNIQUE PRIMARY KEY,
    name TEXT NOT NULL,
    bought_on TIMESTAMP
);

CREATE TABLE IF NOT EXISTS persons_things (
    person_id INT NOT NULL REFERENCES persons(id),
    thing_id INT NOT NULL REFERENCES things(id),
    PRIMARY KEY (person_id, thing_id)
);

INSERT INTO persons (name, age) VALUES
('Spiderman', 20),
('Ironman', 40),
('Doctor Strange', 44);

INSERT INTO things (name, bought_on) VALUES
('mask', now()),
('armor', NULL),
('cape', NULL);

INSERT INTO persons_things (person_id, thing_id) VALUES
(1, 1),
(2, 1),
(2, 2),
(3, 3);