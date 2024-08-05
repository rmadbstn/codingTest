package programmers

func SolutionHiddenPhoneNumber(phone_number string) string {

	b := ""

	for idx, a := range phone_number {
		if idx >= len(phone_number)-4 {
			b += string(a)
		} else {
			b += "*"
		}
	}

	return b

}
