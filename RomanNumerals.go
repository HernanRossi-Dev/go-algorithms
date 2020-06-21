package main
import (
	"strings"
	"fmt"
)

func intToRoman(num int) string {
		if num == 0 {
			return ""
		}
		
		var result [4]string
		ones := [9]string {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		tens := [9]string {"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
		hunds := [9]string {"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
		thous := [3]string {"M", "MM", "MMM"}

		var one int = num % 10
		if one > 0 {
			result[3] = ones[one -1]
		}
		num = num / 10
		var ten int = num % 10
		if ten > 0 {
			result[2] = tens[ten -1]
		}
		num = num / 10
		var hund int = num % 10
		if hund > 0 {
			result[1] = hunds[hund-1]
		}
		num = num / 10.0
		var thou int = num % 10
		if thou > 0 {
			result[0] = thous[thou-1]
		}
		return strings.Join(result[:], "")
}

type test struct {
	input int
	expected string
}

func main() {
	tests := []test{}
	testOne := test{input: 58, expected: "LVIII"  }
	tests = append(tests, testOne)
	testTwo := test{input: 1999, expected: "MCMXCIX"  } 
	tests = append(tests, testTwo)

	for _, test := range tests {

		result := intToRoman(test.input)
		fmt.Print("Test Result: ", result, " Expected: ", test.expected)
	}
}