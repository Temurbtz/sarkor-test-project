package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"test-project/dtos"
	"test-project/services"
	"github.com/gin-gonic/gin"
	"test-project/utils"
)
// контроллер для фон
type PhoneController struct {
	phoneService *services.PhoneService
}

func NewPhoneController(phoneService *services.PhoneService) *PhoneController {
	return &PhoneController{phoneService: phoneService}
}

// создает новый фон
func (phc *PhoneController) CreatePhone(c *gin.Context) {
	// принимает дто 
	var phoneCreateDto dtos.PhoneCreateDTO
	// проверяет биндинг
	if err := c.ShouldBindJSON(&phoneCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// через мидлвар получает клаим
	claims, exists:=c.Get("claims")
	if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error 1"})
        return
    }
	// парсит в нужный тип
	authClaims, ok := claims.(*utils.AuthClaims)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert claims to AuthClaims"})
        return
    }
	// вытаскивает айди юзера, передает в сервис вместе дто
	phoneGetDTO, err := phc.phoneService.CreatePhone(&phoneCreateDto,authClaims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,phoneGetDTO)
}
// ищет фон с помошью фон_намбер
func (phc *PhoneController) GetPhone(c *gin.Context) {
	// проводит валидацию
	phoneNumber:=c.Param("phone_number")
	phoneRegex := `^\+?\d{1,}$|^\d{3}-\d{3}-\d{4}$`
	_,err := regexp.MatchString(phoneRegex, phoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"it is not phone number"})
	}
	// передает параметр в сервис и получает список фон которые сходятся с даннам паттерном
	phoneGetDTOs, err := phc.phoneService.SearchPhone(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK,phoneGetDTOs)
}
// обновляет фон
func (phc *PhoneController) UpdatePhone(c *gin.Context) {
	var phonePutDto dtos.PhoneDTO
	if err := c.ShouldBindJSON(&phonePutDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	phoneDTO, err := phc.phoneService.PutPhone(&phonePutDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK,phoneDTO)
}
// удаляет фон
func (phc *PhoneController) DeletePhone(c *gin.Context) {
	phoneID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone ID"})
		return
	}
	err = phc.phoneService.DeletePhone(phoneID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Phone not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phone deleted successfully"})
}
