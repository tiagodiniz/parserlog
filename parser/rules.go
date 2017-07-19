package parser

// used regexp2 instead regexp for better performance
import (
	"github.com/dlclark/regexp2"
)

// rules
var kill, killer, killed *regexp2.Regexp

func init() {
	kill = regexp2.MustCompile(`:\s([^:]+)\skilled\s(.*?)\sby\s[a-zA-Z_]+`, 0)
	killer = regexp2.MustCompile(`(?<=\d:\s)(.*?)(?=\skilled)`, 0)
	killed = regexp2.MustCompile(`(?<=killed\s)(.*?)(?=\sby)`, 0)
}
