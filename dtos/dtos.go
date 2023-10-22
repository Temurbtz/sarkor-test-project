package dtos
// структура для создания фон, где не должен быть айди
type PhoneCreateDTO struct {
	Phone string `json:"phone_number" binding:"required"`
	IsFax bool  `json:"is_fax"`
	Description string  `json:"description" binding:"required"`
}
//  дто для возврата после создания и для обносления
type PhoneDTO struct {
	PhoneId int64 `json:"phone_id" binding:"required"`
	Phone string `json:"phone_number" binding:"required"`
	IsFax bool  `json:"is_fax"`
	Description string  `json:"description" binding:"required"`
}
// по требованию 
//ответ в JSON список тех, у кого есть этот номер
//user_id, phone, description, is_fax
type PhoneGetDTO struct {
	UserId int64 `json:"phone_id" binding:"required"`
	Phone string `json:"phone_number" binding:"required"`
	IsFax bool  `json:"is_fax"`
	Description string  `json:"description" binding:"required"`
}


// дто для авторизации
type AuthenticationDTO struct{
   Login string `json:"login" binding:"required"`
   Password string  `json:"password" binding:"required"`
}

// дто для получения юзеров по имени
type UserGetDTO struct{
	UserId int64 `json:"user_id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required"`
}
// дто для регисртации
type RegistrationForm struct{
	Login string  `form:"login" binding:"required"`
	Name string   `form:"name" binding:"required"`
	Age int64   `form:"age" binding:"required`
	Password string  `form:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required"`
}

