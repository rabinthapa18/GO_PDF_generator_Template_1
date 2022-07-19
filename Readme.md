# How to write data and images to template file.

### Process
<pre>
  call upload template API     ->      call upload images API    ->    call API to write data to template
           |                                      |                                   |
           V                                      V                                   V
  Image is uploaded to S3               Image uploaded to S3            send body according to the model
           |                                      |                                   |
           V                                      V                                   V
  response is received                 Response received with               Response with PFD file
                                       logoImage and sealImage
                                       which are names for seal
                                       and logo
 </pre>

### There are 3 APIs that you need to call to complete this process.
1. /uploadTemplate      (This will upload the template to s3)
2. /uploadImages        (This will upload images to s3)
3. /addDataToTemplate   (This will take data and write it to the template)


### Request and Response
**1. Upload Template**
          <pre>
          Request
                  TYPE          -     POST, MULTIPART/FORM
                  input keys    -     file
          Response
                  Message       -     Success
          </pre>
**2. Upload Images**
          <pre>
          Request
                  TYPE          -     POST, MULTIPART/FORM
                  input keys    -     seal, logo
          Response
                  logoImage     -     file name for logo
                  sealImage     -     file name for seal
          </pre>
 **3. Request body with data to print to template**
          <pre>
          Request
                  TYPE          -     APPLICATION/JSON
                  input keys    -     template, definitions, values
          Response
                  PDF File
          </pre>


### CURL example for all APIs
**1. uploadTemplate**

`
curl --location --request POST 'localhost:3000/uploadTemplate' \
--form 'file=@"PATH_TO_PDF/pdffile.pdf"'
`

**2. uploadImages**

`
curl --location --request POST 'localhost:3000/uploadImages' \
--form 'seal=@"PATH_TO_SEAL_IMAGE/seal.png"' \
--form 'logo=@"PATH_TO_LOGO_IMAGE/logo.png"'
`

**3. addDataToTemplate**

`
curl --location --request POST 'localhost:3000/addDataToTemplate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "template":"pdf-template-1.pdf",
    "definitions":{
        "texts": [
            {"fieldName": "ClientName","x": 20,"y": 20,"size": 19,"pageNo":1},
            {"fieldName": "AmountBilled","x": 30,"y": 30,"pageNo":1},
            {"fieldName": "PaymentDueDate","x": 40,"y": 40,"pageNo":1}
        ],
        "images": [
            {"name": "seal","x": 10,"y": 10,"width": 100,"height": 100,"pageNo":1},
            {"name": "logo","x": 50,"y": 50,"width": 100,"height": 100,"pageNo":1}
        ]
    },
    "values":{
        "items": [
      {"fieldName": "ClientName", "value": "株式会社 Jackpod"},
      {"fieldName": "AmountBilled", "value": "¥100,000"},
      {"fieldName": "PaymentDueDate", "value": "2022年10月30日"}
    ],
    "images": [
      {"name": "seal", "objectKey": "KEY_RECIEVED_BY_UPLOAD_IMAGE_API"},
      {"name": "logo", "objectKey": "KEY_RECIEVED_BY_UPLOAD_IMAGE_API"}
    ]
    }
}'
`


__*Please do not forget to set enviroment variables and set correct names of template and images names.*__
