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






func TestIis1(t *testing.T){
	if i := ToArabic("I"); i != 1{
		t.Errorf("ToArabic(I) returned %d, not 1", i)
	}
}


func TestIIis2(t *testing.T){
	if i := ToArabic("II"); i != 2{
		t.Errorf("ToArabic(II) returned %d, not 2", i)
	}
}


func TestIVis4(t *testing.T){
	if i := ToArabic("IV"); i != 4{
		t.Errorf("ToArabic(IV) returned %d, not 4", i)
	}
}

func TestVis5(t *testing.T){
	if i := ToArabic("V"); i != 5{
		t.Errorf("ToArabic(V) returned %d, not 5", i)
	}
}

func TestVIis6(t*testing.T){
	if i := ToArabic("VI"); i != 6{
		t.Errorf("ToArabic(VI) returned %d, not 6", i)
	}
}


func TestXis10(t *testing.T){
	if i := ToArabic("X"); i != 10{
		t.Errorf("ToArabic(X) returned %d, not 10", i)
	}
}

func TestXCIXis99(t *testing.T){
	if i := ToArabic("XCIX"); i != 99{
		t.Errorf("ToArabic(XCIX) returned %d, not 99", i)
	}
}


func TestCXVIIIis143(t *testing.T){
	if i := ToArabic("CXVIII"); i != 143{
		t.Errorf("ToArabic(CXVIII) returned %d, not 143", i)
	}
}