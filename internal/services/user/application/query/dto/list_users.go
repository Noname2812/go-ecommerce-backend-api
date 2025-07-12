package userquerydto

type ListUsersRequest struct {
	Page     int    `json:"-" form:"page,default=1" binding:"min=1"`
	PageSize int    `json:"-" form:"page_size,default=20" binding:"min=1,max=100"`
	Purpose  string `json:"-" form:"purpose,omitempty"`
	Status   string `json:"-" form:"status,omitempty"`
	Search   string `json:"-" form:"search,omitempty"`
}

type ListUsersResponse struct {
	Users      []UserInfoQueryResponse `json:"users"`
	Total      int64                   `json:"total"`
	Page       int                     `json:"page"`
	PageSize   int                     `json:"page_size"`
	TotalPages int                     `json:"total_pages"`
}
