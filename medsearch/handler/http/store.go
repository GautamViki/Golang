package http

import (
	"fmt"
	"net/http"
)
type StoreHandler struct{

}
func (s *StoreHandler) GetStore(w http.ResponseWriter, r *http.Request){
	fmt.Println("555555555555555555555555555555555")
	fmt.Println("Get store")
}