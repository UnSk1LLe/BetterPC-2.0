package updateInputs

type UpdateInput interface {
	Validate() error
	Decompose() (map[string]interface{}, error)
}
