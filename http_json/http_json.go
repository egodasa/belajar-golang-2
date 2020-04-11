package http_json

type Response struct {
  Code int `json:"code,omitempty"`
  Message string `json:"message,omitempty"`
  Data interface{} `json:"data,omitempty"`
}

func NewResponse() Response {
  return Response{
    Code: 200,
    Message: "Ok",
  }
}
