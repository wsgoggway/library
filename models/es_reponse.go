package models

import "encoding/json"

func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Response struct {
	Took     int64  `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index   string   `json:"_index"`
	Type    string   `json:"_type"`
	ID      string   `json:"_id"`
	Score   float64  `json:"_score"`
	Ignored []string `json:"_ignored"`
	Source  Source   `json:"_source"`
}

type Source struct {
	Description string        `json:"description"`
	ID          int64         `json:"id"`
	Meta        []MetaElement `json:"meta"`
	Nickname    string        `json:"nickname"`
	Tags        []string      `json:"tags"`
	Title       string        `json:"title"`
}

type MetaElement struct {
	Meta MetaMeta `json:"meta"`
}

type MetaMeta struct {
	Author       string   `json:"author"`
	Bisac        string   `json:"bisac"`
	OriginalName string   `json:"original_name"`
	Pages        *int64   `json:"pages,omitempty"`
	Rating       string   `json:"rating"`
	Thumbnail    []string `json:"thumbnail"`
	Translator   string   `json:"translator"`
	Duration     *string  `json:"duration,omitempty"`
	Preview      *string  `json:"preview,omitempty"`
	Reader       *string  `json:"reader,omitempty"`
}

type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}
