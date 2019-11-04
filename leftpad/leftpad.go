package leftpad

import (
	"strings"
)

func Format(s string) string {
	return 	"000" + strings.ToUpper(s);
}
