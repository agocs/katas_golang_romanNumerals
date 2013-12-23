package romanNumerals

import "testing"

func Test1isI(t *testing.T){
	if toRoman(1) != "I"{
		t.Errorf("toRoman(1) did not return I")
	}
}

func Test2isII(t *testing.T){
	if toRoman(2) != "II"{
		t.Errorf("toRoman(2) did not return II")
	}

}

func Test3isIII(t *testing.T){
	if toRoman(3) != "III"{
		t.Errorf("toRoman(3) did not return III")
	}

}
func Test5isV(t *testing.T){
	if toRoman(5) != "V"{
		t.Errorf("toRoman(5) did not return V")
	}

}
func Test10isX(t *testing.T){
	if toRoman(10) != "X"{
		t.Errorf("toRoman(10) did not return X")
	}

}

func Test25isXXV(t *testing.T){
	if toRoman(25) != "XXV"{
		t.Errorf("toRoman(25) did not return XXV")
	}

}
func Test1738isSomethingLong(t *testing.T){
	if toRoman(1738) != "MDCCXXXVIII"{
		t.Errorf("toRoman(1738) did not return MDCCXXXVIII")
	}

}
func Test944isCMXLIV(t *testing.T){
	if toRoman(944) != "CMXLIV"{
		t.Errorf("toRoman(944) did not return CMXLIV")
	}

}






func testIis1(t *testing.T){
	if i := toArabic("I"); i != 1{
		t.Errorf("toArabic(I) returned %d, not 1")
	}
}



func testIIis2(t *testing.T){
	if i := toArabic("II"); i != 2{
		t.Errorf("toArabic(I) returned %d, not 2")
	}
}


func testIVis4(t *testing.T){
	if i := toArabic("IV"); i != 4{
		t.Errorf("toArabic(I) returned %d, not 4")
	}
}

func testVis5(t *testing.T){
	if i := toArabic("V"); i != 5{
		t.Errorf("toArabic(I) returned %d, not 5")
	}
}


func testXis10(t *testing.T){
	if i := toArabic("X"); i != 10{
		t.Errorf("toArabic(I) returned %d, not 5")
	}
}

func testXCIXis99(t *testing.T){
	if i := toArabic("XCIX"); i != 99{
		t.Errorf("toArabic(I) returned %d, not 99")
	}
}
