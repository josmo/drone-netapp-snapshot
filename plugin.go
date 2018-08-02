package main

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type Plugin struct {
	URL          string
	UserName     string
	UserPassword string
	StorageVmKey string
	VolumeKey    string
	NamePrefix   string
}
type PostObject struct {
	StorageVMKey string `json:"storage_vm_key"`
	VolumeKey    string `json:"volume_key"`
	Name         string `json:"name"`
}

func (p *Plugin) Exec() error {
	log.Info("NetAPP Plugin built")
	fmt.Printf("%+v\n", p)

	if p.URL == "" || p.UserName == "" || p.UserPassword == "" {
		return errors.New("Eek: Must have url, user-name, secret")
	}

	// "Basic": base64encode(user:password)
	object := PostObject{
		StorageVMKey: p.StorageVmKey,
		VolumeKey:    p.VolumeKey,
		Name:         p.NamePrefix,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("POST", p.URL, b)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(p.UserName, p.UserPassword))

	//response , err := client.Post(p.URL, "applications/json", b)
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", response)

	return nil

}
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
