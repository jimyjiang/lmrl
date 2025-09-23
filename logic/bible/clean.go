package bible

import "strings"

// 激励思考的谈话（提摩太后书2:22-3:9）
func Clean(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	end := len(s)
	if strings.Contains(s, "（") {
		start = strings.Index(s, "（") + len("（")

	}
	if strings.Contains(s, "）") {
		end = strings.LastIndex(s, "）")
	}
	if end < start {
		return s
	}
	return s[start:end]
}
