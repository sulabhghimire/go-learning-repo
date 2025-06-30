# Technical Notes & Cheatsheet

This document contains a collection of technical notes and commands for common development tasks, including creating self-signed SSL certificates and understanding modern web connection protocols.

## Creating a Test SSL Certificate

A self-signed SSL certificate is useful for local development and testing environments where a certificate from a trusted Certificate Authority (CA) is not required. You can generate one easily using `openssl`.

### Generation Command

Run the following command in your terminal. It will prompt you for information to include in the certificate, such as country, organization name, and common name (e.g., `localhost`).

```sh
openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
```

After filling out the requested information, this command will create two files in your current directory:

- `key.pem`: The private key for the certificate. **Keep this file secure and do not share it.**
- `cert.pem`: The public certificate, which you can share and configure in your web server.

> **Note:** Browsers will show a security warning for self-signed certificates because they are not signed by a trusted CA. For local development, you can typically choose to "proceed anyway."

---

### Understanding the `.pem` File Format

The `.pem` (Privacy Enhanced Mail) format is a common standard for storing and transferring cryptographic data.

- It is a Base64 encoded representation of a **DER** (Distinguished Encoding Rules) certificate.
- **DER** is a binary format for encoding data structures, while **PEM** wraps this binary data in a human-readable, text-based format.

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

---

## How Connections are Handled in HTTP/1.1

HTTP/1.1 introduced significant improvements over its predecessor, primarily through the optimization of connection management.

### Connection Behavior

#### 1. Persistent Connections (Keep-Alive)

By default, HTTP/1.1 uses **persistent connections**. This means that after the initial TCP handshake, a single TCP connection remains open to handle multiple requests and responses between a client and a server. This is a major improvement over HTTP/1.0, where a new connection was often required for each request.

#### 2. Connection Reuse and Pipelining

- **Reuse**: The same TCP connection is reused for subsequent requests, which avoids the overhead (CPU time and network round-trips) of establishing a new connection every time.
- **Pipelining**: A client can send multiple requests over the same connection without waiting for each response. However, the server must send responses back in the same order the requests were received. This can lead to **Head-of-Line (HOL) blocking**, where a slow response holds up all subsequent responses. Due to its complexity and limited benefits, pipelining is rarely enabled in modern browsers.

#### 3. Connection Closure

Either the client or the server can close a connection. The `Connection: close` header is used to signal that the connection will be closed after the current transaction is complete.

### Performance Considerations

- **Latency**: Persistent connections significantly reduce latency for fetching multiple resources (e.g., a web page and its CSS, JS, and image files) by eliminating repeated TCP handshakes.
- **Resource Consumption**: Reusing connections saves CPU and memory on both the client and server.

---

## How Connections are Handled in HTTPS (HTTP/1.1 over TLS)

HTTPS is not a separate protocol; it is simply the HTTP/1.1 protocol layered on top of a secure TLS (or legacy SSL) connection. It inherits all the behaviors of HTTP/1.1 but adds a layer of security.

### Connection Behavior

#### 1. The TLS Handshake

Before any HTTP data can be exchanged, the client and server must perform a **TLS handshake**. This initial process establishes a secure, encrypted channel. The handshake adds latency to the _first_ request on a new connection.

#### 2. Persistent and Secure Connections

Once the secure channel is established, the connection behaves just like an HTTP/1.1 persistent connection. Multiple HTTP requests and responses can be sent over this single, encrypted TCP connection.

#### 3. Connection Resumption

To mitigate the latency of the initial handshake, TLS offers **session resumption**. The client and server can use a previously negotiated session (via session IDs or session tickets) to establish a new secure connection much faster, bypassing the most computationally expensive parts of the full handshake.

### Performance Considerations

- **Initial Latency**: The primary performance cost of HTTPS is the initial TLS handshake, which adds several round-trips to the first request.
- **CPU Overhead**: Encrypting and decrypting data consumes more CPU resources on both the client and server compared to unencrypted HTTP.
- **Overall Benefit**: For modern hardware, the security benefits of HTTPS far outweigh the minor performance overhead, especially with optimizations like session resumption.

---

## How Connections are Handled in HTTP/2

HTTP/2 was designed to address the performance limitations of HTTP/1.1, particularly Head-of-Line blocking.

### Connection Behavior

#### 1. Single TCP Connection

A browser opens just **one TCP connection** per origin (e.g., `www.example.com`) and uses it for all communication.

#### 2. Multiplexing via Streams

HTTP/2 breaks down requests and responses into smaller, independent binary frames. These frames are tagged with a **stream ID** and can be interleaved (multiplexed) over the single TCP connection. The client and server can then reassemble these frames into complete messages.

