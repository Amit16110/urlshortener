package utils

const Base62 uint64 = 62

var (
	// Base62 character set, [A-Za-z0-9].
	DEFAULT_CHARS = [Base62]rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	defaultLookup map[rune]int
)

func init() {
	defaultLookup = make(map[rune]int, Base62)
	for i, v := range DEFAULT_CHARS {
		defaultLookup[v] = i
	}
}

// Encode encodes the given integer into a base62 string.
func Encode(in uint64) string {
	return encode(DEFAULT_CHARS, in)
}

func encode(chars [Base62]rune, in uint64) string {
	if in < 1 {
		return ""
	}
	i, tmp := 0, in

	for tmp > 0 {
		i++

		if tmp == Base62 {
			break
		}
		tmp /= Base62
	}

	out := make([]rune, i)

	for in > 0 {
		i--

		// Overflows when modulus 62 % 62. Last character, set it and break.
		if in%Base62 == 0 {
			out[i] = chars[Base62-1]
			//If we divide in by 62, we will get 1,which will repeat another cycle. Terminate it.
			if in == Base62 {
				break
			}
		} else {
			out[i] = chars[in%Base62-1]
		}

		in /= Base62
	}
	return string(out)
}

func Decode(s string) uint64 {
	return decode(defaultLookup, s)
}

func decode(lookup map[rune]int, s string) uint64 {
	var sum uint64

	for i, v := range s {

		val := uint64(lookup[v] + 1)

		if i > 0 {
			val = val % Base62
		}
		sum = sum*Base62 + val
	}
	return sum
}
