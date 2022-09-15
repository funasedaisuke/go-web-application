package auth

import (
	_ "embed"
)

//go:embed cert/secret.pem
var rawPriveKey []byte


//go:embed cert/public.pem
var rawPubKey []byte