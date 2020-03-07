package mocks

type Context struct {
}

func (c Context) String(code int, message string) error {
	return nil
}

func (c Context) HTML(code int, message string) error {
	return nil
}

func (c Context) JSON(code int, data interface{}) error {
	return nil
}

func (c Context) Redirect(code int, url string) error {
	return nil
}

func (c Context) File(file string) error {
	return nil
}
