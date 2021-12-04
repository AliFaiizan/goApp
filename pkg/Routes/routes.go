package Routes

import (
	"github.com/alifaiizan/goApp/pkg/Controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", Controller.Home)
	http.HandleFunc("/about", Controller.About)
}
