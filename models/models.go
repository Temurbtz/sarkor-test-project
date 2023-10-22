package models
//  модел для юзера
type User struct{
  UserId int64  
  Login string 
  Password string 
  Name string 
  Age int 
}
// модель для фон
type Phone struct{
	PhoneId int64
  PhoneNumber string
	IsFax bool
	Description string
  UserId int64
}