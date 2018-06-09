package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
	"github.com/kunickiaj/vault-issuer/pkg/issuer"
)

type Vault struct {
	client *api.Client
}

func New(options ...func(*Vault)) issuer.Interface {
	vault := Vault{}

	for _, option := range options {
		option(&vault)
	}

	return &vault
}

func (v *Vault) Issue(ctx context.Context, cert issuer.Certificate) {
}

func (v *Vault) Renew(ctx context.Context, cert issuer.Certificate) {

}
