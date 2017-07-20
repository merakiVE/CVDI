package tags

import (
	"fmt"
)

var (
	tagsDefaultRules = map[string]HandleTag{
		"default": RuleDefault,
	}
)

func RuleDefault(_model interface{}) () {
	fmt.Println("Handle default")
}
