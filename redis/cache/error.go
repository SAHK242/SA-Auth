package cache

type Error struct {
	Op
	Err error
}

func NewGetError(err error) Error {
	return Error{
		Op:  OpGet,
		Err: err,
	}
}

func NewSetError(err error) Error {
	return Error{
		Op:  OpSet,
		Err: err,
	}
}
