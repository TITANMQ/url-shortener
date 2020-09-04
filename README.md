# url-shortener

url-shortener is a little go application I made which shortens URLs.

## Prerequisites
- Go lang installed to run the application 
- Code editor is preferred
- An API client to send API requests (I use Postman)

## Installation

1)  Download source files from github.

2)  Add the follow environment variables to the example.env file:
```env
port = 8080
host = localhost
```
Rename file "example.env" to ".env"

3)  Rename file "example.database.db" to "database.db"
## Usage

Run the application in a console:
```go
go run ./main.go
```
Once the application has started on your chosen port, use a API client to send a request to the running application.
The request will be for creating a shortened URL.

Request URL should be a POST request to the following endpoint:
```
POST hostname:port/api/v1/url
```
You should add the following body in JSON format: 
```json
{
    "url": "https://www.google.com/"
}
```

If successful then you should receive a similar JSON response:
```json
{
    "data": {
        "id": 1,
        "url": "https://www.google.com/",
        "short_url": "fT8m94"
    },
    "message": "URL created successfully",
    "status": true
}
```
Finally, take the short_url from the response and go to the browser and enter the following URL:
```
hostname:port/fT8m94
```
> **Hostname** and **port** being the hostname and port specified in the .env file

Now, that should redirect you to the URL you set.

If failed, please could you report any problems or bugs you have as issues. Thanks 

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
