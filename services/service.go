package services

import (
	"workspace/goweb/db"
	"workspace/goweb/models"
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
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
