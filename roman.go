package romanNumerals

func toRoman(x int) string {

	switch {
		case x >= 1000:
			return "M" + toRoman(x-1000)
		case x >= 500:
			return "D" + toRoman(x-500)
		case x >= 100:
			return "C" + toRoman(x-100)
		case x >= 50:
			return "L" + toRoman(x-50)
		case x >= 10:
			return "X" + toRoman(x-10)
		case x >= 5:
			return "V" + toRoman(x-5)
		case x >= 1:
			return "I" + toRoman(x-1)
	}
	return ""
}
