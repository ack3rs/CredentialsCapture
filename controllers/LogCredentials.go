package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/acky666/CredentialsCapture/models/credentials"

	l "github.com/acky666/ackyLog"
)

// swagger:operation POST /save Users Save
//
// Save Credentials to the Database
//
// ---
//
// Parameters:
// - name: User
//   in: body
//   description: Save Credentials
//   schema:
//     "$ref": "#/definitions/User"
//   required: true
//
// Responses:
//   200:
//     description: OK
//   400:
//     description: Bad Request
//   500:
//     description: Internal Error (Unable to Save to Database)

func LogCredentials(w http.ResponseWriter, r *http.Request) {

	l.INFO("REQUEST: Referer %v %v %v", r.Referer(), r.RemoteAddr, r.RequestURI)

	U := credentials.User{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&U)
	if err != nil {
		l.ERROR("unable to decode request body error: %v", err)
		SendResponse(w, http.StatusBadRequest, map[string]string{"result": "fail", "resultmessage": "Malformed or Invalid JSON"})
		return
	}

	err = U.Save()
	if err != nil {
		// It could have failed Validation or DB
		SendResponse(w, http.StatusInternalServerError, map[string]string{"result": "fail", "resultmessage": err.Error()})
		return
	}

	SendResponse(w, http.StatusOK, map[string]string{"result": "ok"})
	return
}
