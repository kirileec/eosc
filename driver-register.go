package eosc

import "fmt"

var (
//DefaultProfessionDriverRegister IExtenderRegister = NewExtenderRegister()
)

type ExtenderRegister struct {
	data IRegister
}

func (p *ExtenderRegister) RegisterExtender(name string, factory IExtenderDriverFactory) error {
	err := p.data.Register(name, factory, false)
	if err != nil {
		return fmt.Errorf("register profession  driver %s:%w", name, err)
	}
	return nil
}

func (p *ExtenderRegister) GetExtender(name string) (IExtenderDriverFactory, bool) {
	if v, has := p.data.Get(name); has {
		return v.(IExtenderDriverFactory), true
	}
	return nil, false
}

func NewExtenderRegister() *ExtenderRegister {
	return &ExtenderRegister{
		data: NewRegister(),
	}
}

type IExtenderRegister interface {
	RegisterExtender(name string, factory IExtenderDriverFactory) error
}

type IExtenders interface {
	GetExtender(name string) (IExtenderDriverFactory, bool)
}