- **Solves HOL Blocking**: Since frames from different streams are interleaved, a slow response for one resource (e.g., a large image) no longer blocks the delivery of other resources (e.g., a CSS file).

#### 3. Stream Prioritization

Clients can assign a priority to streams, allowing the server to allocate resources to more important requests first (e.g., sending CSS and HTML before large images).

#### 4. Header Compression (HPACK)

HTTP/2 uses HPACK compression to drastically reduce the size of redundant HTTP headers, saving bandwidth and reducing latency.

### Performance Considerations

- **Reduced Latency**: Multiplexing eliminates HOL blocking and allows for parallel request/response handling over a single connection, dramatically speeding up page load times.
- **Lower Resource Usage**: Using one connection per origin reduces memory and CPU usage on both the client and server compared to managing multiple HTTP/1.1 connections.

---

## The TLS Handshake Explained

The Transport Layer Security (TLS) handshake is a process that establishes a secure connection between a client and a server. Here is a simplified step-by-step overview:

1.  **Client Hello**: The client initiates the handshake by sending a `ClientHello` message. This includes:

    - The TLS versions it supports (e.g., 1.2, 1.3).
    - A list of supported **cipher suites** (encryption algorithms).
    - A random string of bytes (Client Random).

2.  **Server Hello**: The server responds with a `ServerHello` message, which includes:

    - The chosen TLS version from the client's list.
    - The chosen cipher suite from the client's list.
    - Another random string of bytes (Server Random).

3.  **Certificate Exchange**: The server sends its digital certificate to the client. This certificate contains the server's public key and has been signed by a trusted Certificate Authority (CA).

4.  **Server Key Exchange & Server Hello Done**: The server may send additional information needed for key generation and then signals it is done.

5.  **Client Verification and Key Exchange**:

    - The client validates the server's certificate by checking its signature against its list of trusted CAs.
    - The client generates a "pre-master secret" (a third random value), encrypts it with the server's public key (from the certificate), and sends it to the server.

6.  **Session Keys Generated**: Both the client and the server now possess the same three pieces of information (Client Random, Server Random, and Pre-Master Secret). They independently use these to generate a set of identical **session keys**. These keys are symmetric (the same key is used for encryption and decryption) and will be used to encrypt all application data for the rest of the session.

7.  **Finished**: The client and server each send a `Finished` message, encrypted with the newly created session key. This confirms that the handshake was successful. If both parties can successfully decrypt the other's message, the secure channel is established.

---

## SSL vs. TLS: Understanding the Difference

Both SSL (Secure Sockets Layer) and TLS (Transport Layer Security) are cryptographic protocols designed to provide secure communication over a network. The key takeaway is that **TLS is the modern, secure successor to the outdated and insecure SSL**.

### Evolution, Not Competition

- **SSL (Secure Sockets Layer)** was the original protocol developed by Netscape.
  - **SSL 2.0 (1995)**: Had critical security flaws.
  - **SSL 3.0 (1996)**: An improvement, but also found to have significant vulnerabilities (e.g., POODLE). **All versions of SSL are now considered deprecated and insecure.**
- **TLS (Transport Layer Security)** was introduced as the upgrade to SSL 3.0.
  - **TLS 1.0 (1999)** & **TLS 1.1 (2006)**: Early versions that are now also deprecated.
  - **TLS 1.2 (2008)**: A major security overhaul that became the long-standing standard for web security.
  - **TLS 1.3 (2018)**: The current standard, offering superior security and improved performance (e.g., a faster handshake).

### Key Differences Summarized

| Feature        | SSL (Legacy)                                      | TLS (Modern)                                                            |
| :------------- | :------------------------------------------------ | :---------------------------------------------------------------------- |
| **Status**     | Deprecated and insecure.                          | Active standard, with TLS 1.2 and 1.3 in wide use.                      |
| **Security**   | Contains known vulnerabilities (POODLE, DROWN).   | Uses stronger encryption algorithms and is more resilient to attacks.   |
| **Handshake**  | Less efficient and secure.                        | TLS 1.3 significantly streamlined the handshake for better performance. |
| **Common Use** | Should **not** be used in any modern application. | The standard for securing web traffic (HTTPS), email, VPNs, and more.   |

**Why do people still say "SSL Certificate"?**
The term "SSL" became widely known first. Even though the underlying protocol is now TLS, the name "SSL" stuck in popular and marketing language. When you buy an "SSL Certificate" today, you are actually buying a certificate that will be used with the modern TLS protocol.
