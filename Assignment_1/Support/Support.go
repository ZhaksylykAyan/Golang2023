package Support

type Support struct {
	position string
	salary   int
	address  string
}

func (sup *Support) NewSupport() Support {
	return Support{}
}

func (sup *Support) GetPosition() string {
	return sup.position
}

func (sup *Support) GetSalary() int {
	return sup.salary
}

func (sup *Support) GetAddress() string {
	return sup.address
}

func (sup *Support) SetPosition(position string) {
	sup.position = position
}

func (sup *Support) SetSalary(salary int) {
	sup.salary = salary
}

func (sup *Support) SetAddress(address string) {
	sup.address = address
}
