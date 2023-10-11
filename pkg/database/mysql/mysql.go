package mysql

type Person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func (p *Person) Get() {}

func (p *Person) Create() {}