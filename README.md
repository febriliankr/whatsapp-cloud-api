# Whatsapp Cloud API Wrapper for Golang

v14.0 Whatsapp Cloud API Wrapper for Golang

Made this cause I haven't found any reliable whatsapp cloud api wrapper for golang

Go to your facebook developer console to get the token https://developers.facebook.com/apps/2017523718408931/whatsapp-business/wa-dev-console

## Usage

```
wa := whatsapp.NewWhatsapp(token, myPhoneId)


// Send whatsapp text without parameters
request := wa.CreateSendTemplateRequest("RECEIVER_PHONE_NUMBER", "your_template_name", whatsapp.Indonesian, nil)
res, err := wa.SendWithTemplate(request)

// Send whatsapp text with parameters
parameters := wa.GenerateTemplateParameters("text", "48884")
components := wa.TemplateComponent("body", parameters)
request := wa.CreateSendTemplateRequest("RECEIVER_PHONE_NUMBER", "your_template_name", whatsapp.Indonesian, components)
res, err := wa.SendWithTemplate(request)
```

Check more usage example in example/example.go
