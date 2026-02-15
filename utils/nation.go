package utils

// GetNationName returns the display name for the given nation ID.
// Nation 1 is Quanato; any other value maps to Temoz.
func GetNationName(nation byte) string {
	switch nation {
	case 1:
		return "Quanato"
	default:
		return "Temoz"
	}
}
