# TODO

- [ ] Creating media sending wrapper https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media
    ```
    curl -X POST \
    'https://graph.facebook.com/v14.0/FROM_PHONE_NUMBER_ID/media' \
    -H 'Authorization: Bearer ACCESS_TOKEN' \
    -F 'file=@/local/path/file.jpg;type=image/jpeg' 
    -F 'messaging_product=whatsapp'
    ```

    the response: `{"id":"ID"}`