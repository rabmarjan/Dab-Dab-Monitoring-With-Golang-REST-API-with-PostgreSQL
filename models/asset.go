package models

import (
	"time"

	"github.com/lib/pq"

	//"github.com/fxtlabs/date"
	_ "github.com/lib/pq"
)

type Date struct {
	Year  int        // Year. E.g., 2009.
	Month time.Month // Month is 1 - 12. 0 means unspecified.
	Day   int
}

type Header struct {
	TraceID                    string    `json:"traceId"`
	ResponseVersion            string    `json:"responseVersion"`
	RequestReceivedTime        time.Time `json:"requestReceivedTime"`
	RequestID                  string    `json:"requestId"`
	ResponseTime               time.Time `json:"responseTime"`
	HopCount                   int       `json:"hopCount"`
	ResponseMessage            string    `json:"responseMessage"`
	ResponseProcessingTimeInMs int       `json:"responseProcessingTimeInMs"`
	ResponseCode               string    `json:"responseCode"`
}

type Asset struct {
	Oid               string      `json:"oid"`
	OrganizationOid   string      `json:"organizationOid"`
	CustomerOid       string      `json:"customerOid"`
	SiteOid           string      `json:"siteOid"`
	CategoryOid       string      `json:"categoryOid"`
	ManufacturerOid   string      `json:"manufacturerOid"`
	ModelOid          string      `json:"modelOid"`
	AssetName         string      `json:"assetName"`
	ProductSerial     string      `json:"productSerial"`
	AssetID           string      `json:"assetID"`
	PurchaseDate      time.Time   `json:"purchaseDate,omitempty"`
	ShipmentDate      pq.NullTime `json:"shipmentDate,omitempty"`
	DeliveryDate      pq.NullTime `json:"deliveryDate,omitempty"`
	EolDate           pq.NullTime `json:"eolDate,omitempty"`
	EosDate           pq.NullTime `json:"eosDate,omitempty"`
	SpecificationJSON string      `json:"specificationJson"`
	ConfigurationJSON string      `json:"configurationJson"`
	CredentialJSON    string      `json:"credentialJson"`
	// CustomerName      string    `json:"customerName"`
	// SiteName          string    `json:"siteName"`
	// CategoryName      string    `json:"categoryName"`
	// ManufacturerName  string    `json:"manufacturerName"`
	// ModelName         string    `json:"modelName"`
}

type Assets struct {
	AllAssets []Asset `json:"data"`
}
