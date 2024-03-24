package app

import (
	"fmt"
	"net/http"
	"story-api/pkg/e"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	fmt.Printf("after bind cmd...\n")
	fmt.Printf("form: %v\n", form)
	if err != nil {
		fmt.Printf("trap1 Bino!\n")
		fmt.Printf("err: %v\n", err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		fmt.Printf("trap2 Bino!\n")
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		fmt.Println(valid.Errors)
		// MarkErrors(valid.Errors)
		fmt.Printf("trap3 Bino!\n")
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
