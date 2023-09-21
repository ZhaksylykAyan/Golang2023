package Sales

type Sales struct {
	position string
	salary   int
	address  string
}

func (s *Sales) NewSales() Sales {
	return Sales{}
}

func (s *Sales) GetPosition() string {
	return s.position
}

func (s *Sales) GetSalary() int {
	return s.salary
}

func (s *Sales) GetAddress() string {
	return s.address
}

func (s *Sales) SetPosition(position string) {
	s.position = position
}

func (s *Sales) SetSalary(salary int) {
	s.salary = salary
}

func (s *Sales) SetAddress(address string) {
	s.address = address
}
