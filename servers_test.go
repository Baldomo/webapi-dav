package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var (
	endpoints = []string{
		// "api",
		"api/about",
		"api/comunicati",
		"api/comunicati/genitori",
		"api/comunicati/genitori/5",
		"api/comunicati/docenti",
		"api/comunicati/docenti/5",
		"api/comunicati/studenti",
		"api/comunicati/studenti/5",
		"api/version",
		// "api/teapot", --> 418
	}

	client = &http.Client{
		Timeout: 1 * time.Second,
	}
)

func TestEndpoints(t *testing.T) {
	StartServers()
	// HTTP
	for _, endpoint := range endpoints {
		t.Logf("Testing %s", endpoint)
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1%s/%s", GetConfig().HTTP.Port, endpoint), nil)
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatal("Risposta non OK 200")
		}
		raw, _ := ioutil.ReadAll(resp.Body)
		if !json.Valid(raw) {
			t.Errorf("JSON invalido per %s, su HTTP", endpoint)
		}
	}

	// HTTPS
	for _, endpoint := range endpoints {
		t.Logf("Testing %s", endpoint)
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1/%s", endpoint), nil)
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatal("Risposta non OK 200")
		}
		raw, _ := ioutil.ReadAll(resp.Body)
		if !json.Valid(raw) {
			t.Errorf("JSON invalido per %s, su HTTPS", endpoint)
		}
	}
}
