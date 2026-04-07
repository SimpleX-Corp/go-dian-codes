package codes

// DeliveryTerm represents Incoterms delivery condition codes.
// Source: DIAN Caja de Herramientas - TipoEntrega-2.1.gc
type DeliveryTerm string

const (
	IncotermCFR DeliveryTerm = "CFR" // Costo y flete
	IncotermCIF DeliveryTerm = "CIF" // Costo, flete y seguro
	IncotermCIP DeliveryTerm = "CIP" // Transporte y Seguro Pagados hasta
	IncotermCPT DeliveryTerm = "CPT" // Transporte Pagado Hasta
	IncotermDAP DeliveryTerm = "DAP" // Entregado en un Lugar
	IncotermDAT DeliveryTerm = "DAT" // Entregado en Terminal
	IncotermDDP DeliveryTerm = "DDP" // Entregado con Pago de Derechos
	IncotermEXW DeliveryTerm = "EXW" // En Fábrica
	IncotermFAS DeliveryTerm = "FAS" // Franco al costado del buque
	IncotermFCA DeliveryTerm = "FCA" // Franco transportista
	IncotermFOB DeliveryTerm = "FOB" // Franco a bordo
)

var DeliveryTerms = map[DeliveryTerm]string{
	IncotermCFR: "Costo y flete",
	IncotermCIF: "Costo, flete y seguro",
	IncotermCIP: "Transporte y Seguro Pagados hasta",
	IncotermCPT: "Transporte Pagado Hasta",
	IncotermDAP: "Entregado en un Lugar",
	IncotermDAT: "Entregado en Terminal",
	IncotermDDP: "Entregado con Pago de Derechos",
	IncotermEXW: "En Fábrica",
	IncotermFAS: "Franco al costado del buque",
	IncotermFCA: "Franco transportista",
	IncotermFOB: "Franco a bordo",
}

// IsValidDeliveryTerm checks if a delivery term (Incoterm) is valid.
func IsValidDeliveryTerm(code string) bool {
	_, ok := DeliveryTerms[DeliveryTerm(code)]
	return ok
}
