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

