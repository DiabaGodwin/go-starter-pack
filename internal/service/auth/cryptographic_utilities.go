package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func CheckPassword(password, hashedPassword string) (bool, error) {
	res := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if res != nil {
		return false, res
	}
	return true, nil
}
