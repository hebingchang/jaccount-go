package jaccount

/**
* @Author: He Bingchang
* @Date: 2020/9/8
* @Description:
 */

type Profile struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Code     string `json:"code"`
	UserType string `json:"userType"`
	Organize struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"organize"`
	ClassNo string `json:"classNo"`
	Avatars struct {
	} `json:"avatars"`
	Birthday struct {
		BirthYear  string `json:"birthYear"`
		BirthMonth string `json:"birthMonth"`
		BirthDay   string `json:"birthDay"`
	} `json:"birthday"`
	Gender     string     `json:"gender"`
	Email      string     `json:"email"`
	TimeZone   int        `json:"timeZone"`
	Mobile     string     `json:"mobile"`
	Identities []Identity `json:"identities"`
	CardNo     string     `json:"cardNo"`
	CardType   string     `json:"cardType"`
	UnionID    string     `json:"unionId"`
}

type Identity struct {
	Kind      string `json:"kind"`
	IsDefault bool   `json:"isDefault"`
	Code      string `json:"code"`
	UserType  string `json:"userType"`
	Organize  struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"organize"`
	ExpireDate  string `json:"expireDate,omitempty"`
	CreateDate  int    `json:"createDate"`
	UpdateDate  int    `json:"updateDate"`
	MgtOrganize struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"mgtOrganize,omitempty"`
	Status  string `json:"status,omitempty"`
	ClassNo string `json:"classNo,omitempty"`
	Gjm     string `json:"gjm,omitempty"`
	Major   struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"major,omitempty"`
	AdmissionDate string `json:"admissionDate,omitempty"`
	TrainLevel    string `json:"trainLevel,omitempty"`
	GraduateDate  string `json:"graduateDate,omitempty"`
}

type profileResp struct {
	Errno    int       `json:"errno"`
	Error    string    `json:"error"`
	Total    int       `json:"total"`
	Entities []Profile `json:"entities"`
}
