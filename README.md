# Whatsapp Cloud API Wrapper for Golang

## About

Lightweight Whatsapp Cloud API Wrapper v14.0 for Golang with no dependency.

Made this cause I haven't found any reliable whatsapp cloud api wrapper for golang

Content

- [Install](#install)
- [Usage](#usage)
- [Getting the Whatsapp Cloud API Access](#getting-the-whatsapp-cloud-api-access)
- [FAQ](#faq)
- [Usage](#usage)

## Install

```
go get github.com/febriliankr/whatsapp-cloud-api
```

or if you want to use specific version

```
go get github.com/febriliankr/whatsapp-cloud-api@v1.0.2
```

## Usage

### Create an instance of the Whatsapp Cloud API Client

```
wa := whatsapp.NewWhatsapp(token, myPhoneID)
```

Modifying instance 
- Changing the language `wa.Language = TemplateLanguage{ Code: "id" }`
- Changing the API version `wa.APIVersion = "v14.0"`
- Changing the Whatsapp Phone ID `wa.PhoneNumberID = WHATSAPP_PHONE_ID`

### Send a templated message

With one parameter

```
parameters := wa.GenerateTemplateParameters("text", "48884")
components := wa.TemplateComponent("body", parameters)
res, err := wa.SendWithTemplate("RECEIVER_PHONE_NUMBER", "your_template_name", components)
```

with no parameter

```
res, err := wa.SendWithTemplate("RECEIVER_PHONE_NUMBER", "your_template_name", nil)
```

### Send a plain text message

Sending plain text message to a phone number that has messaged your WhatsApp Business account in the past 24 hours.

```
res, err := wa.SendText("RECEIVER_PHONE_NUMBER", "your_message")
```

Check more usage example in example/example.go

## Getting the Whatsapp Cloud API Access

Official documentation:
- https://developers.facebook.com/docs/whatsapp?locale=en_US
- https://developers.facebook.com/docs/whatsapp/cloud-api/get-started
- Console: https://developers.facebook.com/apps/2017523718408931/whatsapp-business/wa-dev-console

## FAQ

- Throughput?

  - Cloud API supports 80 messages per second (mps) for text and media as default, up to 250 mps by request. (https://developers.facebook.com/docs/whatsapp/cloud-api/overview)

- Question: How do I signup to get the token?

  - Answer: Here https://developers.facebook.com. Registration guide here: https://developers.facebook.com/docs/development/register/

- Question: Whats the difference between Qiscus and other provider's Whatsapp API to this Cloud API?
  - Answer: This Cloud API is provided directly by Meta, no need to use 3rd party provider.
  - More: https://techcrunch.com/2022/05/19/whatsapp-ramps-up-revenue-with-global-launch-of-cloud-api-and-soon-a-paid-tier-for-its-business-app/
  - https://www.qiscus.com/id/blog/meta-luncurkan-whatsap-cloud-api-untuk-publik/

## Tutorial for Myself!

Releasing new version

```
$ git commit -m "hello: changes for v1.0.0"
$ git tag v1.0.0
$ git push origin v1.0.0
```
