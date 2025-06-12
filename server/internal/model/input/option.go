package input

type OptionVal struct {
	Value    int64        `json:"value"`
	Label    string       `json:"label"`
	Children []*OptionVal `json:"children"`
}
