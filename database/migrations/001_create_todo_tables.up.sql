CREATE TABLE IF NOT EXISTS identity (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    identityId INT REFERENCES identity(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS todo_id_identity_id ON todo(id, identityId);