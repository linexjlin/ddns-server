package main

import (
	"context"
	"net"
	"net/http"
	"strings"
)

func clientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func UpdateDNS(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	domain, ip := r.FormValue("domain"), r.FormValue("ip")
	if ip == "" {
		ip = clientPublicIP(r)
	}

	if domain == "" || ip == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
