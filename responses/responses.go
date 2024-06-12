package responses

type (
	BasicI interface {
		GetCode() int
		GetMsg() string
	}

	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)

func (c *Basic) GetCode() int {
	return c.Code
}

func (c *Basic) GetMsg() string {
	return c.Msg
}
