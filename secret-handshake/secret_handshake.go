package secret

import "fmt"

func Handshake(n int) []string {
	var binary string = fmt.Sprintf("%05b", n)
	var reverse bool = false
	var out []string

	if len(binary) > 5 || n <= 0 {
		return nil
	}

	if string(binary[0]) == "1" {
		reverse = true
	}

	if string(binary[4]) == "1" {
		out = append(out, "wink")
	}

	if string(binary[3]) == "1" {
		if reverse {
			out = append([]string{"double blink"}, out...)
		} else {
			out = append(out, "double blink")
		}
	}

	if string(binary[2]) == "1" {
		if reverse {
			out = append([]string{"close your eyes"}, out...)
		} else {
			out = append(out, "close your eyes")
		}
	}

	if string(binary[1]) == "1" {
		if reverse {
			out = append([]string{"jump"}, out...)
		} else {
			out = append(out, "jump")
		}
	}

	return out
}
