package cmd

import (
	/* Importing source packages */
	"encoding/json"
	"os"

	/* Importing custom packages */
	"../tfgcv"
	
	/* Importing modules from github */
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)


var convertCmd = &cobra.Command{

	Use:   "convert <tfplan>",

	Short: "Convert resources in a Terraform plan to their Google CAI representation.",

	Long: `Convert (terraform-validator convert) will convert a Terraform plan file

into CAI (Cloud Asset Inventory) resources and output them as a JSON array.



Note:

  Only supported resources will be converted. Non supported resources are

  omitted from results.

  Run "terraform-validator list-supported-resources" to see all supported

  resources.



Example:

  terraform-validator convert ./example/terraform.tfplan --project my-project \

    --ancestry organization/my-org/folder/my-folder

`,

	PreRunE: func(c *cobra.Command, args []string) error {

		if len(args) != 1 {

			return errors.New("missing required argument <tfplan>")

		}

		return nil

	},

	RunE: func(c *cobra.Command, args []string) error {

		assets, err := tfgcv.ReadPlannedAssets(args[0], flags.convert.project, flags.convert.ancestry)

		if err != nil {

			if errors.Cause(err) == tfgcv.ErrParsingProviderProject {

				return errors.New("unable to parse provider project, please use --project flag")

			}

			return errors.Wrap(err, "converting tfplan to CAI assets")

		}



		if err := json.NewEncoder(os.Stdout).Encode(assets); err != nil {

			return errors.Wrap(err, "encoding json")

		}



		return nil

	},

}
