package validate

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValidateStruct(obj interface{}) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("validate")
		comment := field.Tag.Get("comment")
		if strings.Contains(tag, "required") && value.String() == "" {
			return fmt.Errorf("%s不能为空", comment)
		}
		if strings.Contains(tag, "email") && !strings.Contains(value.String(), "@") {
			return fmt.Errorf("%s不是有效的邮箱", comment)
		}
		if strings.Contains(tag, "min") {
			minValue := strings.Split(tag, "=")[1]
			min, err := strconv.Atoi(minValue)
			if err != nil {
				return fmt.Errorf("无效的最小值：%s", comment)
			}
			if value.Int() < int64(min) {
				return fmt.Errorf("%s必须大于等于%d", comment, min)
			}
		}
	}
	return nil
}
