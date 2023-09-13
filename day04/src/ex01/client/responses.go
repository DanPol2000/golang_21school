package main

type InlineResponse201 struct {
	Thanks string `json:"thanks,omitempty"`

	Change int32 `json:"change,omitempty"`
}

type InlineResponse400 struct {
	Error_ string `json:"error,omitempty"`
}
