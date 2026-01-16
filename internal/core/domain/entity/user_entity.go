package entity

type UserEntity struct {
	ID         int64
	Name       string
	Email      string
	Password   string
	RoleName   string
	Address    string
	Latitude   string
	Longitude  string
	Phone      string
	Photo      string
	IsVerified bool
	Token      string
}
