CREATE TABLE comments
(
    id        SERIAL PRIMARY KEY,
    content   TEXT,
    user_name TEXT,
    post_id   INTEGER REFERENCES post (id)
);

