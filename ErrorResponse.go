package go_googlesearchengine

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Domain  string `json:"domain"`
			Reason  string `json:"reason"`
		} `json:"errors"`
		Status  string `json:"status"`
		Details []struct {
			Type     string `json:"@type"`
			Reason   string `json:"reason,omitempty"`
			Domain   string `json:"domain,omitempty"`
			Metadata struct {
				Consumer        string `json:"consumer"`
				QuotaMetric     string `json:"quota_metric"`
				Service         string `json:"service"`
				QuotaLimit      string `json:"quota_limit"`
				QuotaLimitValue string `json:"quota_limit_value"`
				QuotaLocation   string `json:"quota_location"`
			} `json:"metadata,omitempty"`
			Links []struct {
				Description string `json:"description"`
				Url         string `json:"url"`
			} `json:"links,omitempty"`
		} `json:"details"`
	} `json:"error"`
}
