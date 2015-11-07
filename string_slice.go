package gotility

type StringSlice []string

func (t StringSlice) Contains(s string) bool {
	for i := range t {
		if t[i] == s {
			return true
		}
	}

	return false
}

func (t *StringSlice) Add(s ...string) {
	*t = append(*t, s...)
}

func (t *StringSlice) Delete(s string) {
	for i := range *t {
		if (*t)[i] == s {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return
		}
	}
}

func (t *StringSlice) Reverse() {
	var tmp string
	for i := 0; i < len(*t)/2; i++ {
		j := len(*t) - i - 1
		tmp = (*t)[i]

		// swap the element
		(*t)[i] = (*t)[j]
		(*t)[j] = tmp
	}
}
