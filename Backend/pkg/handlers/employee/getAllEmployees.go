package employee

import (
	"ESM-backend-app/pkg/mocks"
	"encoding/json"
	"net/http"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Employees)
}
