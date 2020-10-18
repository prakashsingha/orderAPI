package controllers

import "net/http"

func returnErrorResponse(w http.ResponseWriter, r *http.Request, err error, httpStatus int) {
	if err != nil {
		w.WriteHeader(httpStatus)
		w.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

}
