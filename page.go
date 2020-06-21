package tools

type PageRequest struct {
	PerPage int64 `json:"perPage"`
	Page    int64 `json:"page"`
}

func (i *PageRequest) GetPageInfoNew() *ResultsPageInfo {
	pageInfo := new(ResultsPageInfo)
	if i.PerPage == 0 {
		i.PerPage = 15
	}
	pageInfo.PerPage = i.PerPage
	pageInfo.Page = i.Page
	if pageInfo.Total > 0 {
		pageInfo.TotalPage = pageInfo.Total / pageInfo.PerPage
		pageInfo.CurrentTotal = pageInfo.PerPage
		if supply := pageInfo.Total % pageInfo.PerPage; supply > 0 {
			pageInfo.TotalPage++
			if pageInfo.Page == pageInfo.TotalPage {
				pageInfo.CurrentTotal = supply
			}
		}
	}
	return pageInfo
}

type ResultsPageInfo struct {
	Page         int64
	TotalPage    int64
	Total        int64
	PerPage      int64
	CurrentTotal int64
}

func (i *ResultsPageInfo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"page":         i.Page,
		"totalPage":    i.TotalPage,
		"total":        i.Total,
		"perPage":      i.PerPage,
		"currentTotal": i.CurrentTotal,
	}
}

func (i *ResultsPageInfo) SetByParams(params map[string]interface{}) *ResultsPageInfo {
	for k, v := range params {
		if k == "page" || k == "currentPage" {
			i.Page = StringToInt64(v.(string))
		} else if k == "perPage" || k == "pageSize" {
			i.PerPage = StringToInt64(v.(string))
		}
	}
	if i.Page == 0 {
		i.Page = 1
	}
	return i
}

func PageInfo() *ResultsPageInfo {
	return new(ResultsPageInfo)
}

func (i *ResultsPageInfo) CalcMember() {
	if i.Page == 0 {
		i.Page = 1
	}
	if i.Total > 0 {
		i.TotalPage = i.Total / i.PerPage
		i.CurrentTotal = i.PerPage
		if supply := i.Total % i.PerPage; supply > 0 {
			i.TotalPage++
			if i.Page == i.TotalPage {
				i.CurrentTotal = supply
			} else if i.Page > i.TotalPage {
				i.CurrentTotal = 0
			}
		}
	}
}
