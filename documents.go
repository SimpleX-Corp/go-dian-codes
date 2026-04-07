package codes

// DocumentType represents electronic document types.
// Source: DIAN Caja de Herramientas - TipoDocumento-2.1.gc
type DocumentType string

const (
	DocInvoice         DocumentType = "01" // Factura electrónica de Venta
	DocExportInvoice   DocumentType = "02" // Factura electrónica de venta con propósito de exportación
	DocContingency     DocumentType = "03" // Factura de talonario o papel con numeración de contingencia
	DocContingencyDIAN DocumentType = "04" // Factura electrónica de Venta por Contingencia DIAN
	DocCreditNote      DocumentType = "91" // Nota Crédito
	DocDebitNote       DocumentType = "92" // Nota Débito
)

// DocumentTypes maps codes to descriptions.
var DocumentTypes = map[DocumentType]string{
	DocInvoice:         "Factura electrónica de Venta",
	DocExportInvoice:   "Factura electrónica de venta con propósito de exportación",
	DocContingency:     "Factura de talonario o papel con numeración de contingencia",
	DocContingencyDIAN: "Factura electrónica de Venta por Contingencia DIAN",
	DocCreditNote:      "Nota Crédito",
	DocDebitNote:       "Nota Débito",
}

// IsValidDocumentType checks if a document type code is valid.
func IsValidDocumentType(code string) bool {
	_, ok := DocumentTypes[DocumentType(code)]
	return ok
}

// DocumentTypeName returns the description for a document type.
func DocumentTypeName(code string) string {
	return DocumentTypes[DocumentType(code)]
}

// OperationType represents the operation type for invoices.
// Source: DIAN Caja de Herramientas - TipoOperacionF-2.1.gc
type OperationType string

const (
	OpStandard  OperationType = "10" // Estandar
	OpAIU       OperationType = "09" // AIU
	OpMandatos  OperationType = "11" // Mandatos
	OpTransport OperationType = "12" // Transporte
	OpCambiario OperationType = "13" // Cambiario
)

var OperationTypes = map[OperationType]string{
	OpStandard:  "Estandar",
	OpAIU:       "AIU",
	OpMandatos:  "Mandatos",
	OpTransport: "Transporte",
	OpCambiario: "Cambiario",
}

// IsValidOperationType checks if an operation type code is valid.
func IsValidOperationType(code string) bool {
	_, ok := OperationTypes[OperationType(code)]
	return ok
}

// CreditNoteCorrection represents credit note correction concepts.
// Source: DIAN Caja de Herramientas - ConceptoNotaCredito-2.1.gc
type CreditNoteCorrection string

const (
	CNCPartialReturn CreditNoteCorrection = "1" // Devolución parcial de los bienes y/o no aceptación parcial del servicio
	CNCFullCancel    CreditNoteCorrection = "2" // Anulación de factura electrónica
	CNCDiscount      CreditNoteCorrection = "3" // Rebaja o descuento parcial o total
	CNCPriceAdjust   CreditNoteCorrection = "4" // Ajuste de precio
	CNCOther         CreditNoteCorrection = "5" // Otros
)

var CreditNoteCorrections = map[CreditNoteCorrection]string{
	CNCPartialReturn: "Devolución parcial de los bienes y/o no aceptación parcial del servicio",
	CNCFullCancel:    "Anulación de factura electrónica",
	CNCDiscount:      "Rebaja o descuento parcial o total",
	CNCPriceAdjust:   "Ajuste de precio",
	CNCOther:         "Otros",
}

// IsValidCreditNoteCorrection checks if a credit note correction code is valid.
func IsValidCreditNoteCorrection(code string) bool {
	_, ok := CreditNoteCorrections[CreditNoteCorrection(code)]
	return ok
}

// DebitNoteCorrection represents debit note correction concepts.
// Source: DIAN Caja de Herramientas - ConceptoNotaDebito-2.1.gc
type DebitNoteCorrection string

const (
	DNCInterest    DebitNoteCorrection = "1" // Intereses
	DNCExpenses    DebitNoteCorrection = "2" // Gastos por cobrar
	DNCValueChange DebitNoteCorrection = "3" // Cambio del valor
	DNCOther       DebitNoteCorrection = "4" // Otro
)

var DebitNoteCorrections = map[DebitNoteCorrection]string{
	DNCInterest:    "Intereses",
	DNCExpenses:    "Gastos por cobrar",
	DNCValueChange: "Cambio del valor",
	DNCOther:       "Otro",
}

// IsValidDebitNoteCorrection checks if a debit note correction code is valid.
func IsValidDebitNoteCorrection(code string) bool {
	_, ok := DebitNoteCorrections[DebitNoteCorrection(code)]
	return ok
}
