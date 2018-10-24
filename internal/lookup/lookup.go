package lookup

// Lookup represents a string-to-string lookup.
type Lookup interface {
	Lookup(key string) (value *string, hasLongerMatch bool)
}

// FromMap builds a Lookup from string-to-string map.
func FromMap(m map[string]string) Lookup {
	t := make(tree, len(m))
	for k, v := range m {
		t.Insert(k, v)
	}
	return t
}

// ----------------------------------------------------------------------------

type tree map[rune]*leaf

type leaf struct {
	value    *string
	children tree
}

func (t tree) Lookup(key string) (*string, bool) {
	var l *leaf
	for _, r := range key {
		l = t[r]
		if l == nil {
			return nil, false // break
		}
		t = l.children
	}
	if l == nil {
		return nil, false
	}
	return l.value, len(l.children) > 0
}

func (t tree) Insert(key, value string) {
	var l *leaf
	for _, r := range key {
		l = t[r]
		if l == nil {
			l = &leaf{
				children: tree{},
			}
			t[r] = l
		}
		t = l.children
	}
	l.value = &value
}
