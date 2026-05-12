package services

type Notify struct {
}

func NewNotify() Notify {
	return Notify{}
}
func (n Notify) Send() error {
	return nil
}
