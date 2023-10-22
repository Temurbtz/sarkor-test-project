package services


import (
	"test-project/repositories"
	"test-project/models"
	"test-project/dtos"
)
// сервис для фон модели, где запишутся весь бизнес код связанный с фон
type PhoneService struct {
    phoneRepository *repositories.PhoneRepository
}

func NewPhoneService(phoneRepository *repositories.PhoneRepository) *PhoneService {
    return &PhoneService{phoneRepository: phoneRepository}
}
// контроллеры не должны работать прямо с моделями, из-за этого здесь происходят маппинги между моделями и дто
func (phs *PhoneService) CreatePhone(phoneCreateDto *dtos.PhoneCreateDTO, userId int64) (*dtos.PhoneDTO, error) {
	// преврашает входящий дто в модель
	phone := models.Phone{
		UserId: userId,
        PhoneNumber:  phoneCreateDto.Phone,
        IsFax: phoneCreateDto.IsFax,
		Description: phoneCreateDto.Description,
    }
	// передает модель в репозитории и  получает ответ
	phoneId, err:=phs.phoneRepository.CreatePhone(&phone)
	if err != nil {
        return nil, err
    }
	// превращает модель в нужный вид дто
	phoneCreateGetDto:= dtos.PhoneDTO{
		PhoneId: phoneId,
		Phone: phoneCreateDto.Phone,
		IsFax: phoneCreateDto.IsFax,
		Description: phoneCreateDto.Description,
	 }
    return  &phoneCreateGetDto,nil
}
//  ищет номера, например пришел параметр 991б возвращает все номера которые содержат 9916 
func (phs *PhoneService) SearchPhone(phoneNumber string) ([]dtos.PhoneGetDTO, error){
	phones,err :=phs.phoneRepository.SearchPhonesByNumber(phoneNumber)
	if err != nil {
        return nil, err
    }
	phoneDTOs := make([]dtos.PhoneGetDTO, len(phones))
    for i, phone := range phones {
        phoneDTOs[i] = dtos.PhoneGetDTO{
			UserId: phone.UserId,
			Phone:phone.PhoneNumber,
			IsFax: phone.IsFax,
			Description: phone.Description,
        }
    }
	return phoneDTOs,nil 
}

// обновляет фон модель

func (phs *PhoneService) PutPhone(phonePutDto *dtos.PhoneDTO) (*dtos.PhoneDTO, error) {
	phone := models.Phone{
		PhoneId: phonePutDto.PhoneId,
        PhoneNumber:  phonePutDto.Phone,
        IsFax: phonePutDto.IsFax,
		Description: phonePutDto.Description,
    }
	err:=phs.phoneRepository.UpdatePhone(&phone)
	if err != nil {
        return nil, err
    }
	phoneGetDto:=dtos.PhoneDTO{
		PhoneId: phonePutDto.PhoneId,
        Phone:phonePutDto.Phone,
		IsFax: phonePutDto.IsFax,
		Description: phonePutDto.Description,
	}
    return  &phoneGetDto,nil
}

// удаляет фон модель
func (phs *PhoneService) DeletePhone(phoneId int64) error {
    err := phs.phoneRepository.DeletePhone(phoneId)
    if err != nil {
        return err
    }
    return nil
}

