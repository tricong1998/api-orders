package services

type User struct {
	UserId     string
	Email      string
	IsVerified bool
}

func GetAndValidateUser(userId string) (User, error) {
	user := ReadUser(userId)
	err := ValidateUser(user)

	return user, err
}

func ValidateUser(user User) error {
	return nil
}

func ReadUser(userId string) User {
	return User{
		UserId:     "UserId",
		Email:      "mock@test.com",
		IsVerified: true,
	}
}
