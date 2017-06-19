package main

import (
        "github.com/codegangsta/cli"
        "log"
        "os"
)

//need a volume-key and volume-vm-key


type Rancher struct {
        Url            string `json:"url"`
        UserName      string `json:"user_name"`
        UserPassword      string `json:"user_password"`
        StorageVMKey      string `json:"storage_vm_key"`
        VolumeKey         string `json:"volume_key"`
        NamePrefix        string `json:"name_prefix"`
        //VolumeName        string `json:"volume_name"`
}

var version string // build number set at compile-time

func main() {
        app := cli.NewApp()
        app.Name = "netapp  snapshot"
        app.Usage = "netapp snapshot"
        app.Action = run
        app.Version = version
        app.Flags = []cli.Flag{

                cli.StringFlag{
                        Name:   "url",
                        Usage:  "url to the rancher api",
                        EnvVar: "PLUGIN_URL",
                },
                cli.StringFlag{
                        Name:   "user-name",
                        Usage:  "netapp user name",
                        EnvVar: "PLUGIN_USER_NAME, NETAPP_USER_NAME",
                },
                cli.StringFlag{
                        Name:   "user-password",
                        Usage:  "netapp user password",
                        EnvVar: "PLUGIN_USER_PASSWORD, NETAPP_USER_PASSWORD",
                },
                cli.StringFlag{
                        Name:   "storage-vm-key",
                        Usage:  "Storage vm key",
                        EnvVar: "PLUGIN_STORAGE_VM_KEY",
                },
                cli.StringFlag{
                        Name:   "volume-key",
                        Usage:  "Volume key",
                        EnvVar: "PLUGIN_VOLUME_KEY",
                },
                cli.StringFlag{
                        Name:   "name-prefix",
                        Usage:  "snapshot prefix",
                        EnvVar: "PLUGIN_NAME_PREFIX",
                },
        }

        if err := app.Run(os.Args); err != nil {
                log.Fatal(err)
        }
}

func run(c *cli.Context) error {
        plugin := Plugin{
                URL:            c.String("url"),
                UserName:            c.String("user-name"),
                UserPassword:         c.String("user-password"),
                StorageVmKey:        c.String("storage-vm-key"),
                VolumeKey:    c.String("volume-key"),
                NamePrefix:     c.String("name-prefix"),
        }
        return plugin.Exec()
}
