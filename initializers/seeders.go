package initializers

import (
	"github.com/arbeeorlar/models"
	"gorm.io/gorm"
)

func seedLevel(db *gorm.DB) {
	var count int64
	db.Model(&models.Level{}).Count(&count)

	if count == 0 {
		levels := []models.Level{
			{Level: "100"},
			{Level: "200"},
			{Level: "300"},
			{Level: "400"},
		}
		db.Create(&levels)
		println("✅ Seeded Levels")
	}
}

func seedGrade(db *gorm.DB) {
	var count int64
	db.Model(&models.Grade{}).Count(&count)
	if count == 0 {
		grade := []models.Grade{
			{Grade: "A"},
			{Grade: "B"},
			{Grade: "C"},
			{Grade: "D"},
			{Grade: "E"},
			{Grade: "F"},
		}
		db.Create(&grade)
		println("✅ Seeded Grade")
	}
}

func seedFaculty(db *gorm.DB) {
	var count int64
	db.Model(&models.Faculty{}).Count(&count)
	if count == 0 {
		faculties := []models.Faculty{
			{Faculty: "Faculty of Science"},
			{Faculty: "Faculty of Engineering"},
			{Faculty: "Faculty of Arts"},
			{Faculty: "Faculty of Rural Agriculture"},
			{Faculty: "Faculty of Animal Science"},
			{Faculty: "Faculty of Environmental Science"},
			{Faculty: "Faculty of Medicine"},
			{Faculty: "Faculty of Plant Science"},
			{Faculty: "Faculty of Management Science"},
		}
		db.Create(&faculties)
		println("✅ Seeded Faculties")
	}
}

func seedDepartment(db *gorm.DB) {
	var count int64
	db.Model(&models.Department{}).Count(&count)

	if count == 0 {
		departments := []models.Department{
			{Department: "Computer Science", FacultyID: 1},
			{Department: "Mechanical Engineering", FacultyID: 2},
			{Department: "AEFM", FacultyID: 4},
			{Department: "AERD", FacultyID: 4},
			{Department: "Animal Nutrition", FacultyID: 5},
			{Department: "EMT", FacultyID: 6},
			{Department: "MBBS", FacultyID: 7},
			{Department: "PPCP", FacultyID: 8},
			{Department: "ACCOUNTING", FacultyID: 9},
		}
		db.Create(&departments)
		println("✅ Seeded Departments")
	}
}

func SeedAll(db *gorm.DB) {
	seedLevel(db)
	seedFaculty(db)
	seedDepartment(db)
	seedGrade(db)
}
