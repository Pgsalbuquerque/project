package dto

type AddUserDTO struct {
	ProjectId int32  `json:"project_id"`
	UserId    string `json:"user_id"`
}
