package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"net/http"
	"strconv"
)

func AddNumber(context *gin.Context, db *pg.DB) {
	value, parseErr := strconv.ParseInt(context.Param("number"), 10, 64)
	if parseErr != nil {
		context.String(http.StatusBadRequest, "%s", parseErr.Error())
		return
	}

	_, insertErr := db.Model(Number{Value: value}).Insert()
	if insertErr != nil {
		context.String(http.StatusInternalServerError, "Cannot store %s: %s", value, insertErr.Error())
		return
	}

	context.String(http.StatusOK, "%s stored.", value)
}

func ListNumbers(context *gin.Context, db *pg.DB) {
	var numbers []Number
	listErr := db.Model(numbers).Select()
	if listErr != nil {
		context.String(http.StatusInternalServerError, "Cannot list numbers: %s", listErr.Error())
		return
	}

	context.String(http.StatusOK, "Numbers list: %v", numbers)
}
