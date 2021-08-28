DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts
(
    id         SERIAL PRIMARY KEY,
    author_id  INTEGER REFERENCES authors (id) NOT NULL,
    title      TEXT                            NOT NULL,
    content    TEXT                            NOT NULL,
    created_at BIGINT                          NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Дмитрий');
INSERT INTO authors (id, name) VALUES (1, 'Cyber');
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (0, 'Статья 1', 'Содержание статьи 1', 0);
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (1, 'Статья 2', 'Содержание статьи 2', 0);
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (1, 'Статья 3', 'Содержание статьи 3', 0);
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (0, 'Статья 4', 'Содержание статьи 4', 11);