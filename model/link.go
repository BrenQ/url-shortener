package model

// Model Url

type Url struct {
	Original string `json:"Original,omitempty"`
	Short string `json:"Short,omitempty"`
}

// Struct aux for store request
type UrlRequest struct {
	Url string `json:"url" binding:"required"`
}
// Struct with code request
type CodeRequest struct {
	Code string `uri:"code" binding:"required"`
}
