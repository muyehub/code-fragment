package syntax

import (
	"strconv"
	"time"
)

// AssertTypeToInt32 类型断言转换为 int32 类型
func AssertTypeToInt32(val interface{}) int32 {
	switch val.(type) {
	case int:
		return int32(val.(int))
	case int32:
		return val.(int32)
	case int64:
		return int32(val.(int64))
	case float32:
		return int32(val.(float32))
	case float64:
		return int32(val.(float64))
	case string:
		uidTemp, _ := strconv.Atoi(val.(string))
		return int32(uidTemp)
	default:
		return 0
	}
}

// AssertTypeToString 类型断言转换为 string 类型
func AssertTypeToString(val interface{}) string {
	switch val.(type) {
	case int:
		return strconv.Itoa(val.(int))
	case int32:
		return strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', 1, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', 1, 64)
	case string:
		return val.(string)
	case time.Time:
		return val.(time.Time).Format("2006-01-02 15:04:05")
	default:
		return ""
	}
}

