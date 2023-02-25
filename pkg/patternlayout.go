package pkg

type PatternLayout string

func (p PatternLayout) FormatRecord(r Record) string {
	if len(p) == 0 {
		return ""
	}

	result := string(p)
	for _, p := range Pattern {
		result = p(result, r)
	}

	return result
}
