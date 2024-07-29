package simplefactory

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("hi, %v", name)
}

type helloAPI struct{}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %v", name)
}

func NewAPI(n int) API {
	switch n {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	}
	return nil
}
