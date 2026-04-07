package codes

// UnitCode represents units of measure codes.
// Source: DIAN Caja de Herramientas - UnidadesMedida-2.1.gc (UN/ECE Rec 20)
type UnitCode string

// Common unit codes used in DIAN electronic invoicing.
const (
	UnitEach      UnitCode = "EA"  // Each (Unidad)
	UnitPiece     UnitCode = "C62" // Piece (Pieza)
	UnitPack      UnitCode = "PK"  // Pack (Paquete)
	UnitBox       UnitCode = "BX"  // Box (Caja)
	UnitKilogram  UnitCode = "KGM" // Kilogram (Kilogramo)
	UnitGram      UnitCode = "GRM" // Gram (Gramo)
	UnitLiter     UnitCode = "LTR" // Liter (Litro)
	UnitMeter     UnitCode = "MTR" // Meter (Metro)
	UnitSqMeter   UnitCode = "MTK" // Square meter (Metro cuadrado)
	UnitCuMeter   UnitCode = "MTQ" // Cubic meter (Metro cúbico)
	UnitHour      UnitCode = "HUR" // Hour (Hora)
	UnitDay       UnitCode = "DAY" // Day (Día)
	UnitMonth     UnitCode = "MON" // Month (Mes)
	UnitYear      UnitCode = "ANN" // Year (Año)
	UnitService   UnitCode = "E48" // Service unit (Unidad de servicio)
	UnitActivity  UnitCode = "ACT" // Activity (Actividad)
	UnitDozen     UnitCode = "DZN" // Dozen (Docena)
	UnitGallon    UnitCode = "GLL" // Gallon (Galón)
	UnitPound     UnitCode = "LBR" // Pound (Libra)
)

// CommonUnits contains the most common units used in Colombian invoicing.
var CommonUnits = map[UnitCode]string{
	UnitEach:     "Unidad",
	UnitPiece:    "Pieza",
	UnitPack:     "Paquete",
	UnitBox:      "Caja",
	UnitKilogram: "Kilogramo",
	UnitGram:     "Gramo",
	UnitLiter:    "Litro",
	UnitMeter:    "Metro",
	UnitSqMeter:  "Metro cuadrado",
	UnitCuMeter:  "Metro cúbico",
	UnitHour:     "Hora",
	UnitDay:      "Día",
	UnitMonth:    "Mes",
	UnitYear:     "Año",
	UnitService:  "Unidad de servicio",
	UnitActivity: "Actividad",
	UnitDozen:    "Docena",
	UnitGallon:   "Galón",
	UnitPound:    "Libra",
}

// IsValidCommonUnit checks if a unit code is in the common units list.
// Note: DIAN accepts many more units from UN/ECE Rec 20.
func IsValidCommonUnit(code string) bool {
	_, ok := CommonUnits[UnitCode(code)]
	return ok
}

// Common currency codes used in Colombian invoicing.
// Full list available via IsValidCurrencyCode() in currencies_gen.go
const (
	CurrencyCOP = "COP" // Colombian Peso
	CurrencyUSD = "USD" // US Dollar
	CurrencyEUR = "EUR" // Euro
)
