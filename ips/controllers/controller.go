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

	params, ok := r.URL.Query()["addr"]

	if !ok || (len(params[0]) < 1) {
		// this branch should happen if there isn't a query string value - null bool = true or params is nil
		ips := repo.GetAll()
		j, err := json.Marshal(ips)
		if err != nil {
			http.Error(w, "oh noes! something bad happened", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	} else {
		// this branch should happen if there is a query string value ?Addr=something in params
		ip, err := repo.GetByIpaddr(params[0])
		j, err := json.Marshal(ip)
		if err != nil {
			http.Error(w, "oh noes! something bad happened", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)

	}

}

// todo: implement controller to return details for single IP found by address - line 32
/*func GetIpAddr(ipaddr string) (w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("ips")
	repo := &data.IpRepository{C: col}
	ip, err := repo.GetByIpaddr(ipaddr)
	j, err := json.Marshal(ip)
	if err != nil {
		http.Error(w, "oh noes! something bad happened", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}*/