package askpermission

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-devops/easy-workflow-manager/helpers/http/requests"
	"github.com/golang-devops/easy-workflow-manager/helpers/mailing"
	"github.com/golang-devops/easy-workflow-manager/helpers/webhooks"
	"github.com/golang-devops/easy-workflow-manager/logging"
	"github.com/golang-devops/easy-workflow-manager/util/tokens"
)

func NewEmailAndWebhook(logger logging.Logger, listenAddress, baseURL string, sendGridConfig *mailing.SendGridConfig, recipientsProvider mailing.AddressesProvider, question string, answers []string) QuestionAsker {
	linkSecret := tokens.RandomAlphaNumericString(64)

	answerChan := make(chan string, 1)
	requestHandler := &waitForLinkedClickedRequestHandler{
		logger:       logger,
		linkSecret:   linkSecret,
		answerChan:   answerChan,
		validAnswers: answers,
	}

	return &EmailAndWebhook{
		logger:             logger,
		listener:           webhooks.NewListener(listenAddress, requestHandler),
		answerChan:         answerChan,
		baseURL:            baseURL,
		linkSecret:         linkSecret,
		sendGridConfig:     sendGridConfig,
		recipientsProvider: recipientsProvider,
		question:           question,
		answers:            answers,
	}
}

type EmailAndWebhook struct {
	logger logging.Logger

	listener   *webhooks.Listener
	answerChan chan string

	baseURL            string
	linkSecret         string
	sendGridConfig     *mailing.SendGridConfig
	recipientsProvider mailing.AddressesProvider
	question           string
	answers            []string
}

func (e *EmailAndWebhook) sendEmail() error {
	subject := fmt.Sprintf("Human confirmation needed. %s?", e.question)

	htmlLines := []string{}
	htmlLines = append(htmlLines, fmt.Sprintf(`<h3>%s</h3>`, e.question))

	htmlLines = append(htmlLines, `<ul>`)
	for _, answer := range e.answers {
		queryValues := url.Values{}
		queryValues.Set("token", e.linkSecret)
		queryValues.Set("answer", answer)
		linkHref := strings.TrimRight(e.baseURL, "/") + "?" + queryValues.Encode()
		htmlLines = append(htmlLines, fmt.Sprintf(`<li> <a href="%s">%s</a> </li>`, linkHref, answer))
	}
	htmlLines = append(htmlLines, `</ul>`)

	recipients := e.recipientsProvider.Addresses()
	substitutions := make(map[string]string)

	errorStrs := []string{}
	for _, recipient := range recipients {
		if err := mailing.SendHTMLEmail(e.logger, e.sendGridConfig, subject, strings.Join(htmlLines, "\n"), recipient, substitutions); err != nil {
			errorStrs = append(errorStrs, fmt.Sprintf("Failed to send to '%s', error: %s", recipient.Address, err.Error()))
		}
	}

	if len(errorStrs) > 0 {
		return errors.New(strings.Join(errorStrs, " & "))
	}
	return nil
}

func (e *EmailAndWebhook) GetAnswer() (string, error) {
	if err := e.sendEmail(); err != nil {
		e.logger.WithError(err).Error("Failed to Send email")
		return "", err
	}

	if err := e.listener.ListenAndWait(); err != nil {
		e.logger.WithError(err).Error("Failed to Listen and Wait")
		return "", err
	}

	answer := <-e.answerChan
	return answer, nil
}

type waitForLinkedClickedRequestHandler struct {
	logger       logging.Logger
	linkSecret   string
	validAnswers []string

	answerChan chan string
	isComplete bool
}

func (w *waitForLinkedClickedRequestHandler) isAnswerValid(answer string) bool {
	for _, validAnswer := range w.validAnswers {
		if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(validAnswer)) {
			return true
		}
	}
	return false
}

func (w *waitForLinkedClickedRequestHandler) HandleRequest(writer http.ResponseWriter, req *http.Request) {
	helper := requests.NewHelper(writer, req)

	if err := helper.CheckMethod(w.logger, "GET"); err != nil {
		return
	}

	query := req.URL.Query()

	token := query.Get("token")
	if token == "" {
		userMessage := "Failed to get token from URL"
		helper.WriteError(w.logger.WithField("raw-url-query", req.URL.RawQuery), http.StatusBadRequest, userMessage)
		return
	}

	answer := strings.ToLower(query.Get("answer"))
	if !w.isAnswerValid(answer) {
		userMessage := fmt.Sprintf("Invalid answer '%s' from URL", answer)
		helper.WriteError(w.logger.WithField("raw-url-query", req.URL.RawQuery), http.StatusBadRequest, userMessage)
		return
	}

	w.answerChan <- answer
	w.isComplete = true
	helper.WriteInfo(w.logger, http.StatusOK, fmt.Sprintf("Thank you, answer '%s' received", answer))
}

func (w *waitForLinkedClickedRequestHandler) IsComplete() bool {
	return w.isComplete
}

func (w *waitForLinkedClickedRequestHandler) Error() error {
	return nil
}
