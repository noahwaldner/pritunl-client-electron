package cmd

import (
	"os"
	"fmt"
	"encoding/json"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pritunl/pritunl-client-electron/cli/sprofile"
	"github.com/spf13/cobra"
)

type Profile struct {
    Id string
	Name string
	Connected bool
	ConnectedSince string
}


var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List profiles",
	Run: func(cmd *cobra.Command, args []string) {
		sprfls, err := sprofile.GetAll()
		if err != nil {
			panic(err)
			return
		}

		profiles := "["



		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"ID",
			"Name",
			"State",
			"Autostart",
			"Online For",
			"Server Address",
			"Client Address",
		})
		table.SetBorder(true)

		for _, sprfl := range sprfls {
			prof := &Profile{Id: sprfl.Id, Name:sprfl.FormatedName(), Connected: true, ConnectedSince: sprfl.Profile.FormatedTime()}
			stringified, err := json.Marshal(prof)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return;
			}
			profiles += strings.ToLower(string(stringified))
		}
		profiles += "]"
		fmt.Println(profiles)
	},
}
