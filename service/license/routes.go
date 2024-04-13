package license

import (
	"fmt"
	"net/http"

	"github.com/DGclasher/keylang/types"
	"github.com/DGclasher/keylang/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	licenseService types.LicenseService
}

func NewHandler(license types.LicenseService) *Handler {
	return &Handler{licenseService: license}
}

func (h *Handler) LicenseRoutes(router *mux.Router) {
	router.HandleFunc("/license/create", h.handleLicenseCreation).Methods("POST")
	router.HandleFunc("/license/validate", h.handleLicenseValidation).Methods("POST")
}

func (h *Handler) handleLicenseCreation(w http.ResponseWriter, r *http.Request) {
	var licenseCreateType types.LicenseCreateType
	if err := utils.ParseJSON(r, &licenseCreateType); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	createdLicense, err := h.licenseService.CreateLicense(licenseCreateType.UserEmail, licenseCreateType.ProductID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, createdLicense)
}

func (h *Handler) handleLicenseValidation(w http.ResponseWriter, r *http.Request) {
	var licenseValidateType types.LicenseValidateType
	if err := utils.ParseJSON(r, &licenseValidateType); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	data, err := h.licenseService.ValidateLicense(licenseValidateType.LicenseKey, licenseValidateType.ProductID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if data.UserEmail == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("License not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, data)
}
