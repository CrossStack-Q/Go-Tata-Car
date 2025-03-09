package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"float64"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"float64"`
}

func ValidateRequest(carReq CarRequest) error {
	if err := validateName(carReq.Name); err != nil {
		return err
	}
	if err := validateYear(carReq.Year); err != nil {
		return err
	}
	if err := validateBrand(carReq.Brand); err != nil {
		return err
	}
	if err := validateFuelType(carReq.FuelType); err != nil {
		return err
	}
	if err := validateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := validatePrice(carReq.Price); err != nil {
		return err
	}
	return nil

}

func validateName(name string) error {
	if name == "" {
		return errors.New("Name is Required")
	}

	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("Year is Required")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("Year is invalid")
	}
	currentYear := time.Now().Year()
	if yearInt > currentYear {
		return errors.New("Year is greater than current year")
	}
	if yearInt < 1886 {
		return errors.New("Year is less than 2005")
	}

	return nil

}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("Brand is Required")
	}

	return nil
}

func validateFuelType(fuelType string) error {
	validateFuelTypes := []string{"petrol", "diesal"}

	for _, value := range validateFuelTypes {
		if value == fuelType {
			return nil
		}
	}

	return errors.New("FuelType Must be one of them .")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("Engine ID is Required")
	}
	if engine.Displacement <= 0 {
		return errors.New("Displacement is greater than 0")
	}
	if engine.NoOfCyclinder <= 0 {
		return errors.New("No. of Cyclinders should be valid")
	}
	if engine.CarRange <= 0 {
		return errors.New("carRange must be greater than 0")
	}

	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("Price must be greater than Zero")
	}

	return nil
}
