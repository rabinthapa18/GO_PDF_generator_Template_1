# Please check swagger for models

## GeneratePDF1 is for developing Template 1 from sctatch and does not have any coordinate positioning system

## GeneratePDF is generating PDF by specifying the coordinates.

#### cURL request

`curl --location --request POST 'localhost:3000/addToTemplate' \ --header 'Content-Type: application/json' \ --data-raw '{ "name":{"name":"Grrow","x":20,"y":20,"size":22,"pageNo":1}, "phoneNumber":{"phoneNumber":1234,"x":145,"y":50,"size":22,"pageNo":1}, "zipAddress":{"zipAddress":3421,"x":145,"y":30,"size":22,"pageNo":1}, "address": {"address":"Address goes here","x":145,"y":40,"size":22,"pageNo":1}, "products":[ {"productName":{"name":"chair","x":23,"y":146,"size":12,"pageNo":1},"productQuantity":{"quantity":1,"x":140,"y":146,"size":12,"pageNo":1},"productPrice":{"price":100,"x":150,"y":146,"size":12,"pageNo":1}}, {"productName":{"name":"desk","x":23,"y":152,"size":12,"pageNo":1},"productQuantity":{"quantity":5,"x":140,"y":152,"size":12,"pageNo":1},"productPrice":{"price":200,"x":150,"y":152,"size":12,"pageNo":1}} ], "template":1, "logoData":{"x":23,"y":1,"width":20,"height":20,"pageNo":1}, "sealData":{"x":23,"y":23,"width":20,"height":20,"pageNo":1} }'`
