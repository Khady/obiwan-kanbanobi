GRANT ALL PRIVILEGES ON DATABASE kanban TO kanban;
SET client_encoding = 'UTF8';

CREATE TABLE cards (
    id integer NOT NULL,
    name text NOT NULL,
    desc text NOT NULL,
    column_id integer NOT NULL,
    project_id integer NOT NULL,
    tags text NOT NULL,
    users_id integer NOT NULL,
    scripts_id integer NOT NULL,
    write text NOT NULL,
    origin timestamp NOT NULL
);

CREATE TABLE columns (
    id integer NOT NULL,
    name text NOT NULL,
    project_id integer NOT NULL,
    desc text NOT NULL,
    tags text NOT NULL,
    scripts_id integer NOT NULL,
    write text NOT NULL
);

CREATE TABLE users (
    id integer NOT NULL,
    name text NOT NULL,
    admin boolean NOT NULL DEFAULT(false),
    password string NOT NULL,
    mail string NOT NULL,
    active boolean NOT NULL DEFAULT(false)
);

CREATE TABLE projects (
    id integer NOT NULL,
    name text NOT NULL,
    admins_id text NOT NULL,
    read text NOT NULL
);

CREATE TABLE comments (
    id integer NOT NULL,
    content text NOT NULL,
    cards_id integer NOT NULL,
    author_id integer NOT NULL
);

CREATE TABLE history (
    id integer NOT NULL,
    type integer NOT NULL,
    object_id integer NOT NULL,
    column_name text NOT NULL,
    content text NOT NULL
);

CREATE TABLE scripts (
    id integer NOT NULL,
    type integer NOT NULL,
    object_id integer NOT NULL,
    filename text NOT NULL
);

CREATE SEQUENCE cards_id_seq;
CREATE SEQUENCE columns_id_seq;
CREATE SEQUENCE users_id_seq;
CREATE SEQUENCE projects_id_seq;
CREATE SEQUENCE comments_id_seq;
CREATE SEQUENCE history_id_seq;
CREATE SEQUENCE scripts_id_seq;