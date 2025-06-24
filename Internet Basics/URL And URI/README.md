# How the Internet Works

> The Internet is a global network of interconnected computers that communicates using standardized protocols.

This document provides a high-level overview of the fundamental concepts and processes that power the web.

## Table of Contents

- [Key Concepts](#key-concepts)
- [A Web Request's Journey](#a-web-requests-journey)
- [The DNS Resolution Process (Step-by-Step)](#the-dns-resolution-process-step-by-step)
- [Network Layers](#network-layers)
- [Understanding URL / URI / URN](#understanding-url--uri--urn)

---

## Key Concepts

Some of the foundational components of the internet.

- **Clients and Servers**

  - **Clients**: Devices that request resources. These are typically the end-user's applications, like a web browser or a mobile app.
  - **Servers**: Powerful computers that store and provide resources (like webpages, images, or data) to clients upon request.

- **Protocols**

  - The set of rules that define how data is formatted and transmitted over the internet. They ensure that different devices can understand each other.
  - _Examples_: `HTTP`, `HTTPS`, `TCP/IP`, `FTP`.

- **IP Addresses (Internet Protocol Address)**

  - A unique numerical label assigned to each device connected to a computer network that uses the Internet Protocol for communication.
  - _Example_: `192.0.2.1` (IPv4) or `2001:0db8:85a3:0000:0000:8a2e:0370:7334` (IPv6).

- **Domain Name System (DNS)**
  - Often called the "phonebook of the internet." It's a hierarchical and decentralized system that translates human-readable domain names (like `www.google.com`) into machine-readable IP addresses.

---

## A Web Request's Journey

When you type a web address into your browser and hit Enter, a multi-step journey begins to fetch and display the webpage.

1.  **Entering a URL**: The user types a URL (e.g., `https://www.example.com`) into the browser's address bar.

2.  **DNS Lookup**: The browser cannot connect to `www.example.com` directly; it needs the server's IP address. It performs a DNS lookup to find it. (See the detailed breakdown below).

3.  **Establishing a TCP Connection**: Once the browser has the IP address, it establishes a reliable connection with the server using the **TCP (Transmission Control Protocol)**. This is commonly known as the **Three-Way Handshake**:

    - **SYN**: The browser sends a `SYN` (synchronize) packet to the server to initiate a connection.
    - **SYN-ACK**: The server responds with a `SYN-ACK` (synchronize-acknowledge) packet to acknowledge the request and establish its own parameters.
    - **ACK**: The browser sends an `ACK` (acknowledge) packet back to the server, confirming the connection is established.
    - This connection is **full-duplex**, meaning data can flow in both directions simultaneously.

4.  **Sending an HTTP Request**: With the connection established, the browser sends an **HTTP (Hypertext Transfer Protocol)** request to the server. This request asks for the specific resource, often the HTML file for the webpage.

5.  **Server Processing and Response**: The server receives the HTTP request, processes it (e.g., finds the requested file, runs a script, queries a database), and sends back an **HTTP Response**. The response includes a status code (e.g., `200 OK`) and the requested content (the HTML, CSS, JavaScript, images, etc.).

6.  **Rendering the Webpage**: The browser receives the HTTP response. It parses the HTML to build the page structure (DOM), processes the CSS for styling (CSSOM), and executes JavaScript for interactivity. It then renders the final, visible webpage for the user.

---

## The DNS Resolution Process (Step-by-Step)

DNS resolution is the process of converting a domain name into an IP address. It follows a chain of command, starting from your local machine and moving up to global servers if needed.

1.  **Browser Cache**: The browser first checks its own cache to see if it has recently visited this domain. If the IP address is found and hasn't expired, the process stops here.

2.  **Operating System (OS) Cache**: If not in the browser cache, the browser asks the OS for the IP address. The OS maintains its own cache (and a `hosts` file) of recent DNS lookups.

3.  **Resolver Server Query**: If the IP is still not found, the request is sent to a **DNS Resolver** (or Recursive Server). This server is typically provided by your Internet Service Provider (ISP). The resolver's job is to find the IP address on your behalf.

4.  **Root Server Query**: The resolver, if it doesn't have the IP cached, queries one of the **Root Name Servers**. There are only 13 logical root servers worldwide (replicated many times). The root server doesn't know the IP address, but it knows where to find the server that handles the Top-Level Domain (TLD), like `.com`, `.org`, or `.net`. It responds with the address of the appropriate TLD server.

5.  **TLD Server Query**: The resolver then queries the **TLD Name Server** (e.g., the `.com` server). This server doesn't have the final IP address either, but it knows which server is the **authoritative name server** for the specific domain (`example.com`). It directs the resolver to that server.

6.  **Authoritative Name Server Query**: Finally, the resolver queries the **Authoritative Name Server**. This is the server that holds the official DNS records for the domain in question. It responds with the final IP address for `www.example.com`.

7.  **Response to Client**: The resolver receives the IP address from the authoritative server. It stores this record in its cache for a certain amount of time (defined by the TTL - Time To Live) and sends the IP address back to your OS, which in turn gives it to the browser. The journey is now complete, and the browser can initiate the TCP connection.

---

## Network Layers

The process of communication over the internet is standardized into a set of layers. Each layer has a specific responsibility and interacts only with the layers directly above and below it. This is a simplified 4-layer model (based on the TCP/IP model).

### Application Layer

- **Protocols**: `HTTP`, `HTTPS`, `DNS`, `FTP`, `SMTP`
- **Responsibility**: Provides high-level APIs for applications to use the network. This is where user-facing protocols operate. It handles resource sharing and remote file access.

### Transport Layer

- **Protocols**: `TCP`, `UDP`
- **Responsibility**: Manages end-to-end communication between the client and server. It ensures data is transmitted reliably (`TCP`) or quickly (`UDP`), handles error checking, and controls data flow.

### Internet Layer

- **Protocols**: `IP` (IPv4, IPv6)
- **Responsibility**: Responsible for logical addressing (IP addresses), routing, and forwarding data packets across different networks to their final destination.

### Link Layer

- **Protocols**: `Ethernet`, `Wi-Fi`
- **Responsibility**: Manages the physical connection between devices on the same local network. It handles physical addressing (MAC addresses) and the actual transmission of bits over a physical medium (like cables or radio waves).

---

## Understanding URL / URI / URN

These terms are often used interchangeably, but they have distinct meanings.

- **URI (Uniform Resource Identifier)**: The most general term. A URI is a string of characters that unambiguously identifies a particular resource. It's a superset of both URLs and URNs.
- **URL (Uniform Resource Locator)**: The most common type of URI. A URL not only identifies a resource but also specifies _how and where to access it_ (the "locator"). Every URL is a URI.
- **URN (Uniform Resource Name)**: A URI that identifies a resource by a persistent, location-independent name, but doesn't specify how to locate it. For example, an `isbn` number for a book is a URN.

**Key takeaway**: All URLs are URIs, but not all URIs are URLs.

### Components of a URL

A URL is composed of several parts that provide the browser with all the information it needs.

`https://www.example.com:443/products/laptops?id=123&sort=price#specs`

- **Scheme**: `https`
  - The protocol to be used for accessing the resource. Common schemes are `http`, `https`, `ftp`, and `mailto`.
- **Host**: `www.example.com`
  - The domain name (or IP address) of the server where the resource is located.
- **Port**: `:443`
  - The specific "gate" on the server to connect to. This is often omitted, in which case a default is used (`80` for `http`, `443` for `https`).
- **Path**: `/products/laptops`
  - The specific path to the resource on the server, similar to a file path on your computer.
- **Query**: `?id=123&sort=price`
  - A set of key-value pairs that send extra information to the server. It starts with a `?`.
- **Fragment**: `#specs`
  - An anchor that directs the browser to a specific part of the page after it has loaded. This is processed client-side and is not sent to the server.
