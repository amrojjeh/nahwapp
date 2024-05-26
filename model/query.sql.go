// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package model

import (
	"context"
)

const createQuiz = `-- name: CreateQuiz :one

INSERT INTO quiz (
    name, data, hidden, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, name, data, hidden, created, updated
`

type CreateQuizParams struct {
	Name   string
	Data   []byte
	Hidden bool
}

// *****
// QUIZ TABLE
// *****
func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, createQuiz, arg.Name, arg.Data, arg.Hidden)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Data,
		&i.Hidden,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createQuizSession = `-- name: CreateQuizSession :one

INSERT INTO quiz_session (
    student_id, quiz_id, active, questions_answered, created, updated
) VALUES (
    ?, ?, ?, 0, datetime("now"), datetime("now")
) RETURNING id, student_id, quiz_id, questions_answered, active, created, updated
`

type CreateQuizSessionParams struct {
	StudentID int
	QuizID    int
	Active    bool
}

// *****
// QUIZ_SESSION TABLE
// *****
func (q *Queries) CreateQuizSession(ctx context.Context, arg CreateQuizSessionParams) (QuizSession, error) {
	row := q.db.QueryRowContext(ctx, createQuizSession, arg.StudentID, arg.QuizID, arg.Active)
	var i QuizSession
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.QuizID,
		&i.QuestionsAnswered,
		&i.Active,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createStudent = `-- name: CreateStudent :one

INSERT INTO student (
    username, class_code, created, updated
) VALUES (
    ?, ?, datetime("now"), datetime("now")
) RETURNING id, username, class_code, created, updated
`

type CreateStudentParams struct {
	Username  string
	ClassCode string
}

// *****
// STUDENT TABLE
// *****
func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent, arg.Username, arg.ClassCode)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ClassCode,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const createTagAttempt = `-- name: CreateTagAttempt :one

INSERT INTO tag_attempt (
    student_id, tag, correct, created, updated
) VALUES (
    ?, ?, ?, datetime("now"), datetime("now")
) RETURNING id, student_id, tag, correct, created, updated
`

type CreateTagAttemptParams struct {
	StudentID int
	Tag       string
	Correct   bool
}

// *****
// TAG_ATTEMPT TABLE
// *****
func (q *Queries) CreateTagAttempt(ctx context.Context, arg CreateTagAttemptParams) (TagAttempt, error) {
	row := q.db.QueryRowContext(ctx, createTagAttempt, arg.StudentID, arg.Tag, arg.Correct)
	var i TagAttempt
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.Tag,
		&i.Correct,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const deleteQuiz = `-- name: DeleteQuiz :exec
DELETE FROM quiz
WHERE id=?
`

func (q *Queries) DeleteQuiz(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteQuiz, id)
	return err
}

const deleteQuizSession = `-- name: DeleteQuizSession :exec
DELETE FROM quiz_session
WHERE id=?
`

func (q *Queries) DeleteQuizSession(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteQuizSession, id)
	return err
}

const deleteStudent = `-- name: DeleteStudent :exec
DELETE FROM student
WHERE id=?
`

func (q *Queries) DeleteStudent(ctx context.Context, id int) error {
	_, err := q.db.ExecContext(ctx, deleteStudent, id)
	return err
}

const getActiveQuizSession = `-- name: GetActiveQuizSession :one
SELECT id, student_id, quiz_id, questions_answered, active, created, updated FROM quiz_session
WHERE active=true AND student_id=? AND quiz_id=?
LIMIT 1
`

type GetActiveQuizSessionParams struct {
	StudentID int
	QuizID    int
}

func (q *Queries) GetActiveQuizSession(ctx context.Context, arg GetActiveQuizSessionParams) (QuizSession, error) {
	row := q.db.QueryRowContext(ctx, getActiveQuizSession, arg.StudentID, arg.QuizID)
	var i QuizSession
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.QuizID,
		&i.QuestionsAnswered,
		&i.Active,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getQuiz = `-- name: GetQuiz :one
SELECT id, name, data, hidden, created, updated FROM quiz
WHERE id=?
`

func (q *Queries) GetQuiz(ctx context.Context, id int) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, getQuiz, id)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Data,
		&i.Hidden,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getQuizSession = `-- name: GetQuizSession :one
SELECT id, student_id, quiz_id, questions_answered, active, created, updated FROM quiz_session
WHERE id=?
`

