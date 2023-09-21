package Engineer

type Engineer struct {
	position string
	salary   int
	address  string
}

func NewEngineer() Engineer {
	return Engineer{}
}

func (e *Engineer) GetPosition() string {
	return e.position
}

func (e *Engineer) GetSalary() int {
	return e.salary
}

func (e *Engineer) GetAddress() string {
	return e.address
}

func (e *Engineer) SetPosition(position string) {
	e.position = position
}

func (e *Engineer) SetSalary(salary int) {
	e.salary = salary
}

func (e *Engineer) SetAddress(address string) {
	e.address = address
}
