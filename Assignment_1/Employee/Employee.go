package Employee

type Employee interface {
	GetPosition() string
	SetPosition(position string)
	GetSalary() int
	SetSalary(salary int)
	GetAddress() string
	SetAddress(address string)
}
