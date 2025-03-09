package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCyclinder int64     `json:"no_of_cyclinder"`
	CarRange      int64     `json:"car_range"`
}

type EngineReq struct {
	Displacement int64 `json:"displacement"`
	NoOfCylinder int64 `json:"no_of_cylinder"`
	CarRange     int64 `json:"car_range"`
}

func validateEngineReq(engine EngineReq) error {
	if err := validateDisplacement(engine.Displacement); err != nil {
		return err
	}
	if err := validateCylinders(engine.NoOfCylinder); err != nil {
		return err
	}
	if err := validateCarRange(engine.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("Displacement must be greater than ZERO")
	}
	return nil
}

func validateCylinders(noOfCylinder int64) error {
	if noOfCylinder <= 0 {
		return errors.New("noOfCylinder must be greater than ZERO")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than ZERO")
	}
	return nil
}
