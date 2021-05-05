package dto

import "Gin/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	StudentID string `json:"studentID"`
	Status    int    `json:"status"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		StudentID: user.StudentID,
		Status:    user.Status,
	}
}
