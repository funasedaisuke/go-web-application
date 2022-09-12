package store

import 
func (r *Repository)RegisterUSer(ctx context.Context,db Execer, u * entity.User) error{
	u.Created = r.Clocker.Now()
	u.Modified = r.Clocker.Now()
	sql := `INSERT INTO user(
		name,password,role,created,modified
		) VALUES(?,?,?,?,?)`
		result,err := db.ExecContext(ctx,sql,u.Name,u.Password,u.Role,u.Created,u.Modified)
		if err != nil{
			var mysqlErr *mysql.MySQLError
			if errors.As(err,&mysqlErr) && mysqlErr.number == ErrCodeMySQLDuplicateEntry{
				
			}
		}
	
}