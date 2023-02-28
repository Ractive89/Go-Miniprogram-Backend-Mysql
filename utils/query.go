package utils

func AddQuery(url string, params map[string]string) string {
	if params != nil {
		url += "?"
		for key, value := range params {
			url = url + key + "=" + value + "&"
		}
	}
	return url
}
