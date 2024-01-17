// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	CreateAssignmentFile(ctx context.Context, arg CreateAssignmentFileParams) (AssignmentFile, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error)
	CreateCourseModule(ctx context.Context, arg CreateCourseModuleParams) (CourseModule, error)
	CreateLecture(ctx context.Context, arg CreateLectureParams) (Lecture, error)
	CreateSubscribe(ctx context.Context, arg CreateSubscribeParams) (Subscribe, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserDetails(ctx context.Context, arg CreateUserDetailsParams) (UserDetail, error)
	CreateUserRole(ctx context.Context, arg CreateUserRoleParams) (UserRole, error)
	DeleteCategory(ctx context.Context, categoryID int64) error
	DeleteCourse(ctx context.Context, courseID int64) error
	DeleteCourseModule(ctx context.Context, moduleID int64) error
	DeleteLecture(ctx context.Context, lectureID int64) error
	DeleteUser(ctx context.Context, userID int64) error
	DeleteUserDetails(ctx context.Context, userDetailsID int64) error
	DeleteUserRole(ctx context.Context, roleID int64) error
	GetCategory(ctx context.Context, categoryID int64) (Category, error)
	GetCategoryForUpdate(ctx context.Context, categoryID int64) (Category, error)
	GetCourse(ctx context.Context, courseID int64) (Course, error)
	GetCourseForUpdate(ctx context.Context, courseID int64) (Course, error)
	GetCourseModule(ctx context.Context, moduleID int64) (CourseModule, error)
	GetLecture(ctx context.Context, lectureID int64) (Lecture, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	GetUserDetails(ctx context.Context, userDetailsID int64) (UserDetail, error)
	GetUserRole(ctx context.Context, roleID int64) (UserRole, error)
	ListCategory(ctx context.Context, arg ListCategoryParams) ([]Category, error)
	ListCourse(ctx context.Context, arg ListCourseParams) ([]Course, error)
	ListLectures(ctx context.Context, arg ListLecturesParams) ([]Lecture, error)
	ListModules(ctx context.Context, arg ListModulesParams) ([]CourseModule, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	ListUserDetails(ctx context.Context, arg ListUserDetailsParams) ([]UserDetail, error)
	ListUserRole(ctx context.Context, arg ListUserRoleParams) ([]UserRole, error)
	UpdateCategory(ctx context.Context, categoryName string) (Category, error)
	UpdateCourse(ctx context.Context, courseName string) (Course, error)
	UpdateCourseModule(ctx context.Context, moduleName string) (CourseModule, error)
	UpdateLecture(ctx context.Context, status string) (Lecture, error)
	UpdateUser(ctx context.Context, email string) (User, error)
	UpdateUserDetails(ctx context.Context, addressLine1 string) (UserDetail, error)
	UpdateUserRole(ctx context.Context, role string) (UserRole, error)
}

var _ Querier = (*Queries)(nil)
