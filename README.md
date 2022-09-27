# Whatsapp Cloud API Wrapper for Golang

## About

Lightweight Whatsapp Cloud API Wrapper v14.0 for Golang with no dependency.

Made this cause I haven't found any reliable whatsapp cloud api wrapper for golang

Content

- [Installation](#installation)
- [Getting the Whatsapp Cloud API Access](#getting-the-whatsapp-cloud-api-access)
- [FAQ](#faq)
- [Usage](#usage)

## Instal

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

### Send a message with one parameter

```
parameters := wa.GenerateTemplateParameters("text", "48884")
components := wa.TemplateComponent("body", parameters)
request := wa.CreateSendTemplateRequest("RECEIVER_PHONE_NUMBER", "your_template_name", whatsapp.Indonesian, components)
res, err := wa.SendWithTemplate(request)
```

### Send a message with no parameter

```
request := wa.CreateSendTemplateRequest("RECEIVER_PHONE_NUMBER", "your_template_name", whatsapp.Indonesian, nil)
res, err := wa.SendWithTemplate(request)
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

### Changing Whatsapp Cloud API version

```
wa := whatsapp.NewWhatsapp(token, myPhoneID)
wa.APIVersion = "v2"
```

## Tutorial for Myself!

Releasing new version

```
$ git commit -m "hello: changes for v1.0.0"
$ git tag v1.0.0
$ git push origin v1.0.0
```
