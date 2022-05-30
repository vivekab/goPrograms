// number as query param and fact(n) as response
// /factorial?number= GET
// SUCCESS_CASE { "data" : <value> }
// BAD_REQUEST
// 1. when number is not a valid number /factorial?number=abc
// 2. /factorial?number=
// 3. Method other than GET would be MethodNotFound
//

package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Number int `json:"data"`
}

func main() {
	e := gin.New()

	e.GET("/factorial", func(ctx *gin.Context) {
		numberString := ctx.Query("number")

		num, err := strconv.Atoi(numberString)
		if err != nil {
			log.Println("error converting number to string", numberString, err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if num < 0 {
			log.Println("negative number received",num)
			ctx.AbortWithError(http.StatusBadRequest,errors.New("negative number"))
			return
		}

		fact := factorial(num)

		resp := Data{
			Number: fact,
		}

		b, _ := json.Marshal(resp)

		ctx.Writer.Write(b)
		ctx.Writer.WriteHeader(http.StatusAccepted)
	})

	e.Run()
}

func factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}
