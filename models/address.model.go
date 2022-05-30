package models

import (
	"net/http"
	"rest-api-e-comerce/db"
)

type Provinsi struct {
	ProvinceID string `json:"id_province"`
	Province   string `json:"province"`
}
type City struct {
	CityID     string `json:"id_city"`
	ProvinceID string `json:"province_id"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}
type Subdistrict struct {
	Id              string `json:"id"`
	ProvinceID      string `json:"province_id"`
	CityID          string `json:"city_id"`
	SubdistrictName string `json:"subdistrict_name"`
}
type Courier struct {
	Id          string `json:"id"`
	Courier     string `json:"courier"`
	CourierName string `json:"courier_name"`
	Status      string `json:"status"`
}

func FetchAllProvinsi() (Response, error) {
	var res Response
	var prov Provinsi
	var arrProv []Provinsi

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM provinsi"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&prov.ProvinceID, &prov.Province)
		if err != nil {
			return res, err
		}
		arrProv = append(arrProv, prov)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST PROVINSI"
	res.Data = arrProv

	return res, nil
}
func FetchAllCourier() (Response, error) {
	var res Response
	var courier Courier
	var arrCourier []Courier

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM courier"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&courier.Id, &courier.Courier, &courier.CourierName, &courier.Status)
		if err != nil {
			return res, err
		}
		arrCourier = append(arrCourier, courier)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST KURIR"
	res.Data = arrCourier

	return res, nil
}
func FetchAllCourierAktif() (Response, error) {
	var res Response
	var courier Courier
	var arrCourier []Courier

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM courier where status = 'Aktif'"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&courier.Id, &courier.Courier, &courier.CourierName, &courier.Status)
		if err != nil {
			return res, err
		}
		arrCourier = append(arrCourier, courier)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST KURIR"
	res.Data = arrCourier

	return res, nil
}
func FetchAllCity(ID string) (Response, error) {
	var res Response
	var city City
	var arrCity []City

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM city where province_id = " + ID

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&city.CityID, &city.ProvinceID, &city.Type, &city.CityName, &city.PostalCode)
		if err != nil {
			return res, err
		}
		arrCity = append(arrCity, city)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST PROVINSI"
	res.Data = arrCity

	return res, nil
}
func FetchAllSubdistrict(ID string) (Response, error) {
	var res Response
	var sub Subdistrict
	var arrSub []Subdistrict

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM subdistrict where city_id = " + ID

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&sub.Id, &sub.ProvinceID, &sub.CityID, &sub.SubdistrictName)
		if err != nil {
			return res, err
		}
		arrSub = append(arrSub, sub)
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET LIST PROVINSI"
	res.Data = arrSub

	return res, nil
}
