package fixture

import "github.com/funasedaisuke/go-web-application/entity"

func User(u *entity.User) *entity.User{
	result := &entity.User{
	ID: entity.UserID(rand.Int()),
	Name:   "daisuke"+strconv.Itoa(rand.Int())[:5],
	Password:"password",
	Role:"admin",
	Created:time.Now(),
	Modified:time.Now(),
}
if u == nil {
	return result
}
if u.ID != 0{
	result.ID = u.ID
}

if u.Name != nil{
	result.Name = u.Name 
} 

if u.Password != nil{
	result.Password = u.Password
}

if u.Role != ""{
	result.Role = u.Role
}

if !u.Created.IsZero(){
	result.Created = u.Created
}
if !u.Modified.IsZero(){
result.Modified = u.Modified
}
return result