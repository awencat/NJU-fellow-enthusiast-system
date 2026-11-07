package service

import (
	"strings"
	"unicode/utf16"

	"campus-fellow/server/internal/model"
)

// Match browser/TypeScript string length, including surrogate pairs for emoji.
func length(s string, min, max int) bool {
	n := len(utf16.Encode([]rune(s)))
	return n >= min && n <= max
}
func optionalLength(s *string, max int) bool { return s == nil || length(*s, 0, max) }
func trimPointer(s **string) {
	if *s != nil {
		v := strings.TrimSpace(**s)
		*s = &v
	}
}
func nullEmpty(s **string) {
	trimPointer(s)
	if *s != nil && **s == "" {
		*s = nil
	}
}
func contains(value string, options ...string) bool {
	for _, option := range options {
		if value == option {
			return true
		}
	}
	return false
}

func validateExchange(in *model.ExchangeInput) error {
	in.Title = strings.TrimSpace(in.Title)
	in.Content = strings.TrimSpace(in.Content)
	nullEmpty(&in.ImageURL)
	if err := validateTags(&in.Tags, false); err != nil {
		return err
	}
	if !length(in.Title, 4, 100) {
		return fail(400, "标题应为4—100个字符")
	}
	if !length(in.Content, 10, 10000) {
		return fail(400, "正文应为10—10000个字符")
	}
	if !optionalLength(in.ImageURL, 500) {
		return fail(400, "图片地址过长")
	}
	return nil
}
func validateMate(in *model.MateInput) error {
	in.Title = strings.TrimSpace(in.Title)
	in.Details = strings.TrimSpace(in.Details)
	in.Contact = strings.TrimSpace(in.Contact)
	in.Duration = strings.TrimSpace(in.Duration)
	nullEmpty(&in.ImageURL)
	nullEmpty(&in.Location)
	nullEmpty(&in.Notes)
	if !length(in.Title, 4, 100) || !length(in.Details, 10, 10000) {
		return fail(400, "标题应为4—100个字符，详细说明应为10—10000个字符")
	}
	if err := validateTags(&in.Tags, true); err != nil {
		return err
	}
	if !length(in.Contact, 2, 100) || !length(in.Duration, 1, 100) {
		return fail(400, "请填写联系方式和持续时间，且均不超过100个字符")
	}
	if !optionalLength(in.ImageURL, 500) || !optionalLength(in.Location, 100) || !optionalLength(in.Notes, 1000) {
		return fail(400, "图片地址、地点或备注过长")
	}
	return nil
}

func validateTags(tags *[]string, required bool) error {
	if len(*tags) > 6 {
		return fail(400, "最多6个标签")
	}
	clean := []string{}
	seen := map[string]bool{}
	for _, tag := range *tags {
		tag = strings.TrimSpace(tag)
		if !length(tag, 1, 16) {
			return fail(400, "标签应为1—16个字符")
		}
		key := strings.ToLower(tag)
		if !seen[key] {
			clean = append(clean, tag)
			seen[key] = true
		}
	}
	if required && len(clean) == 0 {
		return fail(400, "请填写至少1个标签")
	}
	*tags = clean
	return nil
}
