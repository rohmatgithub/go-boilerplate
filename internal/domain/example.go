package domain

type Example struct {
	ID          string
	ExampleCode string
}

func (Example) TableName() string {
	return "users"
}
