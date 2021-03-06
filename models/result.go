package models

type (
  Result struct {
    Code int32 `json:"code"`
    Msg string `json:"msg"`
    ResNum int32 `json:"num"`
    Data []Recs `json:"data"`
    Distance float64 `json:"distance"`
  }

  Recs interface{}
)
