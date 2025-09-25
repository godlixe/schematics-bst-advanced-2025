package model

import "time"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`

	// Password tidak disimpan di database, hanya hash-nya.
	// Hal tersebut dilakukan agar data password tidak bocor apabila terjadi kebocoran database.
	// verifikasi password akan dilakukan dengan menggunakan PasswordHash.
	// tag gorm `-` digunakan untuk mengecualikan pembuatan kolom dalam database.
	Password     string `json:"-" gorm:"-"`
	PasswordHash string `json:"-"`

	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
