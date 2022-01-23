package mysql

const (
	sqlCourseTable = "course"
)

type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
