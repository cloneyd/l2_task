package builder

// Product
type User struct {
	name    string
	surname string
	email   string
	age     int
	phone   string
}

// Builder
type UserBuilder interface {
	MakeName(string)
	MakeSurname(string)
	MakeEmail(string)
	MakeAge(int)
	MakePhone(string)
	Build() *User
}

// ConcreteBuilder
type SpecificUserBuilder struct {
	name    string
	surname string
	email   string
	age     int
	phone   string
}

func NewSpecificUserBuilder() *SpecificUserBuilder {
	return new(SpecificUserBuilder)
}

func (sub *SpecificUserBuilder) MakeName(name string) {
	sub.name = name
}

func (sub *SpecificUserBuilder) MakeSurname(surname string) {
	sub.surname = surname
}

func (sub *SpecificUserBuilder) MakeEmail(email string) {
	sub.email = email
}

func (sub *SpecificUserBuilder) MakeAge(age int) {
	sub.age = age
}

func (sub *SpecificUserBuilder) MakePhone(phone string) {
	sub.phone = phone
}

func (sub *SpecificUserBuilder) Build() *User {
	u := new(User)
	u.name = sub.name
	u.surname = sub.surname
	u.email = sub.email
	u.age = sub.age
	u.phone = sub.phone

	return u
}
