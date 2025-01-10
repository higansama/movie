package pagination

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	// TotalPages int `json:"total_pages"`
	// Param          string `json:"param"`
	// ParamValue     string `json:"param_value"`
	// SortParam      string `json:"sort_param"`
	// SortParamValue string `json:"sort_param_value"`
}

func (p *Pagination) Validate() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Limit < 0 {
		p.Limit = 1
	}

}
