package get

import (
	"context"
	cloudclient "github.com/otterize/otterize-cli/src/pkg/cloudclient/restapi"
	"github.com/otterize/otterize-cli/src/pkg/cloudclient/restapi/cloudapi"
	"github.com/otterize/otterize-cli/src/pkg/config"
	"github.com/otterize/otterize-cli/src/pkg/output"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var GetOrganizationCmd = &cobra.Command{
	Use:          "get <organization-id>",
	Aliases:      []string{"org"},
	Short:        "Get details for an organization",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(_ *cobra.Command, args []string) error {
		ctxTimeout, cancel := context.WithTimeout(context.Background(), config.DefaultTimeout)
		defer cancel()
		c, err := cloudclient.NewClient(ctxTimeout)
		if err != nil {
			return err
		}

		id := args[0]

		r, err := c.OrganizationQueryWithResponse(ctxTimeout, id)
		if err != nil {
			return err
		}

		output.FormatOrganizations([]cloudapi.Organization{lo.FromPtr(r.JSON200)})
		return nil
	},
}
