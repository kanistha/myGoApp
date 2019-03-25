package http_test

import (
	"testing"
	"net/http/httptest"
	gohttp "net/http"

)

func TestLoginHandlerGet(t *testing.T) {

	req, err := gohttp.NewRequest(gohttp.MethodGet,"/", nil)
	if(err != nil){
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handle := gohttp.HandleFunc(login)


	result.ser

}

