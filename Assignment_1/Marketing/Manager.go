package Marketing

type Manager struct {
	position string
	salary   int
	address  string
}

func (m *Manager) NewManager() Manager {
	return Manager{}
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) GetSalary() int {
	return m.salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) SetSalary(salary int) {
	m.salary = salary
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}
