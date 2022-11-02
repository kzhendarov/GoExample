package class

import "fmt"

/* PersonData: Формат данных для обработки данных типа Person (дублирует поля структуры Person, но поля возможны для доступа из стороннего пакета
[используется при необходимости получения значения всех полей структуры и при работе ])
*/
type Person_data struct {
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Age       int    `json:"age"`
	Contacts_data
}

type Contacts_data struct {
	Email string `json:"e-mail"`
	Phone string `json:"phone"`
}

//: Person структура данных профиля (Все поля приватные, могут быть выведены или изменены соответствующим методами)
type Person struct {
	firstName string
	lastName  string
	age       int
	contacts
}

// Структура контактов профиля
type contacts struct {
	email string
	phone string
}

// NewPerson: Псевдо-конструктор типа данных Person
func NewPerson(first_name, last_name string, age int, email, phone string) *Person {
	return &Person{first_name, last_name, age, contacts{email, phone}}
}

func (p Person) IString() string {
	return fmt.Sprintf(`Карточка профиля (%s)
	Имя:	%s
	Фамилия:%s
	возраст: %d
	Контакты:
		E-mail: %s
		Phone:	%s
`, p.fullName(), p.firstName, p.lastName, p.age, p.email, p.phone)
}

func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// GetData - can return slice of some ghjperties of obj Person type
func (p Person) GetData() Person_data {
	return Person_data{p.firstName, p.lastName, p.age, Contacts_data{p.email, p.phone}}
}
