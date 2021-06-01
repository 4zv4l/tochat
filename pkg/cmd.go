package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"unicode"
)

// Clear the screen
// works on Windows and Unix/based system
func Clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// rot13(alphabets) + rot5(numeric)
func encrypt(input string) string {

	var result []rune
	rot5map := map[rune]rune{'0': '5', '1': '6', '2': '7', '3': '8', '4': '9', '5': '0', '6': '1', '7': '2', '8': '3', '9': '4'}

	for _, i := range input {
		switch {
		case !unicode.IsLetter(i) && !unicode.IsNumber(i):
			result = append(result, i) // doesn't change special caracters
		case i >= 'A' && i <= 'Z':
			result = append(result, 'A'+(i-'A'+13)%26) // rot 13
		case i >= 'a' && i <= 'z':
			result = append(result, 'a'+(i-'a'+13)%26) // rot13
		case i >= '0' && i <= '9':
			result = append(result, rot5map[i]) // rot 5
		case unicode.IsSpace(i):
			result = append(result, ' ')
		}
	}
	return fmt.Sprintf(string(result[:]))
}
