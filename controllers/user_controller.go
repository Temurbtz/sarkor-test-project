package controllers

import (
	"net/http"
	"test-project/dtos"
	"test-project/services"

	"github.com/gin-gonic/gin"
)
type UserController struct{
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}
// функция для регистрации
func (uc *UserController) Registration(c *gin.Context){
	var userFormData dtos.RegistrationForm
	if err := c.ShouldBind(&userFormData); err != nil {
	   c.HTML(http.StatusBadRequest, "registration.html", gin.H{"error": err.Error()})
	   return
	}
	// валидация инпутов
	var errorList []string
	if 4>len(userFormData.Login) || len(userFormData.Login)>30{
		errorList=append(errorList,"Wrong lenght of login.Login should be between 4 and 30 characters")
	}
	exists:=uc.userService.LoginExists(userFormData.Login)
	if exists{
		errorList=append(errorList,"Login already exists")
	}
	if 4>len(userFormData.Password)|| len(userFormData.Password)>30{
		errorList=append(errorList,"Wrong length of password.Password should be between 4 and 30 letters")
	}
	if userFormData.Password!=userFormData.PasswordConfirmation{
		errorList=append(errorList,"Password and its confirmation should match")
	}
	if(len(errorList)!=0){
		c.HTML(http.StatusBadRequest, "registration.html",gin.H{"Errors":errorList,"FormData": userFormData})
		return
	}
	// передача инпутов в нужный сервис
	uc.userService.CreateUser(&userFormData)
	c.HTML(http.StatusOK, "success.html",nil)
}
// авторизация
func (uc *UserController) Login(c *gin.Context){
	var authenticationDTO dtos.AuthenticationDTO
	if err := c.ShouldBindJSON(&authenticationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    // получение токена
	tokenString, err:=uc.userService.Authenticate(&authenticationDTO)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// вставление токен в куки
	c.SetCookie("SESSTOKEN", tokenString, 3600, "/", "localhost", false, true)
	c.String(200, "Hello from user login")
}
// получение юзера по имени
func (uc *UserController) GetByName(c *gin.Context){
	userName:=c.Param("user_name")
	userGetDTOs, err := uc.userService.GetUserByName(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,userGetDTOs)
}
// для показа формы для регистрации
func(uc *UserController) ShowRegistrationForm(c *gin.Context){
	c.HTML(http.StatusOK,"registration.html",nil)
}

