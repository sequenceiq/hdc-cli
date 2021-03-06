package cmd

import (
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/aws"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/azure"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/gcp"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/openstack"
	_ "github.com/hortonworks/cb-cli/dataplane/cloud/yarn"
	cf "github.com/hortonworks/cb-cli/dataplane/config"
	fl "github.com/hortonworks/cb-cli/dataplane/flags"
	"github.com/hortonworks/cb-cli/dataplane/redbeams"
	"github.com/urfave/cli"
)

func init() {
	DataPlaneCommands = append(DataPlaneCommands, cli.Command{
		Name:  "redbeams",
		Usage: "manage relational databases and database servers",
		Subcommands: []cli.Command{
			{
				Name:  "dbserver",
				Usage: "manage redbeams database servers",
				Subcommands: []cli.Command{
					{
						Name:   "list",
						Usage:  "list all database servers",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentCrn).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.ListDatabaseServers,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentCrn).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "describe",
						Usage:       "describe a database server",
						Description: "To specify a database server, either provide its CRN, or both its environment CRN and name.",
						Before: func(c *cli.Context) error {
							err := cf.CheckConfigAndCommandFlagsWithoutWorkspace(c)
							if err != nil {
								return err
							}
							return cf.CheckResourceAddressingFlags(c)
						},
						Flags:  fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.GetDatabaseServer,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "create",
						Usage:  "create a managed database server",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlDatabaseServerCreationFile).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.CreateManagedDatabaseServer,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlDatabaseServerCreationFile).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "release",
						Usage:  "release a managed database server",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.ReleaseManagedDatabaseServer,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "register",
						Usage:  "register a database server",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlDatabaseServerRegistrationFile).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.RegisterDatabaseServer,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlDatabaseServerRegistrationFile).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "delete",
						Usage:  "delete a database server",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlForceOptional).AddAGlobalFlags().Build(),
						Action: redbeams.DeleteDatabaseServer,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlCrn, fl.FlForceOptional).AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
			{
				Name:  "db",
				Usage: "manage redbeams databases",
				Subcommands: []cli.Command{
					{
						Name:   "list",
						Usage:  "list all databases",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentCrn).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.ListDatabases,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlEnvironmentCrn).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "describe",
						Usage:       "describe a database",
						Description: "To specify a database, either provide its CRN, or both its environment CRN and name.",
						Before: func(c *cli.Context) error {
							err := cf.CheckConfigAndCommandFlagsWithoutWorkspace(c)
							if err != nil {
								return err
							}
							return cf.CheckResourceAddressingFlags(c)
						},
						Flags:  fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.GetDatabase,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "create",
						Usage:  "create a database",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlDatabaseCreationFile).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.CreateDatabase,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlDatabaseCreationFile).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:   "register",
						Usage:  "register a database",
						Before: cf.CheckConfigAndCommandFlagsWithoutWorkspace,
						Flags:  fl.NewFlagBuilder().AddFlags(fl.FlDatabaseRegistrationFile).AddAGlobalFlags().AddOutputFlag().Build(),
						Action: redbeams.RegisterDatabase,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddFlags(fl.FlDatabaseRegistrationFile).AddAGlobalFlags().AddOutputFlag().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
					{
						Name:        "delete",
						Usage:       "delete a database",
						Description: "To specify a database, either provide its CRN, or both its environment CRN and name.",
						Before: func(c *cli.Context) error {
							err := cf.CheckConfigAndCommandFlagsWithoutWorkspace(c)
							if err != nil {
								return err
							}
							return cf.CheckResourceAddressingFlags(c)
						},
						Flags:  fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().Build(),
						Action: redbeams.DeleteDatabase,
						BashComplete: func(c *cli.Context) {
							for _, f := range fl.NewFlagBuilder().AddResourceAddressingFlags().AddAGlobalFlags().Build() {
								fl.PrintFlagCompletion(f)
							}
						},
					},
				},
			},
		},
	})
}
