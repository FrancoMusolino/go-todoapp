package mailing

type MailConfig struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	FromAddress string
	FromName    string
	ToAddresses string
	CCAddresses string
	Subject     string
	Body        string
}

type IMailService interface {
	SendHTML(m Message) error
}
