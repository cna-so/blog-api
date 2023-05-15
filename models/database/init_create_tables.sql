CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(50),
    last_name  VARCHAR(80),
    password   VARCHAR(100) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE category
(
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL ,
    updated_at TIMESTAMP NOT NULL 

);
CREATE TABLE articles
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(200) NOT NULL,
    description   TEXT         NOT NULL,
    creator       INTEGER REFERENCES users (id),
    category_id   INTEGER REFERENCES category (id),
    created_at    TIMESTAMP NOT NULL ,
    updated_at    TIMESTAMP NOT NULL
);
-- TODO add other tables like comments
