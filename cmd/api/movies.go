package main

import (
	_"encoding/json"
	"fmt"
	"net/http"
	"time"
	"damir/internal/data"
	"damir/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title 		string 		`json:"title"`
		Year 		int32 		`json:"year"`
		Runtime 	data.Runtime		`json:"runtime"`
		Genres		[]string	`json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	//err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil{
		app.badRequestResponse(w, r, err)
		//app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		//return
	}
	movie := &data.Movie{
		Title: input.Title,
		Year: input.Year,
		Runtime: input.Runtime,
		Genres: input.Genres,
	}
	v := validator.New()
	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}	
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		//http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "something",
		Runtime: 102,
		Genres: []string{"drama", "comedy"},
		Version: 1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie":movie}, nil)
	if err != nil{
		app.serverErrorResponse(w, r, err)
		//app.logger.Print(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)	
	}
}
