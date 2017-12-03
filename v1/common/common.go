package common

type Error struct {
	Err string
}

func (e Error) Error() string {
	return e.Err
}
