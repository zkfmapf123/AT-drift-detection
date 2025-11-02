package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/at-plan/client"
	"github.com/zkfmapf123/at-plan/usecase"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "A CLI tool for managing your at-plan",
	Long:  `A CLI tool for managing your at-plan`,
	Run: func(cmd *cobra.Command, args []string) {
		atRequest := cmd.Context().Value("atlantis").(usecase.AtlantisRequest)
		at := client.NewAtlantisRequest(atRequest)

		// validate
		if err := validAtlantis(at); err != nil {
			log.Println("failed to validate atlantis", err)
		}

	},
}

func validAtlantis(at client.AtlantisParams) error {

	var err error

	if err = at.ValidURL(); err != nil {
		return err
	}

	if err = at.ValidConfigFile(); err != nil {
		return err
	}

	return nil
}
