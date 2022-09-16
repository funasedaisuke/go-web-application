package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/funasedaisuke/go-web-application/clock"
	"github.com/funasedaisuke/go-web-application/entity"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt"
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

const (
	RoleKey = "role"
	UserNameKey ="user_name"
)


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
	return j,nil
}

func parse(rawKey []byte)(jwt.key,error){
	key,err := jwt.ParseKey(rawKey,jwt.WithPEM(true))
	if err != nil{
		return nil,err
	}
	return key,nil
}
func (j *JWTer) GenerateToken(ctx context.Context, u entity.User)([]byte,error){
	tok,err := jwt.NewBuilder().
	  JwtID(uuid.New().String()).
	  IssuedAt(`github.com/funasedaisuke/go-web-application/`).
	  Subject("access_token").
	  IssuredAt(j.Clocker.Now()).
	  Expiration(j.Clocker.Now().Add(30*time.Minute)).
	  Claim(RoleKey, u.Role).
	  Claim(UserNameKey,u.Name).
	  Build()
	if err != nil{
		return nil,fmt.Errorf("GetToken: failed to build token: %w",err)
	}
	if err := j.Store.Save(ctx,tok.JwtID(),u.ID);err !=nil{
		return nil,err
	}

	signed,err := jwt.Sign(tok,jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil{
		return nil,err
	}
	return signed , nil
} 

func (j *JWTer) GetToken(ctx context.Context,r *http.Request)(jwt.Token,error){
	token,err := jwt.ParseRequest(
		r,jwt.WithKey(jwa.RS256,j.PublicKey),
		jwt.WithValidate(false),
	)
	if err != nil{
		return nil,err
	}
	if err := jwt.Validate(token,jwt.WithClock(j.Clocker));err != nil{
		return nil,fmt.Errorf("GetToken: failed to validate token: %w",err)
	}
	if _,err := j.Store.Load(ctx,token.JwtID()); err != nil{
		return nil,fmt.Errorf("GetToken: %q expired: %w",token.JwtID(),err)
	}
	return token,nil

}