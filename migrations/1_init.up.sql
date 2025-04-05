CREATE TABLE IF NOT EXISTS tasks (
    tid varchar(36) NOT NULL PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    status text NOT NULL,
    UNIQUE(title)
);