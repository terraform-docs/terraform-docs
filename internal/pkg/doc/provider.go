package doc

type Provider struct {
	Name    string `json:"name"`
	Alias   string `json:"alias,omitempty"`
	Version string `json:"version,omitempty"`
}

type providersSortedByRequired []Provider

func (a providersSortedByRequired) Len() int {
	return len(a)
}

func (a providersSortedByRequired) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a providersSortedByRequired) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}