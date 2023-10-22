package repositories

import (
    "database/sql"
	_ "modernc.org/sqlite"
	"test-project/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}
// lkz htubcnhfwbb
func (ur *UserRepository) CreateUser(user *models.User)  error {
    // хэшинг
	hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    query := "INSERT INTO users (login,password,name,age) VALUES (?, ?, ?, ?)"
    _, err = ur.db.Exec(query, user.Login,hashedPassword,user.Name,user.Age)
    return err
}
// юзеры могут иметь схожие имена, из-за этого метод возврашает несколко элементов
func (ur *UserRepository) SearchUsersByName(name string) ([]models.User, error) {
    query := "SELECT user_id,name,age FROM users WHERE name= ?"
    rows, err := ur.db.Query(query, name)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.UserId,&user.Name,&user.Age); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

// вытаскивает юзер с базы данных по логину
func (ur *UserRepository) GetUserByLogin(login string) (*models.User, error) {
    query := "SELECT * FROM users WHERE login = ?"
    row := ur.db.QueryRow(query, login)

    var user models.User
    err := row.Scan(&user.UserId,&user.Login,&user.Password,&user.Name,&user.Age)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
// хэширует пароль и сравнивает, дает окончательный ответ по авторизации
func (ur *UserRepository) LoginUser(login string, password string) (*models.User, error) {
    user, err := ur.GetUserByLogin(login)
    if err != nil {
        return nil, err
    }
    if err := VerifyPassword(password, user.Password); err != nil {
        return nil, err
    }

    return user, nil
}





func VerifyPassword(password, hashedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}



func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}