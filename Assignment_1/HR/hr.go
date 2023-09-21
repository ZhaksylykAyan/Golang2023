package HR

type HR struct {
	position string
	salary   int
	address  string
}

func (hr *HR) NewHR() HR {
	return HR{}
}

func (hr *HR) GetPosition() string {
	return hr.position
}

func (hr *HR) GetSalary() int {
	return hr.salary
}

func (hr *HR) GetAddress() string {
	return hr.address
}

func (hr *HR) SetPosition(position string) {
	hr.position = position
}

func (hr *HR) SetSalary(salary int) {
	hr.salary = salary
}

func (hr *HR) SetAddress(address string) {
	hr.address = address
}
