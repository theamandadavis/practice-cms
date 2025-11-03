package id

type UUID string

func (u UUID) String() string {
	return string(u)
}

func UUIDFromString(s *string) *UUID {
	if s == nil {
		return nil
	}
	return (*UUID)(s)
}
func (u UUID) IsEmpty() bool {
	return u == ""
}
