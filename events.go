package codes

// Environment represents DIAN environment types.
// Source: DIAN Caja de Herramientas - TipoAmbiente-2.1.gc
type Environment string

const (
	EnvProduction Environment = "1" // Producción
	EnvTest       Environment = "2" // Pruebas
)

var Environments = map[Environment]string{
	EnvProduction: "Producción",
	EnvTest:       "Pruebas",
}

// IsValidEnvironment checks if an environment code is valid.
func IsValidEnvironment(code string) bool {
	_, ok := Environments[Environment(code)]
	return ok
}

// DocumentEvent represents document event codes.
// Source: DIAN Caja de Herramientas - EventoDocumento-2.1.gc
type DocumentEvent string

const (
	EventAuthorized         DocumentEvent = "02"  // Uso Autorizado por la DIAN
	EventCorrectionRequest  DocumentEvent = "030" // Solicitación de Corrección en Documento
	EventRejected           DocumentEvent = "031" // Rechazo de Documento
	EventGoodsReceived      DocumentEvent = "032" // Recibimiento de los Bienes
	EventAccepted           DocumentEvent = "033" // Aceptación de Documento
	EventNegotiationOffered DocumentEvent = "040" // Factura Ofrecida para Negociación como Título Valor
	EventNegotiated         DocumentEvent = "041" // Factura Negociada como Título Valor
)

var DocumentEvents = map[DocumentEvent]string{
	EventAuthorized:         "Uso Autorizado por la DIAN",
	EventCorrectionRequest:  "Solicitación de Corrección en Documento",
	EventRejected:           "Rechazo de Documento",
	EventGoodsReceived:      "Recibimiento de los Bienes",
	EventAccepted:           "Aceptación de Documento",
	EventNegotiationOffered: "Factura Ofrecida para Negociación como Título Valor",
	EventNegotiated:         "Factura Negociada como Título Valor",
}

// IsValidDocumentEvent checks if a document event code is valid.
func IsValidDocumentEvent(code string) bool {
	_, ok := DocumentEvents[DocumentEvent(code)]
	return ok
}

// ClaimConcept represents claim/rejection concepts.
// Source: DIAN Caja de Herramientas - Concepto de Reclamo.gc
type ClaimConcept string

const (
	ClaimInconsistencies    ClaimConcept = "01" // Documento con inconsistencias
	ClaimNotDeliveredFull   ClaimConcept = "02" // Mercancía no entregada totalmente
	ClaimNotDeliveredPartial ClaimConcept = "03" // Mercancía no entregada parcialmente
	ClaimServiceNotProvided ClaimConcept = "04" // Servicio no prestado
)

var ClaimConcepts = map[ClaimConcept]string{
	ClaimInconsistencies:     "Documento con inconsistencias",
	ClaimNotDeliveredFull:    "Mercancía no entregada totalmente",
	ClaimNotDeliveredPartial: "Mercancía no entregada parcialmente",
	ClaimServiceNotProvided:  "Servicio no prestado",
}

// IsValidClaimConcept checks if a claim concept code is valid.
func IsValidClaimConcept(code string) bool {
	_, ok := ClaimConcepts[ClaimConcept(code)]
	return ok
}
