package condGacha

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GachaGet(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	res := GachaExecute(ps.ByName("userID"))
	fmt.Fprintf(w, "%s\n", res)
}
