package smartid

import "fmt"

// IdentifierType for semantic identifier.
type IdentifierType string

// Available identifier types.
const (
	IdentifierTypePAS IdentifierType = "PAS"
	IdentifierTypeIDC IdentifierType = "IDC"
	IdentifierTypePNO IdentifierType = "PNO"
)

// CountryCode for semantic identifier.
type CountryCode string

// Available country codes.
const (
	CountryEE CountryCode = "EE"
	CountryLV CountryCode = "LV"
	CountryLT CountryCode = "LT"
	CountryKZ CountryCode = "KZ"
)

// NewSemanticIdentifier creates new semantic identifier as string.
func NewSemanticIdentifier(
	idType IdentifierType,
	country CountryCode,
	idCode string,
) string {
	return fmt.Sprintf("etsi/%s%s-%s", idType, country, idCode)
}
