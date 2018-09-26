package models

/**
	the struct of visitors
 */
type Visitor struct {
  Visitorid_string string `json:"visitorid_string"`
  Items [] Items          `json:"items_array"`
}
