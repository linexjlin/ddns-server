package main

import (
	"context"
	"net/http"
)

func UpdateDNS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	domain, ip := r.FormValue("domain"), r.FormValue("ip")
	if err := service.CloudFlareA(domain, ip); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

type Service struct {
	cf *CloudFlare
}

func (s *Service) CloudFlareA(domain, ip string) error {
	ctx := context.Background()
	return s.cf.UpdateIP(ctx, domain, ip)
}
