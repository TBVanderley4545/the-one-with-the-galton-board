package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/TBVanderley4545/the-one-with-the-galton-board/pkg/utils"
)

// Function to handle the root route request.
func RootRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	numberToDouble, err := strconv.ParseFloat(r.URL.Query().Get("double"), 64)

	doubledNumber := utils.Double(numberToDouble)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		log.Print("Query param parsing error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, doubledNumber)
}
