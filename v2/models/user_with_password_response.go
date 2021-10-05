package models

type UserWithPasswordResponse struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
	UserID   *int64  `json:"userId"`
}
