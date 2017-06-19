package main

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"bytes"
	"encoding/json"
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

	client := http.DefaultClient

	object := PostObject{
		StorageVMKey: p.StorageVmKey,
		VolumeKey:    p.VolumeKey,
		Name:         p.NamePrefix,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)
	_, err := client.Post(p.URL, "applications/json", b)
	if err != nil {
		return err
	}

	return nil

}
