package gacha

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GachaIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%s\n", "gacha")
}

func GachaGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
