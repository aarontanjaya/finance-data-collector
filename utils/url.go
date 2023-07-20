package util

import (
	"net/url"
	"reflect"
)

func BuildURLQuery(u *url.URL, q interface{}) {
	query := reflect.ValueOf(q)
	queryType := query.Type()
	qnya := url.Values{}

	for i := 0; i < query.NumField(); i++ {
		qnya.Set(queryType.Field(i).Name, query.FieldByIndex([]int{i}).String())
	}

	u.RawQuery = qnya.Encode()
}
