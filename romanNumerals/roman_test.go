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
	if i := toArabic("I"); i != 1{
		t.Errorf("toArabic(I) returned %d, not 1", i)
	}
}


func TestIIis2(t *testing.T){
	if i := toArabic("II"); i != 2{
		t.Errorf("toArabic(II) returned %d, not 2", i)
	}
}


func TestIVis4(t *testing.T){
	if i := toArabic("IV"); i != 4{
		t.Errorf("toArabic(IV) returned %d, not 4", i)
	}
}

func TestVis5(t *testing.T){
	if i := toArabic("V"); i != 5{
		t.Errorf("toArabic(V) returned %d, not 5", i)
	}
}

func TestVIis6(t*testing.T){
	if i := toArabic("VI"); i != 6{
		t.Errorf("toArabic(VI) returned %d, not 6", i)
	}
}


func TestXis10(t *testing.T){
	if i := toArabic("X"); i != 10{
		t.Errorf("toArabic(X) returned %d, not 10", i)
	}
}

func TestXCIXis99(t *testing.T){
	if i := toArabic("XCIX"); i != 99{
		t.Errorf("toArabic(XCIX) returned %d, not 99", i)
	}
}


func TestCXVIIIis143(t *testing.T){
	if i := toArabic("CXVIII"); i != 143{
		t.Errorf("toArabic(CXVIII) returned %d, not 143", i)
	}
}