## GO_PDF_generator_Template_1

<br>
This repo will generate a bill with the data provided by API.
Data provided should be in multipart/formdata.

Models

model.PdfData
{
    "name"                    :    string,
    "phoneNumber"      :    int,
    "address"                :    string,
    "zipAddress"            :   int,
    "products"                :    [ ]model.ProductData,
    "logo"                       :    multipart.File
}

model.ProductData
{
    "productName"    :    string,
    "price"                  :    int,
    "quantity"             :    int
}

<br>
API -  /addData  [POST]

Input params :-

| Key         | Value data type       |
| ----------- | --------------------- |
| name        | string                |
| phoneNumber | int                   |
| zipAddress  | int                   |
| address     | string                |
| product     | [ ] model.ProductData |
| logo        | png                   |

<br>
<br>
