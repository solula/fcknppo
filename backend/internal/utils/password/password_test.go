package password

import "testing"

func TestPasswordHashing(t *testing.T) {
	password := "admin"
	hash, err := GenerateHash(password)
	if err != nil {
		t.Errorf("Failed to hash password: %v", err)
	}
	t.Log(hash)

	match := CheckHash(password, hash)
	if !match {
		t.Errorf("Password hash does not match original password")
	}
}
