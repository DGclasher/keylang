package types

type LicenseService interface {
	CreateLicense(email string, productID string) (*LicenseCreateType, error)
	ValidateLicense(licenseKey string, productId string) (*LicenseValidateType, error)
}

type LicenseCreateType struct {
	UserEmail string `json:"user_email"`
	ProductID string `json:"product_id"`
	LicenseKey string `json:"license_key"`
}

type LicenseValidateType struct {
	LicenseKey string `json:"license_key"`
	ProductID string `json:"product_id"`
	UserEmail string `json:"user_email"`
}