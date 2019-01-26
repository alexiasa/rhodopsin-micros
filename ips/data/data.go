package data

import "C"
import (
	"rhodopsin-micros/ips/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Define struct for mongo collection
type IpRepository struct {
	C *mgo.Collection
}

// Return all of the IP addresses in the collection.
func (r *IpRepository) GetAll() []models.IpDetails {
	var ips []models.IpDetails
	iter := r.C.Find(nil).Iter()
	result := models.IpDetails{}
	for iter.Next(&result) {
		ips = append(ips, result)
	}
	return ips
}

// Return an IP address by internal MongoDB ID
func (r *IpRepository) GetById(id string) (ip models.IpDetails, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&ip)
	return
}

// Return an IP address by ipaddr value using query
func (r *IpRepository) GetByIpaddr(ipaddr string) (ip models.IpDetails, err error) {
	err = r.C.Find(bson.M{"ipaddr": ipaddr}).One(&ip)
	return
}
