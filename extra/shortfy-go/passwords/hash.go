package passwords

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
