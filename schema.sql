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
