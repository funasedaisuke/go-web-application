package auth

import (
	"context"
	"fmt"

	"github.com/funasedaisuke/go-web-application/clock"
	"github.com/funasedaisuke/go-web-application/entity"
	"github.com/lestrrat-go/jwx/jwt"
)

type JWTer struct {
	PrivateKey,PublicKey jwt.Key
	Store   Store
	Clocker clock.Clocker

}




type Store interface{
	Save(ctx context.Context, key string,userID entity.UserID) error
	Load(ctx context.Context,key string)(entity.UserID,error)
}

func NewJWTer(s Store)(*JWTer,error){
	j := &JWTer{Store:s}
	privatekey,err := parse(rawPrivateKey)
	if err != nil{
		return nil, fmt.Errorf("failed in NewJWTer:private key: %w",err)
	}
	pubkey,err :=parse(rawPubKey)
	if err != nil{
		return nil,fmt.Errorf("failed in NewJWTer:public key: %w",err)
	}
	j.PrivateKey = privatekey
	j.PublicKey = pubkey
	j.Clocker = clock.RealClocker{}
}

func parse(rawKey []byte)(jwt.key,error){
	key,err := jwt.ParseKey(rawKey,jwt.WithPEM(true))
	if err != nil{
		return nil,err
	}
	return key,nil
}