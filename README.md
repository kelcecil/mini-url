mini-url
--------

A demonstration url shortening service in Go.

# Using the app

## Building and Running

```
go build -o url-shortener-service *.go
./url-shortener-service
```

Tests are also available and can be ran using:

```
go test -v
```

## Adding a new URL

New urls can be added by sending a PUT to the server's root path with a simple JSON that includes the URL to shorten. For example, the following curl will add the About page of my blog to the store and returned a shortened key for use:

```
curl -XPUT -H "Content-Type: application/json" -d '{"url":"http://kelcecil.com/about"}' localhost:8080
> {"key":"a"}
```

## Using a shortened URL

You can use the key returned in the previous example to now access your submitted page. The user need only send an HTTP GET request using the key as the path. The server looks up the URL using whatever storage is in use (local usage defaults to using a map) and returns a temporary redirect to the new page. Below is an example output using curl to show the redirect:

```
curl -v -XGET http://localhost:8080/a
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /a HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 307 Temporary Redirect
< Location: http://kelcecil.com/about
< Date: Wed, 03 Aug 2016 13:38:39 GMT
< Content-Length: 61
< Content-Type: text/html; charset=utf-8
<
<a href="http://kelcecil.com/about">Temporary Redirect</a>.
```

# The Code

The application begins in the main function (main.go). This function is kept to a minimum and only initializes the storage, creates the http handler, and starts the service.

The mentioned storage mechanism is the MapUrlStorage struct (map_url_storage.go) and takes advantage of Go maps to provide an easy local storage for development. MapUrlStorage satisfies the ShortUrlStorage interface (url_persistence.go) to aid the task of replacing the MapUrlStorage with a storage option in the cloud (AWS RDS or Dynamo, GCP's Cloud SQL, etc).

The ShortUrlForwardingHandler (url_handler.go) is created using the initialized storage and is passed to Go's builtin HTTP server to route and handle requests for creating new and using short URLs.

Tests exist most for pieces of code and are available in the `*_test.go` source files.

# Further Improvements

There's several things that need to be completed for this to be a complete demonstration. This list may change through additional work.

* Better error handling for missing keys - The server currently just refuses the connection if the key doesn't exist. There should be a friendly message at a minimum.
* Better 12-Factor app compliance - Configurations such as selecting the store, server port, etc should be read as environment variables. Necessary for easy cluster deployment.
* Postgres or other persistent URL storage - A map-based data store was first created due to it's simplicity, but other storage mechanisms can be used by satisfying the ShortURLStorage interface in url_persistence.go. This would be a necessity for scaling in the cloud.
* Avoid URL duplications - The storage doesn't currently ensure that duplicate URLs aren't inserted into our storage. The storage should check for duplicate URLs and simply return the existing key for that URL.
* Dockerfile
