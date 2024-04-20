package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name string, description string, categoryID string) (Course, error) {
	newRandomUuid, _ := uuid.NewRandom()
	id := newRandomUuid.String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) GetAllCourses() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil

}
