package test

import (
	"fmt"
	"testing"
)

func TestUnitTest(t *testing.T) {
	uoms := []string{"BOX", "PCS", "PCS", "PCS"}
	barcodes := []string{"BR01", "", "BR04", ""}

	uomBarcodeMap := make(map[string]string)

	// Isi map dengan barcode yang ada
	for i := 0; i < len(uoms); i++ {
		if uoms[i] != "" && barcodes[i] != "" {
			uomBarcodeMap[uoms[i]] = barcodes[i]
		}
	}

	// Isi barcode yang kosong dengan nilai dari map
	for i := 0; i < len(uoms); i++ {
		if uoms[i] != "" && barcodes[i] == "" {
			if val, exists := uomBarcodeMap[uoms[i]]; exists {
				barcodes[i] = val
			}
		}
	}

	for i, v := range barcodes {
		fmt.Printf("uom %d : %s\n", i, v)
	}
}
