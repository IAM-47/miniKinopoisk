-- miniKinopoisk: initial schema
-- Run once against an empty database to bootstrap the project.

CREATE TABLE IF NOT EXISTS users (
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL,
    password_hash TEXT         NOT NULL,
    role          VARCHAR(50)  NOT NULL DEFAULT 'user',
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    CONSTRAINT user_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS movies (
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(255) NOT NULL,
    producer     VARCHAR(255) NOT NULL,
    director     VARCHAR(255) NOT NULL,
    release_year INT
);

CREATE TABLE IF NOT EXISTS actors (
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100)   NOT NULL,
    last_name  VARCHAR(100)   NOT NULL,
    birth_date DATE,
    salary     NUMERIC(15, 2) NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS movie_actor (
    id_movie INT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    id_actor INT NOT NULL REFERENCES actors(id) ON DELETE CASCADE,
    PRIMARY KEY (id_movie, id_actor)
);

CREATE TABLE IF NOT EXISTS budget_and_fees (
    id                    SERIAL PRIMARY KEY,
    id_movie              INT            NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    total_budget          NUMERIC(20, 2) NOT NULL DEFAULT 0,
    fees_in_prod_country  NUMERIC(20, 2) NOT NULL DEFAULT 0,
    fees_in_other         NUMERIC(20, 2) NOT NULL DEFAULT 0,
    CONSTRAINT budget_movie_key UNIQUE (id_movie)
);
