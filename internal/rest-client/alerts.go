package client

type Variables struct {
	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
}

type Alert struct {
	Name          string    `json:"name"`
	Id            string    `json:"id"`
	Action        string    `json:"action"`
	Code          int64     `json:"code"`
	ComponentName string    `json:"component_name"`
	ComponentType string    `json:"component_type"`
	Created       int64     `json:"created"`
	Description   string    `json:"description"`
	Flagged       bool      `json:"flagged"`
	Index         int64     `json:"index"`
	KBurl         string    `json:"knowledge_base_url"`
	Notified      int64     `json:"notified"`
	Severity      string    `json:"severity"`
	State         string    `json:"state"`
	Summary       string    `json:"summary"`
	Updated       int64     `json:"updated"`
	Vars          Variables `json:"variables"`
}

type AlertsList struct {
	CntToken      string  `json:"continuation_token"`
	TotalItemsCnt int     `json:"total_item_count"`
	Items         []Alert `json:"items"`
}

func (fb *FBClient) GetAlerts(filter string) *AlertsList {
	uri := "/alerts"
	result := new(AlertsList)
	req := fb.RestClient.R().SetResult(&result)
	if filter != "" {
		req = req.SetQueryParam("filter", filter)
	}
	res, _ := req.Get(uri)
	if res.StatusCode() == 401 {
		fb.RefreshSession()
		_, _ = req.Get(uri)
	}
	return result
}
