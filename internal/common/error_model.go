package common

type ErrorModel struct {
	Code      int
	ErrorCode string
	CausedBy  error
	// ErrorParameter        []ErrorParameter
	AdditionalInformation interface{}
}

func GenerateErrorModel(code int, errCode string, causedBy error) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.ErrorCode = errCode
	errModel.CausedBy = causedBy
	return errModel
}

// Fungsi untuk menghasilkan error Unauthorized
func GenerateUnauthorizedError() ErrorModel {
	return GenerateErrorModel(401, "E-401-AUTH-1", nil) // Token tidak valid atau telah kedaluwarsa
}

// Fungsi untuk menghasilkan error Session Expired
func GenerateSessionExpiredError() ErrorModel {
	return GenerateErrorModel(401, "E-401-AUTH-2", nil) // Sesi telah kedaluwarsa, silakan masuk lagi
}

// Fungsi untuk menghasilkan error Forbidden
func GenerateForbiddenError() ErrorModel {
	return GenerateErrorModel(403, "E-403-ACL-1", nil) // Anda tidak memiliki izin untuk melakukan tindakan ini
}

// Fungsi untuk menghasilkan error Access Forbidden
func GenerateAccessForbiddenError() ErrorModel {
	return GenerateErrorModel(403, "E-403-ACL-2", nil) // Akses terlarang
}

// Fungsi untuk menghasilkan error Bad Request - Field Required
func GenerateFieldRequiredError(field string) ErrorModel {
	return GenerateErrorModel(400, "E-400-VAL-1", nil) // {{.field}} diperlukan
}

// Fungsi untuk menghasilkan error Bad Request - Minimum Characters
func GenerateMinimumCharactersError(field string, min int) ErrorModel {
	return GenerateErrorModel(400, "E-400-VAL-2", nil) // {{.field}} harus terdiri dari setidaknya {{.min}} karakter
}

// Fungsi untuk menghasilkan error Bad Request - Maximum Characters
func GenerateMaximumCharactersError(field string, max int) ErrorModel {
	return GenerateErrorModel(400, "E-400-VAL-3", nil) // {{.field}} tidak boleh melebihi {{.max}} karakter
}

// Fungsi untuk menghasilkan error Bad Request - Invalid Format
func GenerateInvalidFormatError(field string) ErrorModel {
	return GenerateErrorModel(400, "E-400-VAL-4", nil) // Format {{.field}} tidak valid
}

// Fungsi untuk menghasilkan error Bad Request - Missing DTO Fields
func GenerateMissingDTOFieldsError() ErrorModel {
	return GenerateErrorModel(400, "E-400-DTO-5", nil) // Field DTO yang diperlukan hilang
}

// Fungsi untuk menghasilkan error Bad Request - Validation Failed
func GenerateValidationFailedError() ErrorModel {
	return GenerateErrorModel(400, "E-400-DTO-6", nil) // Validasi gagal
}

// Fungsi untuk menghasilkan error Bad Request - Invalid JSON Format
func GenerateInvalidJSONFormatError() ErrorModel {
	return GenerateErrorModel(400, "E-400-JSON-7", nil) // Format JSON tidak valid
}

// Fungsi untuk menghasilkan error Internal Server Error
func GenerateInternalServerError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-500-SRV-1", causedBy) // Kesalahan server internal
}

// Fungsi untuk menghasilkan error Database Error
func GenerateDatabaseError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-500-DBS-2", causedBy) // Kesalahan basis data
}

// Fungsi untuk menghasilkan error Service Unavailable
func GenerateServiceUnavailableError() ErrorModel {
	return GenerateErrorModel(500, "E-500-SRV-3", nil) // Layanan tidak tersedia
}
