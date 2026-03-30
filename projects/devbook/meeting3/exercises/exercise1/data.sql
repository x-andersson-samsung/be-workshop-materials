CREATE SCHEMA exercise2;

SET search_path TO exercise2;

CREATE TABLE students
(
    id    INTEGER PRIMARY KEY,
    name  TEXT    NOT NULL,
    age   INTEGER NOT NULL,
    grade CHAR
);

INSERT INTO students (id, name, age, grade)
VALUES (0, 'Alice', 20, 'A'),
       (1, 'Bob', 22, 'C'),
       (2, 'Charlie', 19, null),
       (3, 'Damian', 20, null),
       (4, 'Elsa', 21, 'B'),
       (5, 'Arnold', 19, null),
       (6, 'Frank', 22, 'D')
;

