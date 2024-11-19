package models

type Universe map[string]bool

func NewUniverse(s []string) Universe {
	u := make(Universe)
	for _, i := range s {
		u[i] = true
	}
	return u
}

func (u Universe) ContainSet(s []string) bool {
	for _, i := range s {
		if !u[i] {
			return false
		}
	}
	return true
}
