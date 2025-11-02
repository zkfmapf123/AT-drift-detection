package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zkfmapf123/at-plan/client"
	"github.com/zkfmapf123/at-plan/usecase"
	"github.com/zkfmapf123/at-plan/utils"
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
		tfOutputs := make(map[string]string)

		for dir, outs := range outputs {
			for _, out := range outs.ProjectResults {
				tfOutputs[dir] = out.PlanSuccess.TerraformOutput
			}
		}

		// slack send
		if err := utils.SendSlack(atRequest.SlackBotToken, atRequest.SlackChannel, tfOutputs); err != nil {
			log.Printf("slack send error: %s\n", err)
		}
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
