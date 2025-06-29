# Technical Notes & Cheatsheet

This document contains a collection of technical notes and commands for common development tasks, including creating self-signed SSL certificates and understanding HTTP/1.1 connection handling.

## Creating a Test SSL Certificate

A self-signed SSL certificate is useful for local development and testing environments. You can generate one easily using `openssl`.

### Generation Command

Run the following command in your terminal. It will prompt you for information to include in the certificate, such as country, organization name, and common name (e.g., `localhost`).

```sh
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
```

After filling out the requested information, this command will create two files in your current directory:

- `key.pem`: The private key for the certificate. **Keep this file secure.**
- `cert.pem`: The public certificate.

---

### Understanding the `.pem` File Format

The `.pem` (Privacy Enhanced Mail) format is a common standard for storing and transferring cryptographic data.

- It is a Base64 encoded representation of a **DER** (Distinguished Encoding Rules) certificate.
- **DER** is a binary format for encoding data structures, while **PEM** wraps this binary data in a text-based format.

#### Key Characteristics of `.pem` Files

- **ASCII Text Format**: Because they are Base64 encoded, `.pem` files are plain ASCII text. This makes them easy to read, copy, and transport through text-based protocols like email without corruption.
- **Headers and Footers**: `.pem` files use distinct headers and footers to clearly identify the type of content they contain. For example:
  ```
  -----BEGIN CERTIFICATE-----
  (Base64 encoded data...)
  -----END CERTIFICATE-----
  ```
  and
  ```
  -----BEGIN PRIVATE KEY-----
  (Base64 encoded data...)
  -----END PRIVATE KEY-----
  ```

## How Connections are Handled in HTTP/1.1

HTTP/1.1 introduced significant improvements over its predecessor, primarily through the optimization of connection management.

### Connection Behavior

#### 1. Persistent Connections (Keep-Alive)

By default, HTTP/1.1 uses a model of **persistent connections**. This means that a single TCP connection can remain open to handle multiple requests and responses between a client and a server. This is a major improvement over HTTP/1.0, where a new connection was often required for each request.

#### 2. Connection Reuse

Persistent connections allow the client and server to reuse the same TCP connection for subsequent requests. This has several benefits:

- It avoids the overhead (CPU time and network round-trips) of the TCP handshake for every single request.
- A client can send multiple requests over the same connection without waiting for a response to the previous one (a feature known as pipelining, though its practical use is limited).

#### 3. Connection Closure

Either the client or the server can decide to close a connection. The server typically signals its intent to close the connection by including the `Connection: close` header in its response.

### Performance Considerations

The connection handling model of HTTP/1.1 directly impacts performance:

- **Latency**: By eliminating the need for repeated TCP handshakes, persistent connections significantly reduce the latency for fetching multiple resources from the same server (e.g., a web page and its associated CSS, JavaScript, and image files).

- **Resource Consumption**: Reusing connections saves CPU and memory on both the client and the server, as fewer resources are needed to manage a smaller number of active connections.
