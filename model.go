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

type Course struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Teacher struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Organize struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct {
	Kind   string `json:"kind"`
	Week   int    `json:"week"`
	Day    int    `json:"day"`
	Period int    `json:"period"`
	Start  int    `json:"start"`
	Finish int    `json:"finish"`
}

type Classroom struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Class struct {
	Schedule  Schedule  `json:"schedule"`
	Classroom Classroom `json:"classroom"`
}

type Lesson struct {
	Name     string    `json:"name"`
	Kind     string    `json:"kind"`
	Bsid     string    `json:"bsid"`
	Code     string    `json:"code"`
	Course   Course    `json:"course"`
	Teachers []Teacher `json:"teachers"`
	Organize Organize  `json:"organize"`
	Hours    int       `json:"hours"`
	Credits  int       `json:"credits"`
	Classes  []Class   `json:"classes"`
}

type lessonsResp struct {
	Errno    int      `json:"errno"`
	Error    string   `json:"error"`
	Total    int      `json:"total"`
	Entities []Lesson `json:"entities"`
}

type profileResp struct {
	Errno    int       `json:"errno"`
	Error    string    `json:"error"`
	Total    int       `json:"total"`
	Entities []Profile `json:"entities"`
}
