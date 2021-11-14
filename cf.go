package main

import (
	"context"
	"github.com/weppos/publicsuffix-go/publicsuffix"

	"github.com/cloudflare/cloudflare-go"
)

func GetDomainSuffix(domain string) string {
	s, _ := publicsuffix.Domain(domain)
	return s
}

type CloudFlare struct {
	client *cloudflare.API
}

func NewCloudFlare(key, email string) (*CloudFlare, error) {
	client, err := cloudflare.New(key, email)
	if err != nil {
		return nil, err
	}
	return &CloudFlare{
		client: client,
	}, nil
}

func (c *CloudFlare) GetDomainZoneID(domain string) (string, error) {
	zone := GetDomainSuffix(domain)
	id, err := c.client.ZoneIDByName(zone)
	if err != nil {
		return "", nil
	}
	return id, nil
}

func (c *CloudFlare) UpdateIP(ctx context.Context, domain, ip string) error {
	zid, err := c.GetDomainZoneID(domain)
	if err != nil {
		return err
	}
	rs, err := c.client.DNSRecords(ctx, zid, cloudflare.DNSRecord{
		Name: domain,
	})
	if err != nil {
		return err
	}
	for _, r := range rs {
		if r.Type == "A" {
			if r.Content == ip {
				continue
			}
			//oldIP := r.Content
			r.Content = ip
			err = c.client.UpdateDNSRecord(ctx, zid, r.ID, r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
