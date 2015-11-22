package raindrops

import "fmt"

const TestVersion = 1

func Convert(input int) string {
	var output string;

	if (input % 3 == 0) {
		output += "Pling"
	}
	if (input % 5 == 0) {
		output += "Plang"
	}
	if (input % 7 == 0) {
		output += "Plong"
	}

	if (output == "") {
		output = fmt.Sprintf("%v",input);
	}

	return output
}
