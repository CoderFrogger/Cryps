package things

var alpha = []string{" ", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w",
	"x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z", ".", ",", "'", "/", "!", "?", ":",
	";", "(", ")", "-"}

var beta = []string{"000", "001", "002", "003", "004", "005", "006", "007", "008", "009", "010", "011",
	"012", "013", "014", "015", "016", "017", "018", "019", "020", "021", "022", "023",
	"024", "025", "026", "027", "028", "029", "030", "031", "032", "033", "034", "035",
	"036", "037", "038", "039", "040", "041", "042", "043", "044", "045", "046", "047",
	"048", "049", "050", "051", "052", "053", "054", "055", "056", "057", "058", "059",
	"060", "061", "062", "063", "064", "065", "066"}

func Encrypt(content []byte) []string {
	s := string(content)
	var newContent []string

	for _, c := range s {
		for i := 0; i < len(alpha); i++ {
			if string(c) == alpha[i] {
				newContent = append(newContent, beta[i])
			}
		}
	}

	return newContent
}

func Decrypt(content []byte) []string {
	s := string(content)
	partition := ""
	var newContent []string
	index := 2

	for j, c := range s {
		partition += string(c)
		if j == index {
			for i := 0; i < len(beta); i++ {
				if partition == beta[i] {
					newContent = append(newContent, alpha[i])
				}
			}
			partition = ""
			index += 3
		}
	}

	return newContent
}
