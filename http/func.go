/*
 * @Author: Justin
 * @Date: 2025-07-21 05:59:00
 * @filename:
 * @version:
 * @Description:
 * @LastEditTime: 2025-07-21 06:47:29
 */

package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test_server/config"

	clog "github.com/kpango/glg"
)

var SingleConfig *config.SingleConfig

func HandleConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	singleConfig := config.DecodeSingleConfig(body)

	clog.Infof("📥 [%s] Body: %s\n", "/StartChain", string(body))
	clog.Infof("📥 [%s] Config: %+v\n", "/StartChain", singleConfig)
	SingleConfig = singleConfig

	resp := Response{
		Status:  "success",
		Message: fmt.Sprintf("%s received", "/config"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func HandleStartChain(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)

	clog.Infof("📥 [%s] Body: %s\n", "/StartChain", string(body))

	resp := Response{
		Status:  "success",
		Message: fmt.Sprintf("%s received", "/StartChain"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func HandleSetupContracts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)

	clog.Infof("📥 [%s] Body: %s\n", "/SetupContracts", string(body))

	resp := Response{
		Status:  "success",
		Message: fmt.Sprintf("%s received", "/SetupContracts"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func HandleStartRelayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, _ := io.ReadAll(r.Body)

	clog.Infof("📥 [%s] Body: %s\n", "/StartRelayer", string(body))

	resp := Response{
		Status:  "success",
		Message: fmt.Sprintf("%s received", "/StartRelayer"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
