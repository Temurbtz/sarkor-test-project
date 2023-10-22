package services

import (
	"test-project/repositories"
	"test-project/dtos"
    "test-project/models"
    "test-project/utils"
    "github.com/dgrijalva/jwt-go"
    "time"
)
// сервис для юзерь модели, где запишутся весь бизнес код связанный юзерь моделью
type UserService struct {
    userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
    return &UserService{userRepository: userRepository}
}
// получение юзера по имени
func (us *UserService) GetUserByName(name string)([]dtos.UserGetDTO, error){
	users,err :=us.userRepository.SearchUsersByName(name)
	if err != nil {
        return nil, err
    }
	userDTOs := make([]dtos.UserGetDTO, len(users))
    for i, user := range users {
        userDTOs[i] = dtos.UserGetDTO{
			UserId :user.UserId,
			Name: user.Name,
            Age: user.Age,
        }
    }
	return userDTOs,nil 
}
// этот метод нужен чтобы не повторялись логины при регистрации
func (us *UserService) LoginExists(login string) bool{
    user,_:=us.userRepository.GetUserByLogin(login)
    if user!=nil{
        return true;
    } else {
        return false
    }
}
// для регистрации
func (us *UserService) CreateUser(registrationData *dtos.RegistrationForm) error{
    user:= models.User{
       Login:registrationData.Login,
       Password:registrationData.Password,
       Name:registrationData.Name,
       Age:int(registrationData.Age),
    }
    return us.userRepository.CreateUser(&user)
}
// для авторизации
func (us *UserService) Authenticate(authenticationDTO *dtos.AuthenticationDTO)(string,error){
    // проверка с базы данных
    user,err:=us.userRepository.LoginUser(authenticationDTO.Login,authenticationDTO.Password)
    if err != nil {
        return  "", err
    }
    // создание клаим
    claims := &utils.AuthClaims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
        },
        UserID: user.UserId,
        Login:  user.Login,
    }
    // создание токена
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("jwt_secret_key"))
	if err != nil {
		panic("Error generating JWT token")
	}
    return tokenString, nil
}




