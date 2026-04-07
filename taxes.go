package codes

// TaxType represents tax type codes.
// Source: DIAN Caja de Herramientas - TipoImpuesto-2.1.gc
type TaxType string

const (
	TaxIVA            TaxType = "01" // IVA
	TaxIC             TaxType = "02" // IC (Impuesto al Consumo)
	TaxICA            TaxType = "03" // ICA
	TaxINC            TaxType = "04" // INC (Impuesto Nacional al Consumo)
	TaxReteIVA        TaxType = "05" // ReteIVA
	TaxReteFuente     TaxType = "06" // ReteFuente
	TaxReteICA        TaxType = "07" // ReteICA
	TaxReteCREE       TaxType = "08" // ReteCREE
	TaxFtoHorticultura TaxType = "20" // FtoHorticultura
	TaxTimbre         TaxType = "21" // Timbre
	TaxBolsas         TaxType = "22" // Bolsas
	TaxINCarbono      TaxType = "23" // INCarbono
	TaxINCombustibles TaxType = "24" // INCombustibles
	TaxSobretasaComb  TaxType = "25" // Sobretasa Combustibles
	TaxSordicom       TaxType = "26" // Sordicom
	TaxOther          TaxType = "ZZ" // Nombre de la figura tributaria
)

var TaxTypes = map[TaxType]string{
	TaxIVA:            "IVA",
	TaxIC:             "IC",
	TaxICA:            "ICA",
	TaxINC:            "INC",
	TaxReteIVA:        "ReteIVA",
	TaxReteFuente:     "ReteFuente",
	TaxReteICA:        "ReteICA",
	TaxReteCREE:       "ReteCREE",
	TaxFtoHorticultura: "FtoHorticultura",
	TaxTimbre:         "Timbre",
	TaxBolsas:         "Bolsas",
	TaxINCarbono:      "INCarbono",
	TaxINCombustibles: "INCombustibles",
	TaxSobretasaComb:  "Sobretasa Combustibles",
	TaxSordicom:       "Sordicom",
	TaxOther:          "Nombre de la figura tributaria",
}

// IsValidTaxType checks if a tax type code is valid.
func IsValidTaxType(code string) bool {
	_, ok := TaxTypes[TaxType(code)]
	return ok
}

// TaxTypeName returns the description for a tax type.
func TaxTypeName(code string) string {
	return TaxTypes[TaxType(code)]
}

// IVARate represents IVA tax rates.
// Source: DIAN Caja de Herramientas - TarifaImpuestoIVA-2.1.gc
type IVARate string

const (
	IVAExempt   IVARate = "0.00"  // Exento
	IVA5        IVARate = "5.00"  // Bienes / Servicios al 5%
	IVA16       IVARate = "16.00" // Contratos firmados con el estado antes de ley 1819
	IVA19       IVARate = "19.00" // Tarifa general
)

var IVARates = map[IVARate]string{
	IVAExempt: "Exento",
	IVA5:      "Bienes / Servicios al 5%",
	IVA16:     "Contratos firmados con el estado antes de ley 1819",
	IVA19:     "Tarifa general",
}

// IsValidIVARate checks if an IVA rate is valid.
func IsValidIVARate(rate string) bool {
	_, ok := IVARates[IVARate(rate)]
	return ok
}

// TaxResponsibility represents tax responsibility codes.
// Source: DIAN Caja de Herramientas - TipoResponsabilidad-2.1.gc
type TaxResponsibility string

const (
	RespGranContribuyente TaxResponsibility = "O-13" // Gran contribuyente
	RespAutorretenedor    TaxResponsibility = "O-15" // Autorretenedor
	RespAgenteReteIVA     TaxResponsibility = "O-23" // Agente de retención IVA
	RespRegimenSimple     TaxResponsibility = "O-47" // Régimen simple de tributación
	RespNoAplica          TaxResponsibility = "ZZ"   // No aplica
)

var TaxResponsibilities = map[TaxResponsibility]string{
	RespGranContribuyente: "Gran contribuyente",
	RespAutorretenedor:    "Autorretenedor",
	RespAgenteReteIVA:     "Agente de retención IVA",
	RespRegimenSimple:     "Régimen simple de tributación",
	RespNoAplica:          "No aplica",
}

// IsValidTaxResponsibility checks if a tax responsibility code is valid.
func IsValidTaxResponsibility(code string) bool {
	_, ok := TaxResponsibilities[TaxResponsibility(code)]
	return ok
}

// INCRate represents INC (Impuesto Nacional al Consumo) tax rates.
// Source: DIAN Caja de Herramientas - TarifaImpuestoINC-2.1.gc
type INCRate string

const (
	INC2  INCRate = "2.00"  // Tarifa especial
	INC4  INCRate = "4.00"  // Tarifa especial
	INC8  INCRate = "8.00"  // Tarifa general
	INC16 INCRate = "16.00" // Tarifa especial
)

var INCRates = map[INCRate]string{
	INC2:  "Tarifa especial 2%",
	INC4:  "Tarifa especial 4%",
	INC8:  "Tarifa general 8%",
	INC16: "Tarifa especial 16%",
}

// IsValidINCRate checks if an INC rate is valid.
func IsValidINCRate(rate string) bool {
	_, ok := INCRates[INCRate(rate)]
	return ok
}

// WithholdingRate represents common withholding tax rates (ReteFuente).
// Source: DIAN Caja de Herramientas - TarifaImpuestoReteFuente-2.1.gc
// Note: This is a simplified list. Full list has 40+ specific rates by concept.
var WithholdingRates = map[string]string{
	"0.10":  "Compras de combustibles derivados del petróleo",
	"0.50":  "Compras de café pergamino o cereza",
	"1.00":  "Enajenación activos fijos / Compras vehículos / Transporte carga",
	"1.50":  "Compras con tarjeta / Productos agrícolas sin procesar",
	"2.00":  "Vigilancia y aseo / IPS / Construcción",
	"2.50":  "Compras generales (declarantes) / Bienes raíces",
	"3.00":  "Juegos de suerte y azar",
	"3.50":  "Compras generales (no declarantes) / Software / Hoteles",
	"4.00":  "Servicios generales (declarantes) / Arrendamiento muebles",
	"6.00":  "Servicios generales (no declarantes)",
	"7.00":  "Intereses o rendimientos financieros",
	"10.00": "Honorarios y comisiones (no declarantes)",
	"11.00": "Honorarios y comisiones (personas jurídicas)",
	"20.00": "Loterías, rifas, apuestas y similares",
}

// IsValidWithholdingRate checks if a withholding rate is in the common list.
func IsValidWithholdingRate(rate string) bool {
	_, ok := WithholdingRates[rate]
	return ok
}
