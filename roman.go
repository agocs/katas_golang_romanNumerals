package romanNumerals

func toRoman(x int) string {

	switch {
		case x >= 1000:
			return "M" + toRoman(x-1000)
		case x >= 900:
			return "CM" + toRoman(x-900)
		case x >= 500:
			return "D" + toRoman(x-500)
		case x >= 400:
			return "CD" + toRoman(x-400)
		case x >= 100:
			return "C" + toRoman(x-100)
		case x >= 90:
			return "XC" + toRoman(x-90)
		case x >= 50:
			return "L" + toRoman(x-50)
		case x >= 40:
			return "XL" + toRoman(x-40)
		case x >= 10:
			return "X" + toRoman(x-10)
		case x >= 9:
			return "IX" + toRoman(x-9)
		case x >= 5:
			return "V" + toRoman(x-5)
		case x >= 4:
			return "IV" + toRoman(x-4)
		case x >= 1:
			return "I" + toRoman(x-1)
	}
	return ""
}
