package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusCode int

const (
	Success StatusCode = iota
	NotFound
	Full
	Updated
	Deleted
	Error
)

func (s StatusCode) Get() (string, int) {
	switch s {
	case Success:
		return "Success", http.StatusOK
	case NotFound:
		return "Not Found In Database", http.StatusBadRequest
	case Full:
		return "Full", http.StatusBadRequest
	case Updated:
		return "Updated", http.StatusOK
	case Deleted:
		return "Deleted", http.StatusOK
	case Error:
		return "Error", http.StatusBadRequest
	default:
		return "UNKNOWN", -1
	}
}

func (cc *CustomerController) statusSend(s StatusCode, ctx *gin.Context, options ...string) {
	data, status := s.Get()
	ctx.JSON(status, gin.H{
		"message": data,
	})

	record := ""
	for _, option := range options {
		record += option
	}

	if s != Error {
		cc.SugerLogger.Infoln(data, status, record)

	} else {

		cc.SugerLogger.Errorln(data, status, record)

	}

}
