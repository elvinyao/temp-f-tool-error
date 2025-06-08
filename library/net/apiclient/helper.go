package apiclient

import (
	"net/url"
	"path"
)

func GenerateURL(baseURL string, elems ...string) (urlResult string) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	for _, elem := range elems {
		u.Path = path.Join(u.Path, elem)
	}

	return u.String()
}

func FindMapKeysForValue(m map[string]string, value string) []string {
	var keys []string
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
		}
	}
	return keys
}
