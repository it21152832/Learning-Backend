// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: assignment.sql

package db

import (
	"context"
)

const createAssignment = `-- name: CreateAssignment :one
INSERT INTO "assignment"(
    user_id,
    course_module_id,
    lecture_id
) VALUES (
    $1, $2, $3
)
RETURNING assignment_id, user_id, course_module_id, lecture_id
`

type CreateAssignmentParams struct {
	UserID         int64 `json:"user_id"`
	CourseModuleID int64 `json:"course_module_id"`
	LectureID      int64 `json:"lecture_id"`
}

func (q *Queries) CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, createAssignment, arg.UserID, arg.CourseModuleID, arg.LectureID)
	var i Assignment
	err := row.Scan(
		&i.AssignmentID,
		&i.UserID,
		&i.CourseModuleID,
		&i.LectureID,
	)
	return i, err
}
