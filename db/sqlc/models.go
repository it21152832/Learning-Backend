// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

type Assignment struct {
	AssignmentID   int64         `json:"assignment_id"`
	UserID         int64 `json:"user_id"`
	CourseModuleID int64 `json:"course_module_id"`
	LectureID      int64 `json:"lecture_id"`
}

type AssignmentFile struct {
	AssignmentFileID int64         `json:"assignment_file_id"`
	AssignmentID     int64 `json:"assignment_id"`
	AssignmentLink   string        `json:"assignment_link"`
}

type Category struct {
	CategoryID   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryDesc string `json:"category_desc"`
}

type Course struct {
	CourseID   int64  `json:"course_id"`
	CourseName string `json:"course_name"`
	CourseDesc string `json:"course_desc"`
	Category   string `json:"category"`
}

type CourseModule struct {
	ModuleID   int64         `json:"module_id"`
	CourseID   int64 `json:"course_id"`
	ModuleName string        `json:"module_name"`
}

type Lecture struct {
	LectureID      int64         `json:"lecture_id"`
	CourseModuleID int64 `json:"course_module_id"`
	LectureDesc    string        `json:"lecture_desc"`
	LectureNumber  int32 `json:"lecture_number"`
	VideoURL       string        `json:"video_URL"`
	Status         string        `json:"status"`
}

type Subscribe struct {
	SubscribeID int64         `json:"subscribe_id"`
	UserID      int64 `json:"user_id"`
	CourseID    int64 `json:"course_id"`
}

type User struct {
	UserID         int64  `json:"user_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	Role           string `json:"role"`
	Username       string `json:"username"`
}

type UserDetail struct {
	UserDetailsID int64         `json:"user_details_id"`
	UserID        int64 `json:"user_id"`
	Phone         int32         `json:"phone"`
	AddressLine1  string        `json:"address_line1"`
	AddressLine2  string        `json:"address_line2"`
}

type UserRole struct {
	RoleID int64  `json:"role_id"`
	Role   string `json:"role"`
}
