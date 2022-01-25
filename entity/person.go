package entity

//Person object for REST(CRUD)
type Person struct {
	ID        int64  `gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
