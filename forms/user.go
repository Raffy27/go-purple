package forms

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type CreateForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`

	Email string `form:"email" binding:"required"`
}
