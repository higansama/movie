package reqres

import (
	"movie-app/utils/exception"

	"github.com/gin-gonic/gin"
)

type DefaultResponse struct {
	Code      int         `json:"code"`
	Detail    string      `json:"message"`
	ErrDetail string      `json:"detail_err,omitempty"`
	Status    bool        `json:"status"`
	Data      interface{} `json:"data"`
}

func JsonResponse(ctx *gin.Context, err error, data interface{}) {
	// fmt.Println("response")
	if err != nil {
		coder, msg, detail := exception.ErrorResponse(err)
		ctx.JSON(coder, DefaultResponse{
			Status:    false,
			Detail:    msg,
			ErrDetail: detail,
			Code:      coder,
			Data:      nil,
		})
		return // Ensure we return to avoid further execution
	}
	// Success response
	ctx.JSON(200, DefaultResponse{
		Status: true,
		Detail: "success",
		Code:   200,
		Data:   data,
	})
}
