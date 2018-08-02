package main

import (
        "github.com/urfave/cli"
        "log"
        "os"
        "fmt"
)

//need a volume-key and volume-vm-key




var (
        version = "0.0.0"
        build   = "0"
)

func main() {
        app := cli.NewApp()
        app.Name = "netapp  snapshot"
        app.Usage = "netapp snapshot"
        app.Action = run
        app.Version = fmt.Sprintf("%s+%s", version, build)
        app.Flags = []cli.Flag{

                cli.StringFlag{
                        Name:   "url",
                        Usage:  "url to the netapp api",
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
