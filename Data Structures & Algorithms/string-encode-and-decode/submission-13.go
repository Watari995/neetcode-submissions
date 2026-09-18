type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// 5#hogeh のように{文字数}#をつけて繋げる
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(fmt.Sprintf("%d#%s", len(s), s))
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	result := make([]string, 0)
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		// 数字の部分が入ってくる想定
		length, _ := strconv.Atoi(encoded[i:j])
		// その文字列をappend
		result = append(result, encoded[j+1:j+1+length])
		// iを進める
		i = j + 1 + length
	}
	return result
}
