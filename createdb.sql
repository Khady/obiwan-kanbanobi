-- GRANT ALL PRIVILEGES ON DATABASE kanban TO kanban;
SET client_encoding = 'UTF8';

CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    content text NOT NULL,
    column_id integer NOT NULL,
    project_id integer NOT NULL,
    tags text NOT NULL,
    user_id integer NOT NULL,
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
    author_id integer NOT NULL,
    comment_date timestamp NOT NULL
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

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id text NOT NULL,
    ident_date timestamp NOT NULL,
    session_key text NOT NULL
);

INSERT INTO users(id, name, admin, password, mail, active) VALUES(1, 'admin', true, 'admin', 'admin@admin.com', true);
INSERT INTO projects(id, name, admins_id, read, content) VALUES(1, 'Exemple', ',1,', '1', 'Ceci est un projet d''exemple');
INSERT INTO columns(id, name, project_id, content, tags, scripts_id, write) VALUES(1, 'Cliquer sur moi pour voir ma description', 1, 'Voici la description d''une colonne.<br>Vous pouvez creer une nouvelle colonne en cliquant sur le bouton ajoutez une colonne.', '', 0, ',1,');
INSERT INTO cards(name, content, column_id, project_id, tags, user_id, scripts_id, write) VALUES('this is a card', 'This is the description', 1, 1, '', 1, 0, ',1,');