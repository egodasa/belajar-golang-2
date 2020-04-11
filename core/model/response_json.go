package api

import (
  "net/http"
  "encoding/json"
)

type Response struct {
  Code int `json:"code,omitempty"`
  Message string `json:"message,omitempty"`
  Data []interface{} `json:"data,omitempty"`
}

func NewResponse() Response {
  return Response{
    Code: 200,
    Message: "Ok",
    Data: []interface{}{},
  }
}

func GetAmbulan(res http.ResponseWriter, req *http.Request) {
  result := NewResponse();
  
  result.Code = 503;
  
  //~ result.Code = 503
  //~ result.Message = string(err.Error())
  //~ result.Data = []interface{}{}
  
  json.NewEncoder(res).Encode(result)
}
