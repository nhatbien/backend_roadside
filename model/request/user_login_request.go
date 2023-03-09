package request

type UserLoginRequest struct {
	Phone    string `json:"phone,omitempty" db:"phone, omitempty" validate:"required"`
	Password string `json:"password,omitempty" db:"password, omitempty" validate:"required"`
}
