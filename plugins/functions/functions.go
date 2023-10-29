package functions

import (
	"fmt"
	"strings"
)

type Functions struct{}

func (f *Functions) CustomLower(s any) (string, error) {
	ss, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("cannot coerce %T as string", s)
	}
	return strings.ToLower(ss), nil
}

func (f *Functions) CustomLower_(cb func(m any, aliases []string, examples [][2]string)) {
	cb(f.CustomLower,
		[]string{"customLower"},
		[][2]string{
			{`{{ "Hello, world!" | lower }}`, "hello, world!"},
		},
	)
}
