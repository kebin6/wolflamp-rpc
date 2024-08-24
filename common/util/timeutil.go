package util

import "strings"

// NormalizeDateTime 函数用于将日期时间字符串规范化为统一格式
func NormalizeDateTime(dateTimeStr string) string {
	parts := strings.Fields(dateTimeStr)      // 分割日期和时间
	dateParts := strings.Split(parts[0], "-") // 分割日期部分
	timeParts := strings.Split(parts[1], ":") // 分割时间部分

	// 规范化日期部分
	var normalizedDateParts []string
	for i, part := range dateParts {
		if i > 0 && len(part) == 1 { // 对于月份和日期，检查是否需要添加前导零
			normalizedDateParts = append(normalizedDateParts, "0"+part)
		} else {
			normalizedDateParts = append(normalizedDateParts, part)
		}
	}

	// 规范化时间部分
	var normalizedTimeParts []string
	for _, part := range timeParts {
		if len(part) == 1 { // 对于小时、分钟和秒，检查是否需要添加前导零
			normalizedTimeParts = append(normalizedTimeParts, "0"+part)
		} else {
			normalizedTimeParts = append(normalizedTimeParts, part)
		}
	}

	// 重新组合日期时间和时间部分
	return strings.Join(normalizedDateParts, "-") + " " + strings.Join(normalizedTimeParts, ":")
}
