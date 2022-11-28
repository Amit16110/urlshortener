# Url Shortener

A URL Shortener service.

## Installation

You can install it by cloning the repository.

```bash
git clone https://github.com/Amit16110/urlshortener.git
```

### or

## Run

Run with the Docker.

```bash
git clone https://github.com/Amit16110/urlshortener.git

docker build -t <imageName> .
docker run -dp 8080:8080 <imageName>
```

## Usage

Using CURL generates a shortener URL

```
$ curl -H "Content-Type: application/json" -X POST -d '{"name":"Amit",
"email":"amit16110",
"url":"https://github.com/Amit16110/urlshortener"}' http://localhost:8080/urlshort

<!-- Response: -->

<!-- {
    "created_at":"2022-11-28T07:07:55.602977133Z",
    "id":11547430018834754140,
    "name":"Amit",
    "short_url":"http://localhost:8080/MuAUyDY4WVj"
} -->
```

### or

Redirect, open short URL in your browser.

```
http://localhost:8080/MuAUyDY4WVj
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
