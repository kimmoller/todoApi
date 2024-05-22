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

INSERT INTO identity(username, password) VALUES('identityToUpdate', '$2a$14$yyVApx6b7RrpeqoyusS7muHjS2Bdv0Vd0ontjgI4SKsEOTYVRoTBa');
INSERT INTO identity(username, password) VALUES('identityToDelete', '$2a$14$uXuTdEy3sJw/RvV2cct9JeGm8BCz5Lxag11MNmYNfTxUnf6VQdXZO');
INSERT INTO identity(username, password) VALUES('todoFetch', '$2a$14$SZmSx5p4LWFAwCwA94uB0uKWE/xEgENrBnJqlqaIPfADy0xA24gQy');
INSERT INTO identity(username, password) VALUES('todoCreate', '$2a$14$2zW/avgk5bM/4taL81SCduBEiN2uChb2sSTeWqQOaWL3vjq7ZZKRS');
INSERT INTO identity(username, password) VALUES('todoUpdate', '$2a$14$oFgaxr3rx8gX4ndGYFQrA.Z71XF49aPf0UhoCgvDgYKBqSN1M8jIy');
INSERT INTO identity(username, password) VALUES('todoDelete', '$2a$14$Is/WX9TxGDIyGVSvEGP.ve0eLHMbagIAJRazA2/9TIyj9j/pcdtty');
INSERT INTO identity(username, password) VALUES('loginTest', '$2a$14$ffZOF.n3TBvPAKYkyeUb3et3OZ9XA.C80rRERhLNY.7xIfMA03Th2');

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
