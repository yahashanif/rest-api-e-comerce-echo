package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-api-e-comerce/db"
	"strings"
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

type Address struct {
	Id            string `json:"id"`
	IdUser        string `json:"id_user"`
	IdProvince    string `json:"id_province"`
	Province      string `json:"province"`
	IdCity        string `json:"id_city"`
	City          string `json:"city"`
	IdSubdistrict string `json:"id_subdistrict"`
	Subdistrict   string `json:"subdstrict"`
	PostalCode    string `json:"postal_code"`
	DetailAddress string `json:"detail_address"`
}

func FetchAllProvinsi() (Response, error) {
	var res Response
	var prov Provinsi
	var arrProv []Provinsi

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM provinsi order by name_provinsi"

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

func Ongkir(origin, originType, destination, destinationType, weight string) (Response, error) {
	var res Response
	con := db.CreateCon()
	url := "https://pro.rajaongkir.com/api/cost"

	sqlStatement := "select courier from courier where status='Aktif'"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	var cou string
	var courier string

	for rows.Next() {
		err = rows.Scan(&cou)
		if err != nil {
			return res, err
		}
		courier = courier + ":" + cou

	}

	payload := strings.NewReader("origin=" + origin + "&originType=" + originType + "&destination=" + destination + "&destinationType=" + destinationType + "&weight=" + weight + "&courier=" + courier[1:len(courier)])

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "80ebd4a124cc35bd4322a8105e34c20f")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	result, _ := http.DefaultClient.Do(req)

	defer result.Body.Close()
	body, _ := ioutil.ReadAll(result.Body)
	jsonMap := make(map[string]map[string]interface{})
	json.Unmarshal([]byte(body), &jsonMap)
	// fmt.Println(jsonMap)
	// fmt.Println(result)
	// fmt.Println(string(body))

	res.Data = jsonMap["rajaongkir"]
	return res, nil
}

func AddAddress(a *Address) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO `address` ( `id_user`, `id_province`, `province`, `id_city`, `city`, `id_subdistrict`, `subdistrict`, `postal_code`, `detail_address`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(a.IdUser, a.IdProvince, a.Province, a.IdCity, a.City, a.IdSubdistrict, a.Subdistrict, a.PostalCode, a.DetailAddress)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES ADD ADDRESS"
	res.Data = result

	return res, nil

}
func EditAddress(a *Address) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE `address` SET `id_province` = ?, `province` = ?, `id_city` = ?, `city` = ?, `id_subdistrict` = ?, `subdistrict` = ?, `postal_code` = ?, `detail_address` = ? WHERE `address`.`id` = ? and `address`.`id_user` = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(a.IdProvince, a.Province, a.IdCity, a.City, a.IdSubdistrict, a.Subdistrict, a.PostalCode, a.DetailAddress, a.Id, a.IdUser)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES UPDATED ADDRESS"
	res.Data = result

	return res, nil

}

func FetchAddressByIdUSer(idUser string) (Response, error) {
	var res Response
	con := db.CreateCon()

	var add Address

	sqlStatement := "SELECT * FROM `address` WHERE id_user = ?"

	err := con.QueryRow(sqlStatement, idUser).Scan(&add.Id, &add.IdUser, &add.IdProvince, &add.Province, &add.IdCity, &add.City, &add.IdSubdistrict, &add.Subdistrict, &add.PostalCode, &add.DetailAddress)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES GET DATA ADDRESS"
	res.Data = add
	return res, nil
}
