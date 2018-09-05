package headers

// represents header data for X-User-Namespace and X-User-Volume headers (encoded in base64)
//
//swagger:model
type UserHeaderData struct {
	// hosting-internal name
	// required: true
	ID string `json:"id"`
	// user-visible label for the object
	// required: true
	Label string `json:"label"`
	// one of: "owner", "read", "write", "read-delete", "none"
	// required: true
	Access string `json:"access"`
}
