package validators

import (
	"crud-api-barang/models"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateBarangCreate(input *models.Barang) error {
	validate := validator.New()

	// Menambahkan validasi untuk Nama (required dan min 3 karakter)
	validate.RegisterValidation("requiredmin3", func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			return len(strings.TrimSpace(value)) == 0 || len(value) >= 3
		}
		return false
	})
	validate.RegisterAlias("nama", "requiredmin3")

	// Menambahkan validasi untuk Harga (numerical dan min 3)
	validate.RegisterValidation("numericalmin3", func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(int); ok {
			return value >= 3
		}
		return false
	})
	validate.RegisterAlias("harga", "numericalmin3")

	// Menambahkan validasi untuk Deskripsi (required dan min 3 karakter)
	validate.RegisterValidation("requiredmin3", func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			return len(strings.TrimSpace(value)) == 0 || len(value) >= 3
		}
		return false
	})
	validate.RegisterAlias("deskripsi", "requiredmin3")

	err := validate.Struct(input)
	if err != nil {
		return err
	}

	return nil
}
