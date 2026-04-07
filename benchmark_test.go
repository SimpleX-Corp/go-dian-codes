package codes

import (
	"testing"
)

// ============================================
// API USAGE EXAMPLES - Idiomatic Go
// ============================================

func Example_usage() {
	// Document types
	if IsValidDocumentType("01") {
		_ = DocumentTypeName("01") // "Factura electrónica de Venta"
	}

	// Tax validation
	if IsValidTaxType("01") {
		_ = TaxTypeName("01") // "IVA"
	}

	// Fiscal ID
	if IsValidFiscalIDType("31") {
		_ = FiscalIDTypeName("31") // "NIT"
	}

	// Geography
	if IsValidMunicipality("11001") {
		mun := GetMunicipality("11001")
		_ = mun.Name       // "Bogotá, D.c."
		_ = mun.Department // "11"
	}

	// Countries & Currencies
	_ = IsValidCountry("CO")
	_ = IsValidCurrencyCode("COP")
	_ = CountryName("CO")
	_ = CurrencyName("COP")

	// Units
	_ = IsValidUnitOfMeasure("KGM")
	_ = UnitOfMeasureName("KGM")

	// Constants for common values
	_ = DocInvoice      // "01"
	_ = TaxIVA          // "01"
	_ = IDNIT           // "31"
	_ = PaymentCash     // "1"
	_ = MeansCash       // "10"
	_ = DeptBogota      // "11"
	_ = IncotermFOB     // "FOB"
	_ = EnvProduction   // "1"
	_ = CurrencyCOP     // "COP"
}

// ============================================
// BENCHMARKS - All lookups should be <20ns
// ============================================

func BenchmarkDocumentType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidDocumentType("01")
	}
}

func BenchmarkDocumentTypeName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DocumentTypeName("01")
	}
}

func BenchmarkTaxType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidTaxType("01")
	}
}

func BenchmarkFiscalIDType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidFiscalIDType("31")
	}
}

func BenchmarkPaymentMeans(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidPaymentMeans("10")
	}
}

func BenchmarkMunicipality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidMunicipality("11001")
	}
}

func BenchmarkMunicipalityGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMunicipality("11001")
	}
}

func BenchmarkCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidCountry("CO")
	}
}

func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidCurrencyCode("COP")
	}
}

func BenchmarkPostalCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidPostalCode("111711")
	}
}

func BenchmarkUnitOfMeasure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidUnitOfMeasure("KGM")
	}
}

func BenchmarkUnitOfMeasureLarge(b *testing.B) {
	// Test with a code near the end of the map
	for i := 0; i < b.N; i++ {
		IsValidUnitOfMeasure("ZZ")
	}
}

// Parallel benchmarks
func BenchmarkMunicipalityParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			IsValidMunicipality("11001")
		}
	})
}

func BenchmarkCountryParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			IsValidCountry("CO")
		}
	})
}
