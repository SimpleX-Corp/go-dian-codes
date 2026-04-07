package codes

import "testing"

func TestDocumentTypes(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
		name  string
	}{
		{"01", true, "Factura electrónica de Venta"},
		{"02", true, "Factura electrónica de venta con propósito de exportación"},
		{"91", true, "Nota Crédito"},
		{"92", true, "Nota Débito"},
		{"99", false, ""},
	}

	for _, tt := range tests {
		if got := IsValidDocumentType(tt.code); got != tt.valid {
			t.Errorf("IsValidDocumentType(%q) = %v, want %v", tt.code, got, tt.valid)
		}
		if tt.valid {
			if got := DocumentTypeName(tt.code); got != tt.name {
				t.Errorf("DocumentTypeName(%q) = %q, want %q", tt.code, got, tt.name)
			}
		}
	}
}

func TestTaxTypes(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
		name  string
	}{
		{"01", true, "IVA"},
		{"02", true, "IC"},
		{"05", true, "ReteIVA"},
		{"06", true, "ReteFuente"},
		{"99", false, ""},
	}

	for _, tt := range tests {
		if got := IsValidTaxType(tt.code); got != tt.valid {
			t.Errorf("IsValidTaxType(%q) = %v, want %v", tt.code, got, tt.valid)
		}
		if tt.valid {
			if got := TaxTypeName(tt.code); got != tt.name {
				t.Errorf("TaxTypeName(%q) = %q, want %q", tt.code, got, tt.name)
			}
		}
	}
}

func TestFiscalIDTypes(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
		name  string
	}{
		{"13", true, "Cédula de ciudadanía"},
		{"31", true, "NIT"},
		{"41", true, "Pasaporte"},
		{"99", false, ""},
	}

	for _, tt := range tests {
		if got := IsValidFiscalIDType(tt.code); got != tt.valid {
			t.Errorf("IsValidFiscalIDType(%q) = %v, want %v", tt.code, got, tt.valid)
		}
		if tt.valid {
			if got := FiscalIDTypeName(tt.code); got != tt.name {
				t.Errorf("FiscalIDTypeName(%q) = %q, want %q", tt.code, got, tt.name)
			}
		}
	}
}

func TestPaymentMethods(t *testing.T) {
	if !IsValidPaymentMethod("1") {
		t.Error("PaymentMethod '1' (Contado) should be valid")
	}
	if !IsValidPaymentMethod("2") {
		t.Error("PaymentMethod '2' (Crédito) should be valid")
	}
	if IsValidPaymentMethod("99") {
		t.Error("PaymentMethod '99' should be invalid")
	}
}

