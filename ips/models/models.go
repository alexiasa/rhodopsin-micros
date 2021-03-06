package models

import (
  "gopkg.in/mgo.v2/bson"
)

type (

//  Location struct {
//    Lat       int           `json:"lat"`
//    Lon       int           `json:"lon"`
//  }

  IpDetails struct {
    Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
    Ipaddr    string        `json:"ipaddr"`
    Asn       string        `json:"asn"`
    Location  string        `json:"location"`
    Malicious bool          `json:"malicious"`

  }


)

