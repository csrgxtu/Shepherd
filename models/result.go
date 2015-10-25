package models

type (
  Result struct {
    Code int32 `json:"code"`
    Msg string `json:"msg"`
    ResNum int32 `json:"resNum"`
    Data []recs `json:"data"`
  }

  recs interface{}
)
