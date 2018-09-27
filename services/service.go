package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"log"
	"net/http"
	"workspace/goweb/db"
	"workspace/goweb/models"

	log "github.com/sirupsen/logrus"
)

// QueryRepos first fetches the repositories data from the db
func QueryRepos(repos *models.Assets) error {
	db, err := db.DbConnection()
	defer db.Close()
	rows, err := db.Query(`
		SELECT
		Oid,              
		OrganizationOid,  
		CustomerOid,      
		SiteOid,          
		CategoryOid,      
		ManufacturerOid,  
		ModelOid,         
		AssetName,        
		ProductSerial,    
		AssetID,          
		PurchaseDate,     
		ShipmentDate,     
		DeliveryDate,     
		EolDate,          
		EosDate,          
		SpecificationJSON,
		ConfigurationJSON,
		CredentialJSON       
		FROM asset`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		asset := models.Asset{}
		err = rows.Scan(
			&asset.Oid,
			&asset.OrganizationOid,
			&asset.CustomerOid,
			&asset.SiteOid,
			&asset.CategoryOid,
			&asset.ManufacturerOid,
			&asset.ModelOid,
			&asset.AssetName,
			&asset.ProductSerial,
			&asset.AssetID,
			&asset.PurchaseDate,
			&asset.ShipmentDate,
			&asset.DeliveryDate,
			&asset.EolDate,
			&asset.EosDate,
			&asset.SpecificationJSON,
			&asset.ConfigurationJSON,
			&asset.CredentialJSON,
		)
		if err != nil {
			return err
		}
		repos.AllAssets = append(repos.AllAssets, asset)
	}
	log.Info("Get All List from Asset Table")
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

// QueryReposSave first fetches the repositories data from the db
func QueryReposSave(w http.ResponseWriter, r *http.Request) {
	db, err := db.DbConnection()
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var asset models.Asset
	var LastInsertId string
	log.Println(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // TODO.....
	log.Println("Request body", r.Body)              // TODO.....
	// log.Println(string(body))
	// var t models.Asset
	// err = json.Unmarshal(body, &t)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(t.Oid)
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		log.Fatal(err)
	}
	myOid, ok := obj["body"]
	var myJSON map[string]interface{}
	if ok {
		switch v := myOid.(type) {
		case map[string]interface{}:
			myJSON = v
		default:
			log.Println()
		}
	}
	log.Println(myJSON["siteOid"].(string))

	json.NewDecoder(r.Body).Decode(&asset)
	log.Println("Encoded Asset", &asset)
	log.Println(myJSON["oid"], myJSON["organizationOid"], myJSON["customerOid"],
		myJSON["siteOid"], myJSON["categoryOid"], myJSON["manufacturerOid"],
		myJSON["modelOid"], myJSON["assetName"], myJSON["productSerial"],
		myJSON["assetID"], myJSON["purchaseDate"])
	err = db.QueryRow(`insert into asset (
		oid, organizationOid, customerOid, siteOid,
		categoryOid, manufacturerOid,modelOid, assetName,
		productSerial, assetID, purchaseDate )values (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning oid;`,
		myJSON["oid"], myJSON["organizationOid"], myJSON["customerOid"],
		myJSON["siteOid"], myJSON["categoryOid"], myJSON["manufacturerOid"],
		myJSON["modelOid"], myJSON["assetName"], myJSON["productSerial"],
		myJSON["assetID"], myJSON["purchaseDate"]).Scan(&LastInsertId)
	if err != nil {
		panic(err)
	}
	log.Println(LastInsertId)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asset)
	log.Println("After JSON")

}

// QueryReposUpdate first update the repositories data from the db
func QueryReposUpdate(w http.ResponseWriter, r *http.Request) {
	db, err := db.DbConnection()
	defer db.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		log.Fatal(err)
	}
	myOid, ok := obj["body"]
	var myJSON map[string]interface{}
	if ok {
		switch v := myOid.(type) {
		case map[string]interface{}:
			myJSON = v
		default:
			log.Println()
		}
	}
	log.Println(myJSON["siteOid"].(string))
	stmt, err := db.Prepare(`update asset 
	set oid=$1,
		organizationOid=$2,
		customerOid=$3,
		siteOid=$4,
		categoryOid=$5,
		manufacturerOid=$6,
		modelOid=$7,
		assetName=$8,
		productSerial=$9,
		assetID=$10,
		purchaseDate=$11
	where oid=$1`)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(myJSON["oid"], myJSON["organizationOid"], myJSON["customerOid"],
		myJSON["siteOid"], myJSON["categoryOid"], myJSON["manufacturerOid"],
		myJSON["modelOid"], myJSON["assetName"], myJSON["productSerial"],
		myJSON["assetID"], myJSON["purchaseDate"])
	if err != nil {
		panic(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	log.Println(affect, "rows changed")
}

func QuerySaveSQLite(w http.ResponseWriter, r *http.Request) {
	db, err := db.SQLiteConn()
	defer db.Close()
	var asset models.Asset
	var LastInsertId string
	log.Println(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // TODO.....
	log.Println("Request body", r.Body)              // TODO.....
	obj := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		log.Fatal(err)
	}
	myOid, ok := obj["body"]
	var myJSON map[string]interface{}
	if ok {
		switch v := myOid.(type) {
		case map[string]interface{}:
			myJSON = v
		default:
			log.Println()
		}
	}
	log.Println(myJSON["siteOid"].(string))

	json.NewDecoder(r.Body).Decode(&asset)
	log.Println("Encoded Asset", &asset)
	stmt, err := db.Prepare(`insert into asset (
		oid, organizationOid, customerOid, siteOid,
		categoryOid, manufacturerOid,modelOid, assetName,
		productSerial, assetID, purchaseDate )values (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(myJSON["oid"], myJSON["organizationOid"], myJSON["customerOid"],
		myJSON["siteOid"], myJSON["categoryOid"], myJSON["manufacturerOid"],
		myJSON["modelOid"], myJSON["assetName"], myJSON["productSerial"],
		myJSON["assetID"], myJSON["purchaseDate"])
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	log.Println(LastInsertId)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asset)
	log.Println("After JSON")

}

//QueryReposSQLite get data from SQLite db
func QueryReposSQLite(repos *models.Assets) error {
	db, err := db.SQLiteConn()
	defer db.Close()
	rows, err := db.Query(`
		SELECT
		Oid,              
		OrganizationOid,  
		CustomerOid,      
		SiteOid,          
		CategoryOid,      
		ManufacturerOid,  
		ModelOid,         
		AssetName,        
		ProductSerial,    
		AssetID,          
		PurchaseDate,     
		ShipmentDate,     
		DeliveryDate,     
		EolDate,          
		EosDate,          
		SpecificationJSON,
		ConfigurationJSON,
		CredentialJSON       
		FROM asset`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		asset := models.Asset{}
		err = rows.Scan(
			&asset.Oid,
			&asset.OrganizationOid,
			&asset.CustomerOid,
			&asset.SiteOid,
			&asset.CategoryOid,
			&asset.ManufacturerOid,
			&asset.ModelOid,
			&asset.AssetName,
			&asset.ProductSerial,
			&asset.AssetID,
			&asset.PurchaseDate,
			&asset.ShipmentDate,
			&asset.DeliveryDate,
			&asset.EolDate,
			&asset.EosDate,
			&asset.SpecificationJSON,
			&asset.ConfigurationJSON,
			&asset.CredentialJSON,
		)
		if err != nil {
			return err
		}
		repos.AllAssets = append(repos.AllAssets, asset)
	}
	log.Info("Get All List from Asset Table")
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
