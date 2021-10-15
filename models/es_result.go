package models

type Pair struct {
	ID     int64
	Weight int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Weight > p[j].Weight }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) GetIDs() (ids []int64) {
	ids = make([]int64, 0)
	for _, item := range p {
		ids = append(ids, item.ID)
	}

	return
}
