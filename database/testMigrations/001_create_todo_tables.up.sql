CREATE TABLE IF NOT EXISTS identity (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    identityId INT REFERENCES identity(id)
);

CREATE INDEX IF NOT EXISTS todo_id_identity_id ON todo(id, identityId);

INSERT INTO identity(username, password) VALUES('identityToUpdate', 'originalPass');
INSERT INTO identity(username, password) VALUES('identityToDelete', 'identityToDelete');
INSERT INTO identity(username, password) VALUES('todoFetch', 'todoFetch');
INSERT INTO identity(username, password) VALUES('todoCreate', 'todoCreate');
INSERT INTO identity(username, password) VALUES('todoUpdate', 'todoUpdate');
INSERT INTO identity(username, password) VALUES('todoDelete', 'todoDelete');

INSERT INTO todo(task, status, identityId) VALUES('firstTodo', 'ONGOING',
 (SELECT id FROM identity WHERE username='todoFetch'));
INSERT INTO todo(task, status, identityId) VALUES('secondTodo', 'ONGOING',
 (SELECT id FROM identity WHERE username='todoFetch'));
INSERT INTO todo(task, status, identityId) VALUES('completedTodo', 'COMPLETED',
 (SELECT id FROM identity WHERE username='todoFetch'));

INSERT INTO todo(task, status, identityId) VALUES('todoToUpdate', 'ONGOING',
 (SELECT id FROM identity WHERE username='todoUpdate'));

INSERT INTO todo(task, status, identityId) VALUES('todoToDelete', 'ONGOING',
 (SELECT id FROM identity WHERE username='todoDelete'));
