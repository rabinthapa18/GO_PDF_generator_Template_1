# Please check swagger for models

## GeneratePDF1 is for developing Template 1 from sctatch and does not have any coordinate positioning system

## GeneratePDF is generating PDF by specifying the coordinates.

#### cURL request

##### For Uploading Template

`curl --location --request POST 'localhost:3000/uploadTemplate' \ --form 'file=@"/path/to/file"'`

##### For generating PDF

`curl --location --request POST 'localhost:3000/addToTemplate' \ --header 'Content-Type: application/json' \ --data-raw '{ "name":{"name":"Grrow","x":20,"y":20,"size":22,"pageNo":1}, "phoneNumber":{"phoneNumber":1234,"x":145,"y":50,"size":22,"pageNo":1}, "zipAddress":{"zipAddress":3421,"x":145,"y":30,"size":22,"pageNo":1}, "address": {"address":"Address goes here","x":145,"y":40,"size":22,"pageNo":1}, "products":[ {"productName":{"name":"chair","x":23,"y":146,"size":12,"pageNo":1},"productQuantity":{"quantity":1,"x":140,"y":146,"size":12,"pageNo":1},"productPrice":{"price":100,"x":150,"y":146,"size":12,"pageNo":1}}, {"productName":{"name":"desk","x":23,"y":152,"size":12,"pageNo":1},"productQuantity":{"quantity":5,"x":140,"y":152,"size":12,"pageNo":1},"productPrice":{"price":200,"x":150,"y":152,"size":12,"pageNo":1}} ], "template":"TEMPLATE_NAME.pdf", "logoData":{"x":23,"y":1,"width":20,"height":20,"pageNo":1}, "sealData":{"x":23,"y":23,"width":20,"height":20,"pageNo":1} }'`
<br>

> **Template name should be same as name used during upload PDF api**

> AFTER STARTING THE APP, YOU CAN VISIT "/" IN YOU BROWSER ACCORDING TO YOU SERVER AND USE IT (FOR EXAMPLE : "localhost:3000/"). THIS WILL REDIRECT YOU TO A PAGE WHERE YOU CAN UPLOAD AND WRITE DATA AND PDF WILL OPEN IN NEW PAGE. YOU WILL HAVE TO ADD ENV VARIABLES IN DOCKERFILE

### <span style="color:red">Do not forget to change template name in request textbox. Please input it same as the pdf name that you are uploading.</span>

<br>
_**Example shown below**_

https://user-images.githubusercontent.com/27823073/178200294-b19993d5-0377-4947-ba3b-3b8a3b049957.mov
