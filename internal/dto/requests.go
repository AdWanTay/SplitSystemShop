package dto

type RegistrationRequest struct {
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangeCredentialsRequest struct {
	NewPhoneNumber string `json:"new_phone_number"`
	NewLastName    string `json:"new_last_name"`
	NewFirstName   string `json:"new_first_name"`
	NewPatronymic  string `json:"new_patronymic"`
	NewEmail       string `json:"new_email"`
}
