package gamegrp

// Cup represents the cup data that can be returned in a request.
type Cup struct {
	Account string `json:"account"`
	Outs    int    `json:"outs"`
}
