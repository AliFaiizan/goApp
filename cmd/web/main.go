package main

import (
	"github.com/alifaiizan/goApp/pkg/Routes"
	"net/http"

)



func main() {
	Routes.Routes()
	_ = http.ListenAndServe(":8080", nil)

}
