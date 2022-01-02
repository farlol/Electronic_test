package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type SmartPhone interface {
	OS() string
}

type applePhone struct {
	brand string
	model string
	type1 string
	os    string
}

type androidPhone struct {
	brand string
	model string
	type1 string
	os    string
}

type radioPhone struct {
	brand        string
	model        string
	type1        string
	buttonscount int
}

func NewApple(model, os string) applePhone {
	new := applePhone{}
	new.brand = "apple"
	new.model = model
	new.type1 = "smartphone"
	new.os = os
	return new
}

func NewAndroid(brand, model, os string) androidPhone {
	new := androidPhone{}
	new.brand = brand
	new.model = model
	new.type1 = "smartphone"
	new.os = os
	return new
}

func NewRadio(brand, model string, buttonscount int) radioPhone {
	new := radioPhone{}
	new.brand = brand
	new.model = model
	new.type1 = "station"
	new.buttonscount = buttonscount
	return new
}

func (a applePhone) Brand() string {
	return a.brand
}
func (a androidPhone) Brand() string {
	return a.brand
}
func (a radioPhone) Brand() string {
	return a.brand
}
func (a applePhone) Model() string {
	return a.model
}
func (a androidPhone) Model() string {
	return a.model
}
func (a radioPhone) Model() string {
	return a.model
}
func (a applePhone) Type() string {
	return a.type1
}
func (a androidPhone) Type() string {
	return a.type1
}
func (a radioPhone) Type() string {
	return a.type1
}
func (a applePhone) OS() string {
	return a.os
}
func (a androidPhone) OS() string {
	return a.os
}
func (a radioPhone) ButtonsCount() int {
	return a.buttonscount
}
