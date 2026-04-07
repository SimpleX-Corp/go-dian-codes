// Package codes provides DIAN Colombia electronic invoicing catalogs and validation.
//
// This package contains all the official codes extracted from:
//   - DIAN Caja de Herramientas - Factura Electrónica Validación Previa v1.8
//   - Listas de valores (Genericode .gc files)
//
// # Document Types (TipoDocumento-2.1.gc)
//
//	01 - Factura electrónica de Venta
//	02 - Factura electrónica de venta con propósito de exportación
//	03 - Factura de talonario o papel con numeración de contingencia
//	04 - Factura electrónica de Venta por Contingencia DIAN
//	91 - Nota Crédito
//	92 - Nota Débito
//
// # Tax Types (TipoImpuesto-2.1.gc)
//
//	01 - IVA
//	02 - IC (Impuesto al Consumo)
//	03 - ICA
//	04 - INC (Impuesto Nacional al Consumo)
//	05 - ReteIVA
//	06 - ReteFuente
//	07 - ReteICA
//
// # Fiscal ID Types (TipoIdFiscal-2.1.gc)
//
//	11 - Registro civil
//	12 - Tarjeta de identidad
//	13 - Cédula de ciudadanía
//	21 - Tarjeta de extranjería
//	22 - Cédula de extranjería
//	31 - NIT
//	41 - Pasaporte
//	42 - Documento de identificación extranjero
//	50 - NIT de otro país
//	91 - NUIP
//
// # Payment Methods (FormasPago-2.1.gc)
//
//	1 - Contado
//	2 - Crédito
//
// # Payment Means (MediosPago-2.1.gc)
//
//	10 - Efectivo
//	20 - Cheque
//	30 - Transferencia Crédito
//	42 - Consignación bancaria
//	48 - Tarjeta Crédito
//	49 - Tarjeta Débito
//
// # Usage
//
//	import codes "github.com/SimpleX-Corp/go-dian-codes"
//
//	// Validate a document type
//	if codes.IsValidDocumentType("01") {
//	    name := codes.DocumentTypeName("01") // "Factura electrónica de Venta"
//	}
//
//	// Validate a tax type
//	if codes.IsValidTaxType("01") {
//	    name := codes.TaxTypeName("01") // "IVA"
//	}
//
//	// Validate municipality code
//	if codes.IsValidMunicipality("11001") {
//	    info := codes.GetMunicipality("11001") // Bogotá D.C.
//	}
//
//	// Get department from municipality
//	dept := codes.GetDepartmentCode("11001") // "11" (Bogotá)
//	name := codes.DepartmentName(dept) // "Bogotá"
package codes
