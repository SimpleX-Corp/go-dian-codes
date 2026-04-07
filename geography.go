package codes

// Department represents Colombian department codes.
// Source: DIAN Caja de Herramientas - Departamentos-2.1.gc
type Department string

const (
	DeptAmazonas          Department = "91"
	DeptAntioquia         Department = "05"
	DeptArauca            Department = "81"
	DeptAtlantico         Department = "08"
	DeptBogota            Department = "11"
	DeptBolivar           Department = "13"
	DeptBoyaca            Department = "15"
	DeptCaldas            Department = "17"
	DeptCaqueta           Department = "18"
	DeptCasanare          Department = "85"
	DeptCauca             Department = "19"
	DeptCesar             Department = "20"
	DeptChoco             Department = "27"
	DeptCordoba           Department = "23"
	DeptCundinamarca      Department = "25"
	DeptGuainia           Department = "94"
	DeptGuaviare          Department = "95"
	DeptHuila             Department = "41"
	DeptLaGuajira         Department = "44"
	DeptMagdalena         Department = "47"
	DeptMeta              Department = "50"
	DeptNarino            Department = "52"
	DeptNorteSantander    Department = "54"
	DeptPutumayo          Department = "86"
	DeptQuindio           Department = "63"
	DeptRisaralda         Department = "66"
	DeptSanAndres         Department = "88"
	DeptSantander         Department = "68"
	DeptSucre             Department = "70"
	DeptTolima            Department = "73"
	DeptValleCauca        Department = "76"
	DeptVaupes            Department = "97"
	DeptVichada           Department = "99"
)

// Departments maps codes to department names.
var Departments = map[Department]string{
	DeptAmazonas:       "Amazonas",
	DeptAntioquia:      "Antioquia",
	DeptArauca:         "Arauca",
	DeptAtlantico:      "Atlántico",
	DeptBogota:         "Bogotá",
	DeptBolivar:        "Bolívar",
	DeptBoyaca:         "Boyacá",
	DeptCaldas:         "Caldas",
	DeptCaqueta:        "Caquetá",
	DeptCasanare:       "Casanare",
	DeptCauca:          "Cauca",
	DeptCesar:          "Cesar",
	DeptChoco:          "Chocó",
	DeptCordoba:        "Córdoba",
	DeptCundinamarca:   "Cundinamarca",
	DeptGuainia:        "Guainía",
	DeptGuaviare:       "Guaviare",
	DeptHuila:          "Huila",
	DeptLaGuajira:      "La Guajira",
	DeptMagdalena:      "Magdalena",
	DeptMeta:           "Meta",
	DeptNarino:         "Nariño",
	DeptNorteSantander: "Norte de Santander",
	DeptPutumayo:       "Putumayo",
	DeptQuindio:        "Quindío",
	DeptRisaralda:      "Risaralda",
	DeptSanAndres:      "San Andrés y Providencia",
	DeptSantander:      "Santander",
	DeptSucre:          "Sucre",
	DeptTolima:         "Tolima",
	DeptValleCauca:     "Valle del Cauca",
	DeptVaupes:         "Vaupés",
	DeptVichada:        "Vichada",
}

// IsValidDepartment checks if a department code is valid.
func IsValidDepartment(code string) bool {
	_, ok := Departments[Department(code)]
	return ok
}

// DepartmentName returns the name for a department code.
func DepartmentName(code string) string {
	return Departments[Department(code)]
}

// Municipality represents a Colombian municipality.
type Municipality struct {
	Code       string // 5-digit DANE code
	Name       string
	Department string // 2-digit department code
}

// GetMunicipality returns municipality info by code.
func GetMunicipality(code string) *Municipality {
	name, ok := municipalities[code]
	if !ok {
		return nil
	}
	dept := ""
	if len(code) >= 2 {
		dept = code[:2]
	}
	return &Municipality{
		Code:       code,
		Name:       name,
		Department: dept,
	}
}

// IsValidMunicipality checks if a municipality code is valid.
func IsValidMunicipality(code string) bool {
	_, ok := municipalities[code]
	return ok
}

// MunicipalityName returns the name for a municipality code.
func MunicipalityName(code string) string {
	return municipalities[code]
}

// GetDepartmentCode extracts the department code from a municipality code.
func GetDepartmentCode(municipalityCode string) string {
	if len(municipalityCode) >= 2 {
		return municipalityCode[:2]
	}
	return ""
}
