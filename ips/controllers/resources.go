package controllers

import "rhodopsin-micros/ips/models"

// JSON Resource Models


type (
	// GET all IPs
	IpsResource struct {
		Data []models.IpDetails `json:"data"`
	}
	// GET single IP
	IdResource struct {
		Data models.IpDetails   `json:"data"`
	}
	// GET malicious IPs
	MalResource struct {
		Data []models.IpDetails	`json:"data"`
	}
)