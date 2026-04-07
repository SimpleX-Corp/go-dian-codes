package codes

// PaymentMethod represents payment method codes.
// Source: DIAN Caja de Herramientas - FormasPago-2.1.gc
type PaymentMethod string

const (
	PaymentCash   PaymentMethod = "1" // Contado
	PaymentCredit PaymentMethod = "2" // Crédito
)

var PaymentMethods = map[PaymentMethod]string{
	PaymentCash:   "Contado",
	PaymentCredit: "Crédito",
}

// IsValidPaymentMethod checks if a payment method code is valid.
func IsValidPaymentMethod(code string) bool {
	_, ok := PaymentMethods[PaymentMethod(code)]
	return ok
}

// PaymentMeans represents payment means codes.
// Source: DIAN Caja de Herramientas - MediosPago-2.1.gc
type PaymentMeans string

const (
	MeansNotDefined       PaymentMeans = "1"   // Instrumento no definido
	MeansACHCredit        PaymentMeans = "2"   // Crédito ACH
	MeansACHDebit         PaymentMeans = "3"   // Débito ACH
	MeansCash             PaymentMeans = "10"  // Efectivo
	MeansCheck            PaymentMeans = "20"  // Cheque
	MeansCreditTransfer   PaymentMeans = "30"  // Transferencia Crédito
	MeansDebitTransfer    PaymentMeans = "31"  // Transferencia Débito
	MeansBankDeposit      PaymentMeans = "42"  // Consignación bancaria
	MeansBankCreditTrans  PaymentMeans = "45"  // Transferencia Crédito Bancario
	MeansBankDebitTrans   PaymentMeans = "47"  // Transferencia Débito Bancaria
	MeansCreditCard       PaymentMeans = "48"  // Tarjeta Crédito
	MeansDebitCard        PaymentMeans = "49"  // Tarjeta Débito
	MeansPromissoryNote   PaymentMeans = "60"  // Nota promisoria
	MeansMutualAgreement  PaymentMeans = "ZZZ" // Acuerdo mutuo
)

// PaymentMeansFull contains all payment means from DIAN.
var PaymentMeansFull = map[PaymentMeans]string{
	"1":   "Instrumento no definido",
	"2":   "Crédito ACH",
	"3":   "Débito ACH",
	"4":   "Reversión débito de demanda ACH",
	"5":   "Reversión crédito de demanda ACH",
	"6":   "Crédito de demanda ACH",
	"7":   "Débito de demanda ACH",
	"8":   "Mantener",
	"9":   "Clearing Nacional o Regional",
	"10":  "Efectivo",
	"11":  "Reversión Crédito Ahorro",
	"12":  "Reversión Débito Ahorro",
	"13":  "Crédito Ahorro",
	"14":  "Débito Ahorro",
	"15":  "Bookentry Crédito",
	"16":  "Bookentry Débito",
	"17":  "Concentración de la demanda en efectivo /Desembolso Crédito (CCD)",
	"18":  "Concentración de la demanda en efectivo / Desembolso (CCD) débito",
	"19":  "Crédito Pago negocio corporativo (CTP)",
	"20":  "Cheque",
	"21":  "Proyecto bancario",
	"22":  "Proyecto bancario certificado",
	"23":  "Cheque bancario",
	"24":  "Nota cambiaria esperando aceptación",
	"25":  "Cheque certificado",
	"26":  "Cheque Local",
	"27":  "Débito Pago Negocio Corporativo (CTP)",
	"28":  "Crédito Negocio Intercambio Corporativo (CTX)",
	"29":  "Débito Negocio Intercambio Corporativo (CTX)",
	"30":  "Transferencia Crédito",
	"31":  "Transferencia Débito",
	"32":  "Concentración Efectivo / Desembolso Crédito plus (CCD+)",
	"33":  "Concentración Efectivo / Desembolso Débito plus (CCD+)",
	"34":  "Pago y depósito pre acordado (PPD)",
	"35":  "Concentración efectivo ahorros / Desembolso Crédito (CCD)",
	"36":  "Concentración efectivo ahorros / Desembolso Débito (CCD)",
	"37":  "Pago Negocio Corporativo Ahorros Crédito (CTP)",
	"38":  "Pago Negocio Corporativo Ahorros Débito (CTP)",
	"39":  "Crédito Negocio Intercambio Corporativo (CTX)",
	"40":  "Débito Negocio Intercambio Corporativo (CTX)",
	"41":  "Concentración efectivo/Desembolso Crédito plus (CCD+)",
	"42":  "Consignación bancaria",
	"43":  "Concentración efectivo / Desembolso Débito plus (CCD+)",
	"44":  "Nota cambiaria",
	"45":  "Transferencia Crédito Bancario",
	"46":  "Transferencia Débito Interbancario",
	"47":  "Transferencia Débito Bancaria",
	"48":  "Tarjeta Crédito",
	"49":  "Tarjeta Débito",
	"50":  "Postgiro",
	"51":  "Telex estándar bancario francés",
	"52":  "Pago comercial urgente",
	"53":  "Pago Tesorería Urgente",
	"60":  "Nota promisoria",
	"61":  "Nota promisoria firmada por el acreedor",
	"62":  "Nota promisoria firmada por el acreedor, avalada por el banco",
	"63":  "Nota promisoria firmada por el acreedor, avalada por un tercero",
	"64":  "Nota promisoria firmada por el banco",
	"65":  "Nota promisoria firmada por un banco avalada por otro banco",
	"66":  "Nota promisoria firmada",
	"67":  "Nota promisoria firmada por un tercero avalada por un banco",
	"70":  "Retiro de nota por el acreedor",
	"74":  "Retiro de nota por el acreedor sobre un banco",
	"75":  "Retiro de nota por el acreedor, avalada por otro banco",
	"76":  "Retiro de nota por el acreedor, sobre un banco avalada por un tercero",
	"77":  "Retiro de una nota por el acreedor sobre un tercero",
	"78":  "Retiro de una nota por el acreedor sobre un tercero avalada por un banco",
	"91":  "Nota bancaria transferible",
	"92":  "Cheque local transferible",
	"93":  "Giro referenciado",
	"94":  "Giro urgente",
	"95":  "Giro formato abierto",
	"96":  "Método de pago solicitado no usado",
	"97":  "Clearing entre partners",
	"ZZZ": "Acuerdo mutuo",
}

// IsValidPaymentMeans checks if a payment means code is valid.
func IsValidPaymentMeans(code string) bool {
	_, ok := PaymentMeansFull[PaymentMeans(code)]
	return ok
}

// PaymentMeansName returns the description for a payment means code.
func PaymentMeansName(code string) string {
	return PaymentMeansFull[PaymentMeans(code)]
}
