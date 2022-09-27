package whatsapp

// Templates in https://business.facebook.com/wa/manage/message-templates
const (
	TemplateVerifyPhoneNumberID               string = "verify_phone_number_id"
	TemplateCertificateAnnouncementIndonesian string = "certificate_announcement_indonesian"
	TemplatePurchaseFeedback                  string = "purchase_feedback"
	TemplateOTP                               string = "otp"
)

var (
	LanguageIndonesian = TemplateLanguage{
		Code: "id",
	}
)
