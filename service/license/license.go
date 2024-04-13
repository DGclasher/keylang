package license

import (
	"database/sql"

	"github.com/DGclasher/keylang/types"
	"github.com/DGclasher/keylang/utils"
)

type License struct {
	db *sql.DB
}

func NewLicense(db *sql.DB) *License {
	return &License{db: db}
}

func (l *License) CreateLicense(email string, productID string) (*types.LicenseCreateType, error) {
	licenseKey := utils.GenLicenseKey(email, productID)
	_, err := l.db.Exec("INSERT INTO licenses (user_email, product_id, license_key) VALUES ($1, $2, $3)", email, productID, licenseKey)
	if err != nil {
		return nil, err
	}
	return &types.LicenseCreateType{UserEmail: email, ProductID: productID, LicenseKey: licenseKey}, nil
}

func (l *License) ValidateLicense(licenseKey string, productId string) (*types.LicenseValidateType, error) {
	rows, err := l.db.Query("SELECT license_key, product_id, user_email FROM licenses WHERE license_key = $1 AND product_id = $2", licenseKey, productId)
	if err != nil {
		return nil, err
	}
	ll := new(types.LicenseValidateType)
	for rows.Next() {
		ll, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}
	return ll, nil
}

func scanRow(rows *sql.Rows) (*types.LicenseValidateType, error) {
	ll := new(types.LicenseValidateType)
	err := rows.Scan(&ll.LicenseKey, &ll.ProductID, &ll.UserEmail)
	if err != nil {
		return nil, err
	}
	return ll, nil
}