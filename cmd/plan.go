package cmd

import (
	"fmt"

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
		if err := validAtlantis(&at); err != nil {
			panic(err)
		}

		outputs := at.Plan()

		// outputs
		for dir, outs := range outputs {
			fmt.Println(">>>", dir)
			for _, out := range outs.ProjectResults {
				fmt.Println(out.PlanSuccess.TerraformOutput)
			}
		}

		// slack send
	},
}

func validAtlantis(at *client.AtlantisParams) error {

	var err error

	if err = at.ValidURL(); err != nil {
		return err
	}

	if err = at.ValidRepository(); err != nil {
		return err
	}

	if err := at.ValidConfigFile(); err != nil {
		return err
	}

	if _, err = at.SetConfigParmas(); err != nil {
		return err
	}

	return nil
}
