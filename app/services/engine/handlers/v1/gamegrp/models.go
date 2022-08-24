package gamegrp

// Cup represents the cup response.
type Cup struct {
	Account string `json:"account"`
	Outs    int    `json:"outs"`
}

// Claim represents the claim response.
type Claim struct {
	Account string `json:"account"`
	Number  int    `json:"number"`
	Suite   int    `json:"suite"`
}
