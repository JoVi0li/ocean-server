package shared

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hashedPasssword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasssword), []byte(password))
	return err
}
