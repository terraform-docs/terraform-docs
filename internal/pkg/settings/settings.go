package settings

// Setting represents a setting
type Setting uint

// Settings represents all settings
type Settings uint

// Add adds a setting to settings
func (s *Settings) Add(setting Setting) {
	*s |= Settings(1 << setting)
}

// Has indicates if settings contain a setting
func (s *Settings) Has(setting Setting) bool {
	return *s&Settings(1<<setting) != 0
}

/*
package settings

// Setting represents a setting
type Setting uint

// Settings represents all settings
type Settings struct {
	s uint
}

func (s *Settings) add(setting Setting) {
	s.s |= 1 << setting
}

func (s *Settings) has(setting Setting) bool {
	return s.s&(1<<setting) != 0
}
*/
