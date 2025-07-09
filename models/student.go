package models

import "time"

type Faculty struct {
	FacultyID   uint `gorm:"primaryKey"`
	Faculty     string
	Departments []Department `gorm:"foreignKey:FacultyID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Department struct {
	DepartmentID uint `gorm:"primaryKey"`
	Department   string
	FacultyID    uint
	Faculty      Faculty
	Students     []Student `gorm:"foreignKey:DepartmentID"`
	Courses      []Course  `gorm:"foreignKey:DepartmentID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Level struct {
	LevelID   uint `gorm:"primaryKey"`
	Level     string
	Students  []Student `gorm:"foreignKey:LevelID"`
	Courses   []Course  `gorm:"foreignKey:LevelID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Student struct {
	ID           uint `gorm:"primaryKey"`
	FirstName    string
	MiddleName   string
	LastName     string
	DateOfBirth  string
	MatricNumber string `gorm:"not null;unique"`
	Gender       string
	PhoneNumber  string `gorm:"not null;unique"`
	DepartmentID uint
	Department   Department
	LevelID      uint
	Level        Level
	Password     string
	Email        string `gorm:"not null;unique"`
	Role         string
	Scores       []Score `gorm:"foreignKey:StudentID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Course struct {
	CourseID     uint `gorm:"primaryKey"`
	Course       string
	CourseCode   string
	DepartmentID uint
	Department   Department
	LevelID      uint
	Level        Level
	Scores       []Score `gorm:"foreignKey:CourseID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Grade struct {
	ID        uint `gorm:"primaryKey"`
	Grade     string
	Scores    []Score `gorm:"foreignKey:GradeID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Score struct {
	ID        uint `gorm:"primaryKey"`
	Score     uint
	StudentID uint
	Student   Student
	CourseID  uint
	Course    Course
	GradeID   uint
	Grade     Grade
	CreatedAt time.Time
	UpdatedAt time.Time
}
