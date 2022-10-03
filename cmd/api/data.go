package main

import (
	"encoding/json"
	"github.com/dontunee/drone_nav_serv/cmd/internal/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const defaultError = "Error occurred processing request"

type response struct {
	Loc float64 `json:"loc"`
}

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
	error := c.Init(input.X, input.Y, input.Z)
	if error != "" {
		app.errorResponse(w, r, http.StatusBadRequest, error)
		return
	}

	l, error := c.Getloc(input.Vel)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, error)
		return
	}

	res := response{
		Loc: l,
	}
	app.writeJSON(w, 200, res, nil)
	return
}
