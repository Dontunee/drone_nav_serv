package data

import (
	"math"
	"strconv"
)

const invalidInput = "Invalid input"
const invalidVec = "Invalid vec"

const sectorId = 1

type Coordinate struct {
	X, Y, Z float64
}

func (c *Coordinate) Init(x, y, z string) string {
	var err error
	var sx, sy, sz float64
	if sx, err = strconv.ParseFloat(x, 64); err != nil {
		return invalidInput
	}

	if sy, err = strconv.ParseFloat(y, 64); err != nil {
		return invalidInput
	}

	if sz, err = strconv.ParseFloat(z, 64); err != nil {
		return invalidInput
	}
	c.X = sx
	c.Y = sy
	c.Z = sz
	return ""
}
func (c *Coordinate) Getloc(v string) (float64, string) {
	if vel, err := strconv.ParseFloat(v, 64); err == nil {
		res := (c.X * sectorId) + (c.Y * sectorId) + (c.Z * sectorId) + vel
		res = toFixed(res, 2)
		return res, ""
	}
	return 0, invalidVec
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
