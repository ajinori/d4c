package mocks

type Context struct {
}

func (c Context) String(code int, message string) error {
	return nil
}

func (c Context) HTML(code int, message string) error {
	return nil
}
