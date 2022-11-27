package forms

type Authorize struct {
	Login    string `schema:"login,required"`
	Password string `schema:"password,required"`
}
