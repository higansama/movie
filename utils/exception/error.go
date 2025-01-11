package exception

import (
	"fmt"
	"strconv"
	"strings"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

type ErrorMovie struct {
	HttpCode int
	Message  string
	Detail   string
}

// Implementasikan metode Error() untuk membuat ErrorSSO kompatibel dengan interface error
func (e ErrorMovie) Error() string {
	return fmt.Sprintf(
		`%s~%s~%s`,
		strconv.Itoa(e.HttpCode),
		e.Message,
		e.Detail,
	)
}

func (e ErrorMovie) ErrorDBDetail() string {
	return fmt.Sprintf(
		e.Detail,
	)
}

func NewErrorMovie(httpcode int, message string, err error) ErrorMovie {
	engineError := ""
	if err != nil {
		engineError = err.Error()
	}
	return ErrorMovie{
		HttpCode: httpcode,
		Message:  message,
		Detail:   engineError,
	}
}

func ErrorResponse(err error) (http int, message, detail string) {
	// cek err pars
	errSplit := strings.Split(err.Error(), "~")
	if len(errSplit) < 3 {
		return 500, "internal server error", err.Error()
	}

	fmt.Println("====================")
	fmt.Println("error ")
	fmt.Println(errSplit[0])
	fmt.Println(errSplit[1])
	fmt.Println(errSplit[2])
	fmt.Println("====================")
	i, _ := strconv.Atoi(errSplit[0])
	msg := errSplit[1]
	if string(msg[len(msg)-1]) == "~" {
		msg = msg[:len(msg)-1]
	}

	return i, msg, errSplit[2]
}
