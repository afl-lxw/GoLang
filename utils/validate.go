package utils

import "regexp"

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	// 使用正则表达式匹配邮箱格式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

// ValidatePhone 验证手机号格式
func ValidatePhone(phone string) bool {
	// 使用正则表达式匹配手机号格式
	phoneRegex := `^\d{1,3}\d{6,14}$`
	return regexp.MustCompile(phoneRegex).MatchString(phone)
}
