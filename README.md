# Monitoring-System-With-Golang-REST-API-with-PostgreSQL & SQLite Database
A Go based REST API with PostgreSQL and SQLite Database

## How to Run the Project on Your Local Machine with PostgreSQL
1. Create a database (Run the below command)
```sql
-- Database: lumoswg

-- DROP DATABASE lumoswg;

CREATE DATABASE lumoswg
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'en_US.UTF-8'
       LC_CTYPE = 'en_US.UTF-8'
       CONNECTION LIMIT = -1;
       
-- Table: public.asset

-- DROP TABLE public.asset;

CREATE TABLE public.asset
(
  oid character varying(128) NOT NULL,
  organizationoid character varying(128) NOT NULL,
  customeroid character varying(128) NOT NULL,
  siteoid character varying(128) NOT NULL,
  categoryoid character varying(128) NOT NULL,
  manufactureroid character varying(128) NOT NULL,
  modeloid character varying(128) NOT NULL,
  assetname character varying(256) NOT NULL,
  productserial character varying(256) NOT NULL,
  assetid character varying(256),
  purchasedate date,
  shipmentdate date,
  deliverydate date,
  eoldate date,
  eosdate date,
  specificationjson text NOT NULL DEFAULT '{}'::text,
  configurationjson text NOT NULL DEFAULT '{}'::text,
  credentialjson text NOT NULL DEFAULT '{}'::text,
  datajson text NOT NULL DEFAULT '{}'::text,
  createdby character varying(128) NOT NULL DEFAULT 'System'::character varying,
  createdon timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedby character varying(128),
  updatedon timestamp without time zone,
  CONSTRAINT pk_asset PRIMARY KEY (oid)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.asset
  OWNER TO postgres;

```
2. Install Go and Set a Path variable in .bashrc or .profile file

```bash
export GOPATH=$HOME/go
export GOROOT=/usr/lib/go-1.10
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
3. Put the project in src directory and rename project `Monitoring-System-With-Golang-REST-API-with-PostgreSQL-and-SQLite` to `goweb`
```bash
cd $HOME/go/src
mkdir workspace
## put the goweb with all source code into workspace directory so directory will look like following
$HOME/go/src/workspace/goweb
## All source file should inside of the goweb directory
```
4. Run the below command (It will install all project dependency)
```bash
go get
```
5. Change Database credential in db.go file with your db credentials

### Base url
localhost:9010/lumos/

### API end point
1. asset/asset/v1/get-list
2. asset/asset/v1/save
3. asset/asset/v1/update

### Concatenate base url with api end point so it look like below url
localhost:9010/lumos/asset/asset/v1/get-list

### For running service run below command where main.go file located
```bash
go run main.go
```
### Getting data from service
curl -i -X POST -H "Content-Type: application/json" -d '{"header":{"requestId":"f130d221-4ae6-4d09-bf2e-75d4f194a469","requestType":"asset/asset/v1/get-list","requestClient":"lumos","requestSource":"curl","requestSourceService":"terminal","requestVersion":"1.0","requestTimeoutInSeconds":30,"requestRetryCount":0,"hopCount":1,"traceId":"171206IPLI","requestTime":"2018-09-11T08:20:39.778Z"},"meta":{},"body":{"searchParam":{"offSet":0,"limit":10}}}' http://localhost:9010/lumos/asset/asset/v1/get-list

## How to Run the Project on Your Local Machine with SQLite

```sql
CREATE TABLE asset
(
  oid character varying(128) NOT NULL,
  organizationoid character varying(128) NOT NULL,
  customeroid character varying(128) NOT NULL,
  siteoid character varying(128) NOT NULL,
  categoryoid character varying(128) NOT NULL,
  manufactureroid character varying(128) NOT NULL,
  modeloid character varying(128) NOT NULL,
  assetname character varying(256) NOT NULL,
  productserial character varying(256) NOT NULL,
  assetid character varying(256),
  purchasedate date,
  shipmentdate date,
  deliverydate date,
  eoldate date,
  eosdate date,
  specificationjson text DEFAULT '{}',
  configurationjson text  DEFAULT '{}',
  credentialjson text   DEFAULT '{}',
  datajson text DEFAULT '{}',
  createdby character varying(128) DEFAULT '',
  createdon timestamp without time zone,
  updatedby character varying(128) DEFAULT '',
  updatedon timestamp without time zone,
  CONSTRAINT pk_asset PRIMARY KEY (oid)
);
```
### Get all asset from service
#### For insert
curl -i -X POST -H "Content-Type: application/json" -d '{"header":{"requestId":"f9f3dc67-e830-41b2-a45d-43cc3e3bb742","requestType":"asset/asset/sqlite/v1/save","requestClient":"lumos","requestSource":"curl","requestSourceService":"terminal","requestVersion":"1.0","requestTimeoutInSeconds":30,"requestRetryCount":0,"hopCount":1,"traceId":"171206IPLI","requestTime":"2018-09-11T09:05:38.199Z"},"meta":{},"body":{"oid":"120marjan27","organizationOid":"ORG-01","customerOid":"Cust-01","siteOid":"Site-01","categoryOid":"cat-01",
"manufacturerOid":"OEM-01","modelOid":"AM-01","assetName":"1234asdf","productSerial":"123456789",
"assetID":"1234","purchaseDate":"2018-07-31T10:51:00.603Z"}}' http://localhost:9010/lumos/asset/asset/sqlite/v1/save

#### For get all asset
curl -i -X POST -H "Content-Type: application/json" -d '{"header":{"requestId":"f130d221-4ae6-4d09-bf2e-75d4f194a469","requestType":"asset/asset/sqlite/v1/get-list","requestClient":"lumos","requestSource":"curl","requestSourceService":"terminal","requestVersion":"1.0","requestTimeoutInSeconds":30,"requestRetryCount":0,"hopCount":1,"traceId":"171206IPLI","requestTime":"2018-09-11T08:20:39.778Z"},"meta":{},"body":{"searchParam":{"offSet":0,"limit":10}}}' http://localhost:9010/lumos/asset/asset/sqlite/v1/get-list
