package main

func mask(s string) string {
	masked := []byte(s)
	toMask := false

	for i := 0; i < len(s); i++ {
		if toMask {
			masked[i] = '*'
		}

		if s[i] == ':' {
			toMask = true
		}
	}

	return string(masked)
}
