package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(c context.Context, userID uint) (*models.User, error) {
	return s.repo.GetUserById(c, userID)
}

func (s *UserService) RegisterUser(c context.Context, input dto.RegistrationRequest) (*models.User, error) {

	existingUser, err := s.repo.GetUserByEmail(c, input.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("Аккаунт с такой почтой уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("внутренняя ошибка сервера: %w", err)
	}

	newUser := &models.User{
		LastName:    input.LastName,
		FirstName:   input.FirstName,
		Patronymic:  input.Patronymic,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    string(hashedPassword),
	}

	err = s.repo.CreateUser(c, newUser)

	if err != nil {
		return nil, fmt.Errorf("не удалось создать новый аккаунт: %w", err)
	}

	return newUser, nil
}

func (s *UserService) LoginUser(c context.Context, input dto.LoginRequest) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(c, input.Email)
	if err != nil {
		return nil, fmt.Errorf("неверная почта или пароль")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, fmt.Errorf("неверная почта или пароль")
	}

	return user, nil
}

func (s *UserService) ChangePassword(c context.Context, userID uint, newPassword string) error {
	user, err := s.repo.GetUserById(c, userID)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	if err != nil {
		return fmt.Errorf("внутренняя ошибка сервера: %w", err)
	}
	return s.repo.Update(c, user)
}

func (s *UserService) DeleteUser(c context.Context, userID uint) error {
	return s.repo.Delete(c, userID)
}

func (s *UserService) GetFirstNameAndLastName(c context.Context, userID uint) (string, string, error) {
	user, err := s.repo.GetUserById(c, userID)
	if err != nil {
		return "", "", err
	}

	return user.FirstName, user.LastName, nil
}

func (s *UserService) ChangeCredentials(c context.Context, userID uint, newCredentials *dto.ChangeCredentialsRequest) error {
	user, err := s.repo.GetUserById(c, userID)
	if err != nil {
		return err
	}
	user.Email = newCredentials.NewEmail
	user.FirstName = newCredentials.NewFirstName
	user.LastName = newCredentials.NewLastName
	user.Patronymic = newCredentials.NewPatronymic
	user.PhoneNumber = newCredentials.NewPhoneNumber
	return s.repo.Update(c, user)
}

func (s *UserService) GetFavorites(c context.Context, userID uint) ([]models.SplitSystem, error) {
	return s.repo.GetFavorites(c, userID)
}

func (s *UserService) GetCart(c context.Context, userID uint) ([]models.SplitSystem, error) {
	return s.repo.GetCart(c, userID)
}

func (s *UserService) GetOrders(c context.Context, userID uint) ([]models.Order, error) {
	return s.repo.GetOrders(c, userID)
}

func (s *UserService) GetUserRole(c context.Context, userID uint) (string, error) {
	user, err := s.repo.GetUserById(c, userID)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func (s *UserService) IsInCart(c context.Context, userID, splitSystemID uint) bool {
	cart, err := s.repo.GetCart(c, userID)
	if err != nil {
		return false
	}

	for _, item := range cart {
		if item.ID == splitSystemID {
			return true
		}
	}
	return false
}

func (s *UserService) IsInFavorites(c context.Context, userID, splitSystemID uint) bool {
	favorites, err := s.repo.GetFavorites(c, userID)
	if err != nil {
		return false
	}

	for _, item := range favorites {
		if item.ID == splitSystemID {
			return true
		}
	}
	return false
}

func (s *UserService) HasProcessingOrders(c context.Context, userID uint) bool {
	orders, err := s.repo.GetOrders(c, userID)
	if err != nil {
		return false
	}
	hasProcessing := false
	for _, order := range orders {
		if order.Status == "В обработке" {
			hasProcessing = true
			break
		}
	}
	return hasProcessing
}
