package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && prefix != strs[i][:len(prefix)] {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

func main() {
	prefixes := []string{"flow", "flower", "flight"}

	println(longestCommonPrefix(prefixes))
}
