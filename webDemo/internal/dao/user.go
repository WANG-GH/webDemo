package dao

import "webDemo/internal/model"

func (d *Dao) CreateUser(email string,name string, passwd string, privilege uint32) error {
	user := model.User{
		UserName:      name,
		Password:  passwd,
		Privilege: privilege,
		Email: email,
	}
	return user.CreateUser(d.engine)
}

func (d *Dao) UpdateUser(name string, passwd string, id uint32, privilege uint32) (model.User,error) {
	user := model.User{
		UserName:      name,
		Password:  passwd,
		ID:        id,
		Privilege: privilege,
	}
	return user.Update(d.engine)
}

func (d *Dao)CountUser(name string) (int, error){
	user := model.User{
		UserName:      name,
	}
	return user.Count(d.engine)
}

func (d *Dao) GetUserPasswd(name string) (string, error){
	user := model.User{
		UserName:      name,
	}
	return user.GetPasswd(d.engine)
}

func (d *Dao) GetStatus(name string) (model.User, error){
	user := model.User{
		UserName:      name,
	}
	return user.GetStatus(d.engine)
}
