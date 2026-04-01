package utility

import (
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(to, subject, token, zhTWRemark, enRemark string) error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	smtpAddr := smtpHost + ":" + smtpPort
	senderEmail := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")
	var url, mime, body, message string
	var msg []byte

	switch subject {
	case "Confirm Email":
		subject = "SEL Confirm Email"
		url = os.Getenv("FRONTEND_URL") + "/register?token=" + token
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body = `
		<p>歡迎！非常感謝您註冊本網站。請點擊下方連結以填寫詳細資訊並完成註冊流程。</p>
		<p>Welcome! Thank you for registering on our platform. Please click the link below to provide additional information and complete the registration.</p><br />

		<p>注意：該連結於 10 分鐘後失效。</p>
		<p>Note: This link will expire in 10 minutes.</p><br />

		<p>若您沒有註冊本網站，請忽略此信件。</p>
		<p>If you did not register, please ignore this email.</p><br />

		<p><a href="` + url + `">點擊我完成註冊流程 / Click here to complete your registration</a></p><br />
		<p>Cheers!</p>
	`
		body += GetMailSignature()
		message = "Subject:" + subject + "\r\n" + "To: " + to + "\r\n" + mime + "\r\n" + body
		msg = []byte(message)
	default:
		return nil // If the subject is not recognized, do not send an email
	}

	return smtp.SendMail(
		smtpAddr,
		smtp.PlainAuth("", senderEmail, senderPassword, smtpHost),
		senderEmail,
		[]string{to},
		msg,
	)
}

func GetMailSignature() string {
	officialEmail := os.Getenv("MAIL_DEFAULT_SENDER")
	frontendURL := os.Getenv("FRONTEND_URL")

	return `
	<hr />
	<p>
	客服信箱 / Contact Email: <a href="mailto:` + officialEmail + `">` + officialEmail + `</a><br />
	官方網站 / Official Website: <a href="` + frontendURL + `">` + frontendURL + `</a>
	</p>

	<p style="color: gray;">
	<strong>請勿直接回覆本郵件。</strong><br />
	<strong>Please do not reply to this email.</strong><br />
	如有任何問題，請使用上方聯絡方式與我們聯繫。<br />
	If you have any questions, please contact us using the information above.
	</p>
	`
}
