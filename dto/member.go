package dto

type MemberDto struct {
	Id     int    `json:"id"`
	Passwd string `json:"passwd"`
	Email  string `json:"email"`
}

func (member *MemberDto) MemberDto(id int, passwd string, email string) {
	member.Id = id
	member.Passwd = passwd
	member.Email = email
}

func (member *MemberDto) GetId() int {
	return member.Id
}

func (member *MemberDto) SetId(id int) {
	member.Id = id
}

func (member *MemberDto) GetPasswd() string {
	return member.Passwd
}

func (member *MemberDto) SetPasswd(passwd string) {
	member.Passwd = passwd
}

func (member *MemberDto) GetEmail() string {
	return member.Email
}

func (member *MemberDto) SetEmail(email string) {
	member.Email = email
}
