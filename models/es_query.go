// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    query, err := UnmarshalQuery(bytes)
//    bytes, err = query.Marshal()

package models

import "encoding/json"

func UnmarshalQuery(data []byte) (Query, error) {
	var r Query
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Query) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Query struct {
	Query     *QueryClass `json:"query"`
	Highlight *Highlight  `json:"highlight"`
}

type Highlight struct {
	Fields *Fields `json:"fields"`
}

type Fields struct {
	Title *Title `json:"title"`
}

type Title struct {
}

type QueryClass struct {
	MultiMatch *MultiMatch `json:"multi_match"`
}

type MultiMatch struct {
	Query     string   `json:"query"`
	Fuzziness int64    `json:"fuzziness"`
	Fields    []string `json:"fields"`
}
