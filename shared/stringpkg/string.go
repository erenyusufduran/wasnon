package stringpkg

func NullableString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
