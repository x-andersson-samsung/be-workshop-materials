# SQL Basics Workshop Exercises

## Exercise 1: Creating a Simple Table
Create a table named `students` with the following columns:
- `id` (integer, primary key)
- `name` (text, not null)
- `age` (integer)
- `grade` (text)

## Exercise 2: Inserting Data
Insert the following student records into the `students` table:
- Student 1: Name "Alice", Age 20, Grade "A"
- Student 2: Name "Bob", Age 22, Grade "B"
- Student 3: Name "Charlie", Age 19, Grade "A"

## Exercise 3: Basic Selection
Write a query to select all columns for all students in the `students` table.

## Exercise 4: Selection with Conditions
Write a query to select all columns for students who are 20 years old or older.

## Exercise 5: Selection with Text Condition
Write a query to select only the names of students who have grade "A".

## Exercise 6: Updating Data
Update the age of the student named "Bob" to 23.

## Exercise 7: Updating with Condition
Change the grade to "B" for all students who are younger than 21.

## Exercise 8: Creating Another Table
Create a table named `courses` with the following columns:
- `id` (integer, primary key)
- `student_id` (integer, references students table)
- `course_name` (text, not null)
- `credits` (integer)

## Exercise 9: Inserting Related Data
Add the following course records:
- Alice (student_id=1) is taking "Mathematics" (3 credits)
- Alice (student_id=1) is taking "Physics" (4 credits)
- Bob (student_id=2) is taking "Chemistry" (3 credits)

## Exercise 10: Deleting Records
Delete all students who have grade "B" from the students table.