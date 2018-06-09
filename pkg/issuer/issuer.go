package issuer

import "context"

// Certificate is a base64 encoded PEM
type Certificate string

// Interface to issue/renew certificates
type Interface interface {
	Issue(context.Context, Certificate)
	Renew(context.Context, Certificate)
}
