package request

import "context"

type BaseRequest struct {
	Ctx      context.Context `json:"ctx"`
	Language string          `json:"language"`
}

type BasePortalRequest struct {
	BaseID      int64  `json:"base_id"`
	Phone       string `json:"phone"`
	Role        int    `json:"role_id"`
	CompanyName string `json:"company_name"`
}

type BasePagination struct {
	CurrentPage int `json:"current_page" validate:"required,min=1"`
	PageCount   int `json:"page_count" validate:"required,max=50"`
}

type Test struct {
	BaseRequest
	BasePortalRequest
}
