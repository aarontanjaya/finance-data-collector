package util

import (
	"fmt"
	"net/url"
	"reflect"
)

func BuildURLQuery(u *url.URL, q interface{}) {
	query := reflect.ValueOf(q)
	queryType := query.Type()
	qnya := url.Values{}

	for i := 0; i < query.NumField(); i++ {
		qnya.Set(queryType.Field(i).Tag.Get("json"), fmt.Sprintf("%v", query.FieldByIndex([]int{i})))
	}

	u.RawQuery = qnya.Encode()
}
