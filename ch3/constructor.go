package ch3

// Go言語のコンストラクタはこれ
func NewUser(name string, age int) *User {

	return &User{
		Name: name,
		Age:  age,
	}
}
