# ddns-serve

[![Go](https://github.com/linexjlin/ddns-server/actions/workflows/go.yml/badge.svg)](https://github.com/linexjlin/ddns-server/actions/workflows/go.yml)

## Usage 
1. Run serve
````
CLOUDFLARE_API_KEY=keyxxx CLOUDFLARE_API_EMAIL=mail@gmail.com ./ddns-server
``````

2. Set IP
```
curl -v 'http://127.0.0.1:8010/UpdateDNS?domain=test.linkown.com&ip=127.0.0.2'
```
