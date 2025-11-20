
CREATE TABLE IF NOT EXISTS "users" (
                                      id            SERIAL PRIMARY KEY,
                                      username      VARCHAR(255) NOT NULL UNIQUE,
                                      password      VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255),
                                     author_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                                     category VARCHAR(100),
                                     score INTEGER DEFAULT 0,
                                     views INTEGER DEFAULT 0,
                                     type VARCHAR(50),
                                     text TEXT,
                                     upvote_percentage INTEGER DEFAULT 0,
                                     created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
                                        id SERIAL PRIMARY KEY,
                                        post_id INTEGER REFERENCES posts(id),
                                        author INTEGER REFERENCES users(id) ON DELETE CASCADE,
                                        body TEXT,
                                        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS votes (
                                     id SERIAL PRIMARY KEY,
                                     user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
                                     post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
                                     vote INTEGER CHECK (vote IN (-1, 0, 1)),
                                     created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

