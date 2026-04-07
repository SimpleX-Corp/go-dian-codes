package codes

// FiscalIDType represents fiscal identification type codes.
// Source: DIAN Caja de Herramientas - TipoIdFiscal-2.1.gc
type FiscalIDType string

const (
	IDRegistroCivil      FiscalIDType = "11" // Registro civil
	IDTarjetaIdentidad   FiscalIDType = "12" // Tarjeta de identidad
	IDCedula             FiscalIDType = "13" // Cédula de ciudadanía
	IDTarjetaExtranjeria FiscalIDType = "21" // Tarjeta de extranjería
	IDCedulaExtranjeria  FiscalIDType = "22" // Cédula de extranjería
	IDNIT                FiscalIDType = "31" // NIT
	IDPasaporte          FiscalIDType = "41" // Pasaporte
	IDDocExtranjero      FiscalIDType = "42" // Documento de identificación extranjero
	IDNITOtroPais        FiscalIDType = "50" // NIT de otro país
	IDNUIP               FiscalIDType = "91" // NUIP
)

var FiscalIDTypes = map[FiscalIDType]string{
	IDRegistroCivil:      "Registro civil",
	IDTarjetaIdentidad:   "Tarjeta de identidad",
	IDCedula:             "Cédula de ciudadanía",
	IDTarjetaExtranjeria: "Tarjeta de extranjería",
	IDCedulaExtranjeria:  "Cédula de extranjería",
	IDNIT:                "NIT",
	IDPasaporte:          "Pasaporte",
	IDDocExtranjero:      "Documento de identificación extranjero",
	IDNITOtroPais:        "NIT de otro país",
	IDNUIP:               "NUIP",
}

// IsValidFiscalIDType checks if a fiscal ID type code is valid.
func IsValidFiscalIDType(code string) bool {
	_, ok := FiscalIDTypes[FiscalIDType(code)]
	return ok
}

// FiscalIDTypeName returns the description for a fiscal ID type.
func FiscalIDTypeName(code string) string {
	return FiscalIDTypes[FiscalIDType(code)]
}

// OrganizationType represents organization type codes.
// Source: DIAN Caja de Herramientas - TipoOrganizacion-2.1.gc
type OrganizationType string

const (
	OrgPersonaJuridica OrganizationType = "1" // Persona Jurídica y asimiladas
	OrgPersonaNatural  OrganizationType = "2" // Persona Natural y asimiladas
)

var OrganizationTypes = map[OrganizationType]string{
	OrgPersonaJuridica: "Persona Jurídica y asimiladas",
	OrgPersonaNatural:  "Persona Natural y asimiladas",
}

// IsValidOrganizationType checks if an organization type code is valid.
func IsValidOrganizationType(code string) bool {
	_, ok := OrganizationTypes[OrganizationType(code)]
	return ok
}
