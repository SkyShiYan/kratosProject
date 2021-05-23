package biz

type NoDataErr struct {
	Err error
}

func (err NoDataErr) Error() string {
	return err.Err.Error()
}
