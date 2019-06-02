package types

// MetaObject is object metadata as seen by the object and metadata stores.
type MetaObject struct {
	Oid      string `json:"oid"`
	Size     int64  `json:"size"`
	Existing bool
}

type ObjectError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
