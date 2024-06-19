package request

type User struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required,gte=6,lte=50"`
}
