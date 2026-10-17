package entities

type User struct {
	Id      int64  `cli:"ID,ro"`
	Name    string `cli:"Name"`
	Email   string
	Age     int
	Active  bool
	Secret  string `cli:"-"`
	Address struct {
		City string
		Zip  string
	}
}
