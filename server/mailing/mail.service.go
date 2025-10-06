package mailing

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/FrancoMusolino/go-todoapp/utils/logger"
	mail "github.com/xhit/go-simple-mail/v2"
)

type SimpleMailService struct {
	config   *MailConfig
	jobQueue chan Message
	logger   logger.Logger
}

func NewSimpleMailService(config *MailConfig, jobQueue chan Message) *SimpleMailService {
	return &SimpleMailService{
		config:   config,
		jobQueue: jobQueue,
		logger:   *logger.NewLogger("Mail Service"),
	}
}

func (m *SimpleMailService) getServer() *mail.SMTPServer {
	m.logger.Info(context.Background(), "getServer", "Creating SMTP Server")
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = m.config.Host
	server.Port = m.config.Port
	server.Username = m.config.Username
	server.Password = m.config.Password
	server.Encryption = m.getEncryption(m.config.Encryption)

	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	return server
}

func (m *SimpleMailService) getEncryption(e string) mail.Encryption {
	switch e {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSL
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}

func (m *SimpleMailService) SendHTML(msg Message) error {
	smtpClient, err := m.getServer().Connect()
	if err != nil {
		return err
	}

	if msg.FromAddress == "" {
		msg.FromAddress = m.config.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.config.FromName
	}

	if msg.ToAddresses == "" {
		return errors.New("Missing To Addresses for email")
	}

	if msg.Subject == "" {
		return errors.New("Missing Subject for sending email")
	}

	email := mail.NewMSG()
	email.SetFrom(fmt.Sprintf("%s <%s>", msg.FromName, msg.FromAddress))

	for _, toAddress := range strings.Split(msg.ToAddresses, ";") {
		email.AddTo(toAddress)
	}

	for _, ccAddress := range strings.Split(msg.CCAddresses, ";") {
		email.AddTo(ccAddress)
	}

	email.SetSubject(msg.Subject)
	email.SetBody(mail.TextHTML, msg.Body)

	if email.Error != nil {
		return email.Error
	}

	err = email.Send(smtpClient)
	if err != nil {
		return err
	}

	m.logger.Info(context.Background(), "SendHTML", "Send email successfully")
	return nil
}

func (m *SimpleMailService) SendHTMLAsync(msg Message) {
	m.jobQueue <- msg
}
