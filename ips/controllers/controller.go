package controllers

import (
	"encoding/json"
	"net/http"
	"rhodopsin-micros/ips/data"
)

func GetIps(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("ips")
	repo := &data.IpRepository{C: col}
	ips := repo.GetAll()
	j, err := json.Marshal(ips)
	if err != nil {
		http.Error(w, "oh noes! something bad happened", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

// todo: implement controller to return single IP address - line 32
/*func GetIpAddr(ipaddr string) (w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("ips")
	repo := &data.IpRepository{C: col}
	ip := repo.GetByIpaddr(ipaddr)
	j, err := json.Marshal(ip)
	if err != nil {
		http.Error(w, "oh noes! something bad happened", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}*/