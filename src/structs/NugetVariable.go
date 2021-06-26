package structs

type NugetVariable struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
