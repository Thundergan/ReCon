package helpers

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

// Gets a gin Context, an http status as int and a arbitrary payload
// and writes it back to the requester via the gin context.
func RespondJSON(w *gin.Context, status int, payload interface{}) {
	var res ResponseData

	res.Status = status
	res.Data = payload
	w.JSON(200, res)
}
