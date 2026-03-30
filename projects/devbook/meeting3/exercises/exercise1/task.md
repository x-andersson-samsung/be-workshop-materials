# Exercise 1

Load provided test data.

```
psql -u{Username} -p < data.sql
```

You should now have new schema `exercise1` with data for this exercise.

You should have a table with columns:

- `id` (integer, primary key)
- `name` (text, not null)
- `age` (integer)
- `grade` (text)

# Tasks

List all students.
List all students older than 20.
List all students with grade 'B'.
List all students whose name starts with 'A'.
List all students without a grade.

Update 'Bob' age to 21.
Update 'Damian' grade to 'C'.

Add new student `Gandalf` with age 23 and grade 'A'

Create new table `courses` with columns:
- `id` (integer, primary key)
- `student_id` (integer, references students table)
- `course_name` (text, not null)
- `credits` (integer)

Add following course records:
- Alice (student_id=1) is taking "Mathematics" (3 credits)
- Alice (student_id=1) is taking "Physics" (4 credits)
- Bob (student_id=2) is taking "Chemistry" (3 credits)

List the sum of all credits for each student

Delete all students who have grade 'D' from students table.