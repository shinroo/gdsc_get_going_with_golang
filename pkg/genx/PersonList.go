package genx

// PersonList is a list of pointers to people
type PersonList []*Person

// PersonToBool is a stub function signature that takes a pointer to a person and returns a boolean
type PersonToBool func(*Person) bool

// Filter allows us to pass in a function to PersonToBool function that filters a PersonList
func (al PersonList) Filter(f PersonToBool) PersonList {
	var ret PersonList
	for _, a := range al {
		if f(a) {
			ret = append(ret, a)
		}
	}
	return ret
}
