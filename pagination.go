package zoodmall_sdk

type Pagination struct {
	Page        uint `json:"page"`
	RowsPerPage uint `json:"rowsPerPage"`
	TotalCount  uint `json:"totalCount"`
}
