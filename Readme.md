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

> You can find index.html inside the HTML folder where you can upload template and write the request as you would write in postman ot swagger (in JSON) and it will render the PDF inside the webpage from which you can download it or print it from there.

_We will have to change API URL in fetch functions inside the HTML and set the config vars for the AWS account._

_**Example shown below**_



https://user-images.githubusercontent.com/27823073/178200294-b19993d5-0377-4947-ba3b-3b8a3b049957.mov

