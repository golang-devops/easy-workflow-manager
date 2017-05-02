package requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang-devops/easy-workflow-manager/logging"
)

func NewHelper(writer http.ResponseWriter, request *http.Request) *Helper {
	return &Helper{
		writer:  writer,
		request: request,
	}
}

type Helper struct {
	writer  http.ResponseWriter
	request *http.Request
}

func (h *Helper) Request() *http.Request {
	return h.request
}

func (h *Helper) CheckMethod(logger logging.Logger, allowedMethods ...string) error {
	for _, allowedMethod := range allowedMethods {
		if strings.EqualFold(h.request.Method, allowedMethod) {
			return nil
		}
	}

	userMessage := "Method not allowed"
	h.WriteError(
		logger.WithFields(map[string]interface{}{
			"method":          h.request.Method,
			"allowed-methods": allowedMethods,
		}),
		http.StatusMethodNotAllowed,
		userMessage,
	)
	return errors.New(userMessage)
}

func (h *Helper) CheckBodyContentSignatureFromHeader(logger logging.Logger, signatureHeaderKey string, getExpectedSignature func(bodyBytes []byte) string) ([]byte, error) {
	logger = logger.WithField("signature-header-key", signatureHeaderKey)

	tmpActualSignature, foundSignature := h.request.Header[signatureHeaderKey]
	if !foundSignature {
		userMessage := fmt.Sprintf("Unable to find signature header '%s'", signatureHeaderKey)
		h.WriteError(logger, http.StatusBadRequest, userMessage)
		return nil, errors.New(userMessage)
	}
	actualSignature := tmpActualSignature[0]
	logger = logger.WithField("actual-signature", actualSignature)

	bodyBytes, err := ioutil.ReadAll(h.request.Body)
	if err != nil {
		userMessage := "Failed to read request Body content"
		h.WriteError(logger.WithError(err), http.StatusBadRequest, userMessage)
		return nil, errors.New(userMessage)
	}

	expectedSignature := getExpectedSignature(bodyBytes)
	logger = logger.WithField("expected-signature", expectedSignature)

	if actualSignature != expectedSignature {
		userMessage := "Signature mismatch"
		h.WriteError(logger, http.StatusUnauthorized, userMessage)
		return nil, errors.New(userMessage)
	}

	logger.Info("Request signature Verified")
	return bodyBytes, nil
}

func (h *Helper) WriteInfo(logger logging.Logger, status int, userMessage string) {
	h.writer.WriteHeader(status)
	h.writeString(logger, h.writer, userMessage)
	logger.Info(userMessage)
}

func (h *Helper) WriteWarn(logger logging.Logger, status int, userMessage string) {
	h.writer.WriteHeader(status)
	h.writeString(logger, h.writer, userMessage)
	logger.Warn(userMessage)
}

func (h *Helper) WriteError(logger logging.Logger, status int, userMessage string) {
	h.writer.WriteHeader(status)
	h.writeString(logger, h.writer, userMessage)
	logger.Error(userMessage)
}

func (h *Helper) writeString(logger logging.Logger, writer http.ResponseWriter, str string) {
	if _, err := writer.Write([]byte(str)); err != nil {
		logger.Error(fmt.Sprintf("Failed to write response. Write error: %s. Response to write was: %s", err.Error(), str))
	}
}
