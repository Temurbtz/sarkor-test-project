package repositories 

import (
    "database/sql"
	_ "modernc.org/sqlite"
	"test-project/models"
)

type PhoneRepository struct {
    db *sql.DB
}

func NewPhoneRepository(db *sql.DB) *PhoneRepository {
    return &PhoneRepository{db: db}
}

// создает фон модель
func (phr *PhoneRepository) CreatePhone(phone *models.Phone) (int64, error) {
    result, err := phr.db.Exec("INSERT INTO phones (phone_number,is_fax,description,user_id) VALUES (?, ?, ?, ?)", 
	                                                phone.PhoneNumber,phone.IsFax,phone.Description,phone.UserId)
    if err != nil {
        return 0, err
    }
    return result.LastInsertId()
}

func (phr *PhoneRepository) GetPhoneById(phoneId int64) (*models.Phone, error) {
    query:="SELECT * FROM phones WHERE phone_id=?"
	row:=phr.db.QueryRow(query,phoneId)
	var phone models.Phone
	err:=row.Scan(&phone.PhoneId,&phone.PhoneNumber,&phone.IsFax,&phone.Description,&phone.UserId)
    if err != nil {
        return nil, err
    }
	return &phone, nil
}
// ищет фонб по фон_намбер
func (phr *PhoneRepository) SearchPhonesByNumber(numberPart string) ([]models.Phone, error) {
    query := "SELECT * FROM phones WHERE phone_number LIKE ?"
    rows, err := phr.db.Query(query, "%"+numberPart+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var phones []models.Phone
    for rows.Next() {
        var phone models.Phone
        if err := rows.Scan(&phone.PhoneId,&phone.PhoneNumber,&phone.IsFax,&phone.Description,&phone.UserId); err != nil {
            return nil, err
        }
        phones = append(phones, phone)
    }

    return phones, nil
}
// обновляет фон модель
func (phr *PhoneRepository) UpdatePhone(phone *models.Phone) error {
    query := "UPDATE phones SET phone_number = ?, is_fax=?,description=? WHERE phone_id = ?"
    _, err := phr.db.Exec(query, phone.PhoneNumber,phone.IsFax,phone.Description,phone.PhoneId )
    return err
}
// удаляет фон модель
func (phr *PhoneRepository) DeletePhone(phoneID int64) error {
    query := "DELETE FROM phones WHERE phone_id =?"
    _, err := phr.db.Exec(query, phoneID)
    return err
}