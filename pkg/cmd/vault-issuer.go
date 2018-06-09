package cmd

import (
	"context"
	"io/ioutil"

	"github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// Version of the binary including git commit hash representing the source
	Version string
)

// NewCommand creates the root command of the tool
func NewCommand(name string) *cobra.Command {
	c := &cobra.Command{
		Use:     name,
		Short:   "",
		Long:    ``,
		Version: Version,
		RunE:    run,
	}

	c.InitDefaultVersionFlag()

	c.Flags().String("vault-url", "", "Vault URL")
	c.Flags().String("auth-role", "", "Vault authentication role")
	c.Flags().String("pki-path", "/pki", "path to the PKI mount point")
	c.Flags().String("pki-role", "", "signing role name")
	c.Flags().String("cn", "", "common name in issued certificate")
	// TODO: add flag for TTL, etc

	return c
}

func run(c *cobra.Command, args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	url, err := c.Flags().GetString("vault-url")
	cfg := api.DefaultConfig()
	cfg.Address = url

	client, err := api.NewClient(cfg)
	log.WithError(err).Error("failed to create new Vault client")

	b, err := readKubernetesToken()
	if err != nil {
		log.WithError(err).Error("failed to read service account token")
	}

	jwt := string(b)
	role, err := c.Flags().GetString("auth-role")
	if err != nil {
		log.WithError(err).Error()
	}

	loginPath := "auth/kubernetes/login"
	login := map[string]interface{}{
		"role": role,
		"jwt":  jwt,
	}

	secret, err := client.Logical().Write(loginPath, login)
	client.SetToken(secret.Auth.ClientToken)

	select {
	case <-ctx.Done():
		return nil
	}
}

func readKubernetesToken() ([]byte, error) {
	return ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
}
