package lemonsqueezy

import (
	"fmt"
	"strings"
)

// pathWithFilters adds filters to the path.
//
// https://docs.lemonsqueezy.com/api/products#list-all-products
func pathWithFilters(path string, filterKV ...string) string {
	var prefix = "?"
	if strings.Contains(path, "?") {
		prefix = "&"
		if strings.HasSuffix(path, "?") || strings.HasSuffix(path, "&") {
			prefix = ""
		}
	}
	for i := 0; i < len(filterKV); i += 2 {
		path += fmt.Sprintf("%sfilter[%s]=%s", prefix, filterKV[i], filterKV[i+1])
		prefix = "&"
	}
	return path
}
