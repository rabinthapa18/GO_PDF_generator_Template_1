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

1. /uploadTemplate (This will upload the template to s3)
2. /uploadImages (This will upload images to s3)
3. /addDataToTemplate (This will take data and write it to the template)

### Request and Response

**1. Upload Template**

<pre>
Request
TYPE - POST, MULTIPART/FORM
input keys - file<br/>
Response
Message - Success
</pre>

**2. Upload Images**

<pre>
Request
TYPE - POST, MULTIPART/FORM
input keys - seal, logo<br/>
Response
logoImage - file name for logo
sealImage - file name for seal
</pre>

**3. Request body with data to print to template**

<pre>
Request
TYPE - APPLICATION/JSON
input keys - template, definitions, values<br/>
Response
PDF File
</pre>

### CURL example for all APIs

**1. uploadTemplate**

`curl --location --request POST 'localhost:3000/uploadTemplate' \ --form 'file=@"PATH_TO_PDF/pdffile.pdf"'`

**2. uploadImages**

`curl --location --request POST 'localhost:3000/uploadImages' \ --form 'seal=@"PATH_TO_SEAL_IMAGE/seal.png"' \ --form 'logo=@"PATH_TO_LOGO_IMAGE/logo.png"'`

**3. addDataToTemplate**

`curl --location --request POST 'localhost:3000/addDataToTemplate' \ --header 'Content-Type: application/json' \ --data-raw '{ "template":"PDF_NAME_USED_DURING_UPLOAD_TEMPLATE_API.PDF", "definitions":{ "texts": [ {"fieldName": "ClientName","x": 20,"y": 20,"size": 19,"pageNo":1}, {"fieldName": "AmountBilled","x": 30,"y": 30,"pageNo":1}, {"fieldName": "PaymentDueDate","x": 40,"y": 40,"pageNo":1} ], "images": [ {"name": "seal","x": 10,"y": 10,"width": 100,"height": 100,"pageNo":1}, {"name": "logo","x": 50,"y": 50,"width": 100,"height": 100,"pageNo":1} ], "details":{ "increment":5, "pageNo":1, "schema":[ {"fieldName":"name","x":23,"y":146}, {"fieldName":"quantity","x":140,"y":146}, {"fieldName":"price","x":150,"y":146} ] } }, "values":{ "items": [ {"fieldName": "ClientName", "value": "株式会社 Jackpod"}, {"fieldName": "AmountBilled", "value": "¥100,000"}, {"fieldName": "PaymentDueDate", "value": "2022年10月30日"} ], "images": [ {"name": "seal", "objectKey": "KEY_RECIEVED_BY_UPLOAD_IMAGE_API"}, {"name": "logo", "objectKey": "KEY_RECIEVED_BY_UPLOAD_IMAGE_API"} ], "details":[ [ {"fieldName":"name","value":"p1"}, {"fieldName":"quantity","value":"1"}, {"fieldName":"price","value":"100"} ], [ {"fieldName":"name","value":"p2"}, {"fieldName":"quantity","value":"2"}, {"fieldName":"price","value":"200"} ], [ {"fieldName":"name","value":"p3"}, {"fieldName":"quantity","value":"3"}, {"fieldName":"price","value":"300"} ] ] } }'`

### JSON request for addDataToTemplate

<pre>
{
    "template":"PDF_NAME_USED_DURING_UPLOAD_TEMPLATE_API.PDF",
    "definitions":{
        "texts": [
            {"fieldName": "ClientName","x": 20,"y": 20,"size": 19,"pageNo":1},
            {"fieldName": "AmountBilled","x": 30,"y": 30,"pageNo":1},
            {"fieldName": "PaymentDueDate","x": 40,"y": 40,"pageNo":1}
        ],
        "images": [
            {"name": "seal","x": 10,"y": 10,"width": 100,"height": 100,"pageNo":1},
            {"name": "logo","x": 50,"y": 50,"width": 100,"height": 100,"pageNo":1}
        ],
        "details":{
            "increment":5,
            "pageNo":1,
            "schema":[
                {"fieldName":"name","x":23,"y":146},
                {"fieldName":"quantity","x":140,"y":146},
                {"fieldName":"price","x":150,"y":146}
            ]
        }
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
    ],
    "details":[
        [
            {"fieldName":"name","value":"p1"},
            {"fieldName":"quantity","value":"1"},
            {"fieldName":"price","value":"100"}            
        ],
        [
            {"fieldName":"name","value":"p2"},
            {"fieldName":"quantity","value":"2"},
            {"fieldName":"price","value":"200"} 
        ],
        [
            {"fieldName":"name","value":"p3"},
            {"fieldName":"quantity","value":"3"},
            {"fieldName":"price","value":"300"} 
        ]
    ]
    }
}
</pre>

### **Things to remember**

**_1. Please do not forget to set enviroment variables and set correct names of template and images names._**<br/>
**_2. Image names/objectKeys for images will be received via uploadImages API._**<br/>
**_3. Images should be in PNG format._**<br/>
**_4. "seal" in addDataToTemplate is the registered seal and "logo" is company logo. No other name will work._**<br/>
**_5. "fieldName" can be anything but make sure to that it remains same in both "definitions" and "values" or it wont be printed on template._**
