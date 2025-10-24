package utils

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

func HttpPost(url string, msg []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(msg))
	if err != nil {
		return []byte{}, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return body, nil
}

func StructToValues(v interface{}) url.Values {
	values := url.Values{}
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Name
		}

		value := rv.Field(i).Interface()
		values.Set(key, value.(string))
	}

	return values
}
