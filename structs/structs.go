package structs

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
}

type Ingrident struct {
	Item []string
}

type Recipe struct {
	Menu
	Ingrident
}
