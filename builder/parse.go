package builder

import (
	"net/url"
	"reflect"
	"strconv"
)

func StructToURLValues(baseURL string, i interface{}) (*url.URL, error) {
	// Create a URL object
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	v := reflect.ValueOf(i)

	// Check if the value is a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Iterate through the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i)

		// Add the field value to the url.Values map
		switch value.Kind() {
		case reflect.String:
			values.Add(tag, value.String())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			values.Add(tag, strconv.FormatUint(value.Uint(), 10))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			values.Add(tag, strconv.FormatInt(value.Int(), 10))
		case reflect.Float32, reflect.Float64:
			values.Add(tag, strconv.FormatFloat(value.Float(), 'f', -1, 64))
		}
	}
	u.RawQuery = values.Encode()

	return u, nil
}
