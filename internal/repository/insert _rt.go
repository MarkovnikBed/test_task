package repository

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func (rep *Repository) InsertRT(refreshToken string, id string) error {
	refreshSHA256 := sha256.Sum256([]byte(refreshToken))
	hashedJWT, err := bcrypt.GenerateFromPassword(refreshSHA256[:], bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = rep.DB.Exec("INSERT INTO refresh_tokens (id, hash) VALUES ($1,$2) ON CONFLICT (id) DO UPDATE SET hash = EXCLUDED.hash", id, string(hashedJWT))
	if err != nil {
		return err
	}
	return nil
}
