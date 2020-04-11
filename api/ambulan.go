package api

import (
  "contoh_web/http_json" // module yang berisi response atau request format json
  "contoh_web/database" // module yang berisi koneksi ke database
  "contoh_web/models" // module yang berisi models atau proses ke database
  "net/http" // module yang berisi http request atau response. untuk membuat web
  "encoding/json" // module untuk encode/decode json
  "github.com/gorilla/mux" // module untuk router, serta mengambil query string
)

// method untuk mengembalikan data ambulan secara keseluruhan
func GetAmbulan(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json");
  
  db := database.Connect(); // koneksi kedatabase, dari module database
  result := http_json.NewResponse(); // mempersiapkan response JSON. dari module http_json
  data := []models.Ambulance{} // variabel penampung hasil database
  
  err := db.Select(&data, "SELECT * FROM tb_ambulan") // proses pengambilan kedatabase
  if err != nil { // jika error, maka ubah response JSON tadi ke mode error
    result.Code = 503; // kode 503 menandakan error
    result.Message = err.Error(); // ambil pesan error dan tampilkan ke response JSON
    result.Data = nil // data pada response JSON dikosongkan
  } else {
    result.Data = data // jika tidak error, maka hasil select database dimasukkan ke response JSON
  }

  json.NewEncoder(res).Encode(result) // menampilkan response JSON ke klien
}

func GetAmbulanById(res http.ResponseWriter, req *http.Request) {
  db := database.Connect();
  result := http_json.NewResponse();
  data := models.Ambulance{}
  
  qs := mux.Vars(req); // untuk mengambil data dari URL seperti GET atau parameter route
  err := db.Get(&data, "SELECT * FROM tb_ambulan WHERE id_ambulan=?", qs["id_ambulan"]);
  if err != nil {
    result.Code = 503;
    result.Message = err.Error();
    result.Data = nil
  } else {
    result.Data = data
  }

  json.NewEncoder(res).Encode(result)
}

func InsertAmbulan(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json");
  
  db := database.Connect();
  result := http_json.NewResponse();
  
  data_insert := map[string]interface{}{
    "nm_ambulan": req.FormValue("nm_ambulan"),
    "deskripsi": req.FormValue("deskripsi"),
    "tarif": req.FormValue("tarif"),
    "foto": req.FormValue("foto"),
    "status": req.FormValue("status"),
    "no_polisi": req.FormValue("no_polisi"),
  }
  
  id, err := db.NamedExec(`INSERT INTO tb_ambulan (nm_ambulan, deskripsi, tarif, foto, status, no_polisi) VALUES (:nm_ambulan, :deskripsi, :tarif, :foto, :status, :no_polisi)`, data_insert)
  if err != nil {
    result.Code = 503;
    result.Message = err.Error();
    result.Data = nil
  } else {
    result.Data = id
  }

  json.NewEncoder(res).Encode(result)
}

func UpdateAmbulan(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json");
  
  db := database.Connect();
  result := http_json.NewResponse();
  
  qs := mux.Vars(req);
  
  data_insert := map[string]interface{}{
    "id_ambulan": qs["id_ambulan"],
    "nm_ambulan": req.FormValue("nm_ambulan"),
    "deskripsi": req.FormValue("deskripsi"),
    "tarif": req.FormValue("tarif"),
    "foto": req.FormValue("foto"),
    "status": req.FormValue("status"),
    "no_polisi": req.FormValue("no_polisi"),
  }
  
  id, err := db.NamedExec(`UPDATE tb_ambulan SET nm_ambulan = :nm_ambulan, deskripsi = :deskripsi, tarif = :tarif, foto = :foto, status = :status, no_polisi = :no_polisi WHERE id_ambulan = :id_ambulan`, data_insert)
  if err != nil {
    result.Code = 503;
    result.Message = err.Error();
    result.Data = nil
  } else {
    result.Data = id
  }

  json.NewEncoder(res).Encode(result)
}


func DeleteAmbulan(res http.ResponseWriter, req *http.Request) {
  res.Header().Set("Content-Type", "application/json");
  
  db := database.Connect();
  result := http_json.NewResponse();
  
  qs := mux.Vars(req);
  
  data_delete := map[string]interface{}{
    "id_ambulan": qs["id_ambulan"],
  }
  
  id, err := db.NamedExec(`DELETE FROM tb_ambulan WHERE id_ambulan = :id_ambulan`, data_delete)
  if err != nil {
    result.Code = 503;
    result.Message = err.Error();
    result.Data = nil
  } else {
    result.Data = id
  }

  json.NewEncoder(res).Encode(result)
}
