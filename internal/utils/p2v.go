package utils

func P2VString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func P2VInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func P2VInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}
