package settings

// Flag represents a flag
type Flag uint

// Setting represents a setting
type Setting string

// Settings represents all settings
type Settings struct {
	Flags  Flag
	Values map[Setting]string
}

// Add adds a flag to settings
func (s *Settings) Add(setting Flag) {
	s.Flags |= (1 << setting)
}

// Add adds a setting to settings
func (s *Settings) Set(setting Setting, value string) {
	s.Values[setting] = value
}

// Has indicates if settings contain a flag
func (s *Settings) Has(setting Flag) bool {
	return s.Flags&(1<<setting) != 0
}

// Get fetches the value of setting
func (s *Settings) Get(setting Setting) string {
	return s.Values[setting]
}
