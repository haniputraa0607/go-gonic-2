package dto

// UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required,min=1"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"min=6"`
}
