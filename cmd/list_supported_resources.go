
package cmd



import (
	/* Importing source packages */
	"fmt"


	/* Importing custom packages */
	"../converters/google"
	/* Importing github packages */
	"github.com/spf13/cobra"

)



var listSupportedResourcesCmd = &cobra.Command{

	Use:   "list-supported-resources",

	Short: "List supported terraform resources.",

	RunE: func(c *cobra.Command, args []string) error {

		list := google.SupportedTerraformResources()



		for _, resource := range list {

			fmt.Println(resource)

		}



		return nil

	},

}
