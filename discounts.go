package codes

// DiscountCode represents discount type codes.
// Source: DIAN Caja de Herramientas - CodigoDescuento-2.1.gc
type DiscountCode string

const (
	DiscountUnconditional DiscountCode = "00" // Descuento no condicionado
	DiscountConditional   DiscountCode = "01" // Descuento condicionado
)

var DiscountCodes = map[DiscountCode]string{
	DiscountUnconditional: "Descuento no condicionado",
	DiscountConditional:   "Descuento condicionado",
}

// IsValidDiscountCode checks if a discount code is valid.
func IsValidDiscountCode(code string) bool {
	_, ok := DiscountCodes[DiscountCode(code)]
	return ok
}

// ReferencePriceCode represents reference price codes.
// Source: DIAN Caja de Herramientas - CodigoPrecioReferencia-2.1.gc
type ReferencePriceCode string

const (
	RefPriceCommercial ReferencePriceCode = "01" // Valor comercial
	RefPriceInventory  ReferencePriceCode = "02" // Valor en inventarios
	RefPriceOther      ReferencePriceCode = "03" // Otro valor
)

var ReferencePriceCodes = map[ReferencePriceCode]string{
	RefPriceCommercial: "Valor comercial",
	RefPriceInventory:  "Valor en inventarios",
	RefPriceOther:      "Otro valor",
}

// IsValidReferencePriceCode checks if a reference price code is valid.
func IsValidReferencePriceCode(code string) bool {
	_, ok := ReferencePriceCodes[ReferencePriceCode(code)]
	return ok
}

// ProductCodeType represents product code scheme types.
// Source: DIAN Caja de Herramientas - TipoCodigoProducto-2.1.gc
type ProductCodeType string

const (
	ProductCodeUNSPSC      ProductCodeType = "001" // UNSPSC
	ProductCodeGTIN        ProductCodeType = "010" // GTIN
	ProductCodeContributor ProductCodeType = "999" // Estándar de adopción del contribuyente
)

var ProductCodeTypes = map[ProductCodeType]string{
	ProductCodeUNSPSC:      "UNSPSC",
	ProductCodeGTIN:        "GTIN",
	ProductCodeContributor: "Estándar de adopción del contribuyente",
}

// IsValidProductCodeType checks if a product code type is valid.
func IsValidProductCodeType(code string) bool {
	_, ok := ProductCodeTypes[ProductCodeType(code)]
	return ok
}
