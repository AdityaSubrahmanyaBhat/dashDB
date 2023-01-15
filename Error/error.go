package error

import (
	"fmt"
	color "github.com/AdityaSubrahmanyaBhat/dashDB/Colors"
)

func HandleError(err error) {
	if err != nil {
		fmt.Printf(color.GetColor("Red"), "Error : %v",err, color.GetColor("Reset"))
	}
}