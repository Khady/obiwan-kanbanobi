-- GRANT ALL PRIVILEGES ON DATABASE kanban TO kanban;
SET client_encoding = 'UTF8';

CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    content text NOT NULL,
    column_id integer NOT NULL,
    project_id integer NOT NULL,
    tags text NOT NULL,
    users_id text NOT NULL,
    scripts_id text NOT NULL,
    write text NOT NULL
);

CREATE TABLE columns (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    project_id integer NOT NULL,
    content text NOT NULL,
    tags text NOT NULL,
    scripts_id text NOT NULL,
    write text NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    admin boolean NOT NULL DEFAULT(false),
    password text NOT NULL,
    mail text NOT NULL,
    active boolean NOT NULL DEFAULT(false)
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    admins_id text NOT NULL,
    read text NOT NULL,
    content text NOT NULL
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content text NOT NULL,
    cards_id integer NOT NULL,
    author_id integer NOT NULL
);

CREATE TABLE history (
    id SERIAL PRIMARY KEY,
    change_type integer NOT NULL,
    object_id integer NOT NULL,
    column_name text NOT NULL,
    content text NOT NULL
);

CREATE TABLE scripts (
    id SERIAL PRIMARY KEY,
    script_type integer NOT NULL,
    object_id integer NOT NULL,
    filename text NOT NULL
);

CREATE TABLE metadata (
    id SERIAL PRIMARY KEY,
    object_type integer NOT NULL,
    object_id integer NOT NULL,
    data_key text NOT NULL,
    data_value text NOT NULL
);
