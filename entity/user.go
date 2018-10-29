/* define user */
package entity

type User struct {
	Username string
	Password string
	Mail     string
	Phone    string
}

func (user *User) InitUser(name, psw, email, phone string) {
	user.Username = name
	user.Password = psw
	user.Mail = email
	user.Phone = phone
}

func (user *User) CopyUser(t User) {
	user.Username = t.Username
	user.Password = t.Password
	user.Mail = t.Mail
	user.Phone = t.Phone
}

func (user *User) getName() string {
	return user.Username
}

func (user *User) setName(name string) {
	user.Username = name
}

func (user *User) getPassword() string {
	return user.Password
}

func (user *User) setPassword(psw string) {
	user.Password = psw
}

func (user *User) getMail() string {
	return user.Mail
}

func (user *User) setMail(mail string) {
	user.Mail = mail
}

func (user *User) getPhone() string {
	return user.Phone
}

func (user *User) setPhone(phone string) {
	user.Phone = phone
}
