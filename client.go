package main

import (
	"crypto/tls"
	"github.com/bogdanovich/dns_resolver"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	scheme               = "https://"
	baseURL              = "127.0.0.1"
	port                 = ":9001"
	endpoint             = "/telemetry"
	localDNSResolverConf = "/etc/resolv.conf"
	DNSFallbackAddr      = "8.8.8.8"
)

func ResolveDNSName(hostName string) []string {
	resolver, err := dns_resolver.NewFromResolvConf(localDNSResolverConf)
	if err != nil {
		log.Fatal(err)
	}

	resolver.RetryTimes = 1

	log.Printf("Trying to resolve %v with local resolver...\n", hostName)
	ip, err := resolver.LookupHost(hostName)
	if err != nil {
		log.Fatal(err.Error())

		resolver = dns_resolver.New([]string{DNSFallbackAddr})
		log.Printf("Trying to resolve %v with %v resolver...", hostName, DNSFallbackAddr)
		ip, err = resolver.LookupHost(hostName)
		if err != nil {
			log.Fatal("Unable to resolve host:", err)
		}
	}
	result := make([]string, 0)
	for _, v := range ip {
		result = append(result, v.To16().String())
	}
	return result
}

func SendTelemetryData(client *http.Client, hostIP string) error {
	resp, err := client.PostForm(
		scheme+hostIP+port+endpoint,
		url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		log.Fatalf("Request error: %s", err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error on parsing response body: ", err)
		return err
	}

	log.Println("Got message from server:", string(body))
	return err
}

func main() {

	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   2 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ResponseHeaderTimeout: 5 * time.Second,
		MaxIdleConns:          10,
		IdleConnTimeout:       1 * time.Second,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   60 * time.Second,
	}

	IPs := ResolveDNSName(baseURL)

	if len(IPs) >= 1 {
		err := SendTelemetryData(client, IPs[0])
		if err != nil {
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s)
			randomHostIP := IPs[r.Intn(len(IPs))]
			SendTelemetryData(client, randomHostIP)
		}
	}
}
