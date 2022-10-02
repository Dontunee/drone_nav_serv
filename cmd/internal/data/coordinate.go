package data

import (
	"math"
	"strconv"
)

const sectorId = 1

type Coordinate struct {
	X, Y, Z float64
}

func (c *Coordinate) Init(x, y, z string) error {
	var err error
	var sx, sy, sz float64
	if sx, err = strconv.ParseFloat(x, 64); err != nil {
		return err
	}

	if sy, err = strconv.ParseFloat(y, 64); err != nil {
		return err
	}

	if sz, err = strconv.ParseFloat(z, 64); err != nil {
		return err
	}
	c.X = sx
	c.Y = sy
	c.Z = sz
	return nil

}
func (c *Coordinate) Getloc(v string) (float64, error) {
	var err error
	if vel, err := strconv.ParseFloat(v, 64); err == nil {
		res := (c.X * sectorId) + (c.Y * sectorId) + (c.Z * sectorId) + vel
		res = toFixed(res, 2)
		return res, nil
	}
	return 0, err
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
