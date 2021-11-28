//Package input read string from keyboard
package input

import (
	"bufio"
	"os"
)

//GetLine will read a line from keyboard and return
func GetLine() (string, error) {
	scanner := bufio.NewReader(os.Stdin)
	return scanner.ReadString('\n')
}