func TestPaymentMeans(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"10", true},  // Efectivo
		{"48", true},  // Tarjeta Crédito
		{"49", true},  // Tarjeta Débito
		{"ZZZ", true}, // Acuerdo mutuo
		{"999", false},
	}

	for _, tt := range tests {
		if got := IsValidPaymentMeans(tt.code); got != tt.valid {
			t.Errorf("IsValidPaymentMeans(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestDepartments(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
		name  string
	}{
		{"11", true, "Bogotá"},
		{"05", true, "Antioquia"},
		{"76", true, "Valle del Cauca"},
		{"00", false, ""},
	}

	for _, tt := range tests {
		if got := IsValidDepartment(tt.code); got != tt.valid {
			t.Errorf("IsValidDepartment(%q) = %v, want %v", tt.code, got, tt.valid)
		}
		if tt.valid {
			if got := DepartmentName(tt.code); got != tt.name {
				t.Errorf("DepartmentName(%q) = %q, want %q", tt.code, got, tt.name)
			}
		}
	}
}

func TestMunicipalities(t *testing.T) {
	// Test Bogotá
	if !IsValidMunicipality("11001") {
		t.Error("Municipality '11001' (Bogotá) should be valid")
	}

	mun := GetMunicipality("11001")
	if mun == nil {
		t.Fatal("GetMunicipality('11001') returned nil")
	}
	if mun.Department != "11" {
		t.Errorf("Municipality department = %q, want '11'", mun.Department)
	}

	// Test Medellín
	if !IsValidMunicipality("05001") {
		t.Error("Municipality '05001' (Medellín) should be valid")
	}

	// Test invalid
	if IsValidMunicipality("00000") {
		t.Error("Municipality '00000' should be invalid")
	}

	// Test GetDepartmentCode
	if got := GetDepartmentCode("11001"); got != "11" {
		t.Errorf("GetDepartmentCode('11001') = %q, want '11'", got)
	}
}

func TestTaxResponsibilities(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"O-13", true}, // Gran contribuyente
		{"O-15", true}, // Autorretenedor
		{"O-47", true}, // Régimen simple
		{"ZZ", true},   // No aplica
		{"X-99", false},
	}

	for _, tt := range tests {
		if got := IsValidTaxResponsibility(tt.code); got != tt.valid {
			t.Errorf("IsValidTaxResponsibility(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestCreditNoteCorrections(t *testing.T) {
	if !IsValidCreditNoteCorrection("1") {
		t.Error("CreditNoteCorrection '1' should be valid")
	}
	if !IsValidCreditNoteCorrection("2") {
		t.Error("CreditNoteCorrection '2' should be valid")
	}
	if IsValidCreditNoteCorrection("9") {
		t.Error("CreditNoteCorrection '9' should be invalid")
	}
}

func TestDebitNoteCorrections(t *testing.T) {
	if !IsValidDebitNoteCorrection("1") {
		t.Error("DebitNoteCorrection '1' (Intereses) should be valid")
	}
	if !IsValidDebitNoteCorrection("4") {
		t.Error("DebitNoteCorrection '4' (Otro) should be valid")
	}
	if IsValidDebitNoteCorrection("9") {
		t.Error("DebitNoteCorrection '9' should be invalid")
	}
}

func TestOperationTypes(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"10", true}, // Estandar
		{"09", true}, // AIU
		{"11", true}, // Mandatos
		{"99", false},
	}

	for _, tt := range tests {
		if got := IsValidOperationType(tt.code); got != tt.valid {
			t.Errorf("IsValidOperationType(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestCountries(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"CO", true}, // Colombia
		{"US", true}, // United States
		{"XX", false},
	}

	for _, tt := range tests {
		if got := IsValidCountry(tt.code); got != tt.valid {
			t.Errorf("IsValidCountry(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}

	// Test CountryName
	if name := CountryName("CO"); name == "" {
		t.Error("CountryName('CO') should not be empty")
	}
}

func TestCurrencies(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"COP", true}, // Colombian Peso
		{"USD", true}, // US Dollar
		{"EUR", true}, // Euro
		{"XXX", false},
	}

	for _, tt := range tests {
		if got := IsValidCurrencyCode(tt.code); got != tt.valid {
			t.Errorf("IsValidCurrencyCode(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestPostalCodes(t *testing.T) {
	// Test a known postal code
	if !IsValidPostalCode("111711") {
		// Try another known code
		found := false
		for code := range postalCodes {
			if IsValidPostalCode(code) {
				found = true
				break
			}
		}
		if !found {
			t.Error("No valid postal codes found")
		}
	}

	// Invalid postal code
	if IsValidPostalCode("000000") {
		t.Error("PostalCode '000000' should be invalid")
	}
}

func TestUnitsOfMeasure(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"EA", true},   // Each
		{"KGM", true},  // Kilogram
		{"LTR", true},  // Liter
		{"XXXX", false},
	}

	for _, tt := range tests {
		if got := IsValidUnitOfMeasure(tt.code); got != tt.valid {
			t.Errorf("IsValidUnitOfMeasure(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestEnvironments(t *testing.T) {
	if !IsValidEnvironment("1") {
		t.Error("Environment '1' (Production) should be valid")
	}
	if !IsValidEnvironment("2") {
		t.Error("Environment '2' (Test) should be valid")
	}
	if IsValidEnvironment("3") {
		t.Error("Environment '3' should be invalid")
	}
}

func TestDocumentEvents(t *testing.T) {
	if !IsValidDocumentEvent("033") {
		t.Error("DocumentEvent '033' (Accepted) should be valid")
	}
	if IsValidDocumentEvent("999") {
		t.Error("DocumentEvent '999' should be invalid")
	}
}

func TestDeliveryTerms(t *testing.T) {
	tests := []struct {
		code  string
		valid bool
	}{
		{"FOB", true},
		{"CIF", true},
		{"EXW", true},
		{"XXX", false},
	}

	for _, tt := range tests {
		if got := IsValidDeliveryTerm(tt.code); got != tt.valid {
			t.Errorf("IsValidDeliveryTerm(%q) = %v, want %v", tt.code, got, tt.valid)
		}
	}
}

func TestDiscountCodes(t *testing.T) {
	if !IsValidDiscountCode("00") {
		t.Error("DiscountCode '00' (unconditional) should be valid")
	}
	if !IsValidDiscountCode("01") {
		t.Error("DiscountCode '01' (conditional) should be valid")
	}
}

func TestProductCodeTypes(t *testing.T) {
	if !IsValidProductCodeType("001") {
		t.Error("ProductCodeType '001' (UNSPSC) should be valid")
	}
	if !IsValidProductCodeType("010") {
		t.Error("ProductCodeType '010' (GTIN) should be valid")
	}
}

func BenchmarkIsValidMunicipality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidMunicipality("11001")
	}
}

func BenchmarkGetMunicipality(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMunicipality("11001")
	}
}

func BenchmarkIsValidCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidCountry("CO")
	}
}

func BenchmarkIsValidUnitOfMeasure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidUnitOfMeasure("KGM")
	}
}
