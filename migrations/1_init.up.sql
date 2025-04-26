CREATE TABLE IF NOT EXISTS tasks (
    tid varchar(36) NOT NULL PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    status text NOT NULL,
    UNIQUE(title)
);

CREATE TABLE IF NOT EXISTS users (
    uid varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    UNIQUE(email)
);