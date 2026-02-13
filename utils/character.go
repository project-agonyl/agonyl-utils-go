package utils

// GetClassName returns the display name for the given character class ID.
// Class 1 is Holy Knight, 2 is Mage, 3 is Archer; any other value maps to Warrior.
func GetClassName(class byte) string {
	switch class {
	case 1:
		return "Holy Knight"
	case 2:
		return "Mage"
	case 3:
		return "Archer"
	default:
		return "Warrior"
	}
}
