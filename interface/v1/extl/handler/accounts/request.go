package accounts

type Action struct {
	Id     string `json:"id" validate:"required" example:"65c1de91056ae9755c64ffba"`
	Action int    `json:"action" validate:"required,min=1,max=2" example:"1"`
}

type Get struct {
	Id string `json:"id" validate:"required" example:"65c1de91056ae9755c64ffba" param:"id"`
}

type Premium struct {
	Id string `json:"id" validate:"required" example:"65c1de91056ae9755c64ffba" param:"id"`
}

type List struct {
	Page  int32 `json:"page" validate:"omitempty,min=1" example:"1" query:"page"`
	Limit int32 `json:"limit" validate:"omitempty,min=1" example:"1" query:"limit"`
}
