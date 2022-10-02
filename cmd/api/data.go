package main

import (
	"drone_nav_serv/cmd/internal/data"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const defaultError = "Error occurred processing request"

func (app *application) uploadDataHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input struct {
		X   string `json:"x"`
		Y   string `json:"y"`
		Z   string `json:"z"`
		Vel string `json:"vel"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var c data.Coordinate
	err = c.Init(input.X, input.Y, input.Z)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	l, err := c.Getloc(input.Vel)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	res := struct {
		Loc float64 `json:"loc"`
	}{
		Loc: l,
	}
	app.writeJSON(w, 200, res, nil)
	return
}
