package models

type Ambulance struct {
  Id_ambulan int `json:"id_ambulan"`
  Nm_ambulan string `json:"nm_ambulan"`
  Deskripsi string `json:"deskripsi"`
  Tarif int `json:"tarif"`
  Foto string `json:"foto"`
  Status string `json:"status"`
  No_polisi string `json:"no_polisi"`
}