func (q *Queries) GetQuizSession(ctx context.Context, id int) (QuizSession, error) {
	row := q.db.QueryRowContext(ctx, getQuizSession, id)
	var i QuizSession
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.QuizID,
		&i.QuestionsAnswered,
		&i.Active,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getStudent = `-- name: GetStudent :one
SELECT id, username, class_code, created, updated FROM student
WHERE id=?
`

func (q *Queries) GetStudent(ctx context.Context, id int) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ClassCode,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const getStudentByUsernameAndClassCode = `-- name: GetStudentByUsernameAndClassCode :one
SELECT id, username, class_code, created, updated FROM student
WHERE username=? AND class_code=?
`

type GetStudentByUsernameAndClassCodeParams struct {
	Username  string
	ClassCode string
}

func (q *Queries) GetStudentByUsernameAndClassCode(ctx context.Context, arg GetStudentByUsernameAndClassCodeParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByUsernameAndClassCode, arg.Username, arg.ClassCode)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ClassCode,
		&i.Created,
		&i.Updated,
	)
	return i, err
}

const listQuiz = `-- name: ListQuiz :many
SELECT id, name, data, hidden, created, updated FROM quiz
WHERE name LIKE ?
ORDER BY created
LIMIT ? OFFSET ?
`

type ListQuizParams struct {
	Name   string
	Limit  int64
	Offset int64
}

func (q *Queries) ListQuiz(ctx context.Context, arg ListQuizParams) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, listQuiz, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quiz
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Data,
			&i.Hidden,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStudents = `-- name: ListStudents :many
SELECT id, username, class_code, created, updated FROM student
WHERE username LIKE ? AND class_code LIKE ?
LIMIT ? OFFSET ?
`

type ListStudentsParams struct {
	Username  string
	ClassCode string
	Limit     int64
	Offset    int64
}

func (q *Queries) ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, listStudents,
		arg.Username,
		arg.ClassCode,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.ClassCode,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTagAttemptsByStudent = `-- name: ListTagAttemptsByStudent :many
SELECT id, student_id, tag, correct, created, updated FROM tag_attempt
WHERE student_id=? AND tag=?
ORDER BY created
LIMIT ? OFFSET ?
`

type ListTagAttemptsByStudentParams struct {
	StudentID int
	Tag       string
	Limit     int64
	Offset    int64
}

func (q *Queries) ListTagAttemptsByStudent(ctx context.Context, arg ListTagAttemptsByStudentParams) ([]TagAttempt, error) {
	rows, err := q.db.QueryContext(ctx, listTagAttemptsByStudent,
		arg.StudentID,
		arg.Tag,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TagAttempt
	for rows.Next() {
		var i TagAttempt
		if err := rows.Scan(
			&i.ID,
			&i.StudentID,
			&i.Tag,
			&i.Correct,
			&i.Created,
			&i.Updated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuizSession = `-- name: UpdateQuizSession :one
UPDATE quiz_session
SET active = ?, questions_answered = ?
WHERE id=?
RETURNING id, student_id, quiz_id, questions_answered, active, created, updated
`

type UpdateQuizSessionParams struct {
	Active            bool
	QuestionsAnswered int
	ID                int
}

func (q *Queries) UpdateQuizSession(ctx context.Context, arg UpdateQuizSessionParams) (QuizSession, error) {
	row := q.db.QueryRowContext(ctx, updateQuizSession, arg.Active, arg.QuestionsAnswered, arg.ID)
	var i QuizSession
	err := row.Scan(
		&i.ID,
		&i.StudentID,
		&i.QuizID,
		&i.QuestionsAnswered,
		&i.Active,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
