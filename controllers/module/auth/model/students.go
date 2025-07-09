package model

type StudentModel struct {
	FirstName    string `json:"firstName" binding:"required"`
	MiddleName   string `json:"middleName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	DateOfBirth  string `json:"dateOfBirth" binding:"required"`
	MatricNumber string `json:"matricNumber" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	Department   string `json:"department" binding:"required"`
	Level        string `json:"level" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email"  binding:"required,email"`
	Role         string `json:"role" binding:"required"`
}
