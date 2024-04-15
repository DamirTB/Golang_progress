package main

import (
	_"encoding/json"
	_ "fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// data := map[string]string{
	// 	"status":      "available",
	// 	"environment": app.config.env,
	// 	"version":     version,
	// }
	env := envelope{
		"status" : "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		//app.logger.Print(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// 	js := `{"status":"available", "environment":%q, "version":%q}`
// 	js = fmt.Sprintf(js, app.config.env, version)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(js))
// }

// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "status: available")
// 	fmt.Fprintf(w, "enviroment: %s\n", app.config.env)
// 	fmt.Fprintf(w, "version: %s\n", version)
// }
