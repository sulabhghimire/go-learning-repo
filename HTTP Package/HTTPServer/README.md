# Understanding Computer Ports

This document provides a basic overview of computer networking ports, their ranges, and some common examples.

## What is a Port?

Think of a computer's IP address as the street address of an apartment building. If you want to send a letter (data) to a specific person (application) in that building, you need to know their apartment number. In networking, that apartment number is the **port number**.

A port is a virtual point where network connections start and end. They are managed by the computer's operating system, and each port is associated with a specific process or service. This allows a single computer to handle many different types of network traffic simultaneously (e.g., browsing a website, receiving email, and transferring a file at the same time).

A computer has a total of 65,536 possible ports, numbered from 0 to 65,535. They are divided into three main ranges.

## Port Ranges

The Internet Assigned Numbers Authority (IANA) has categorized ports into three ranges to help organize services and prevent conflicts.

| Range         | Name                      | Description                                                                                                                                                                                                  |
| ------------- | ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| 0 - 1023      | **Well-known Ports**      | These are reserved for common services and system-level processes (e.g., HTTP for web traffic, SMTP for email). On most OSs, administrative privileges are required to run a service on these ports.         |
| 1024 - 49151  | **Registered Ports**      | These ports are for applications and services that are registered with IANA. They are for specific applications (e.g., databases like PostgreSQL on port 5432) but are less universal than well-known ports. |
| 49152 - 65535 | **Dynamic/Private Ports** | These are used for temporary, private, or custom services. Your web browser uses a port from this range for its side of a connection when you visit a website.                                               |

## Common Ports and Their Uses

Here are some of the most frequently encountered ports.

### Standard Service Ports

These ports are typically associated with standardized internet protocols.

| Port    | Protocol | Common Use                    | Notes                                                                                      |
| :------ | :------- | :---------------------------- | :----------------------------------------------------------------------------------------- |
| **21**  | FTP      | File Transfer Protocol        | Used for transferring files. It is unencrypted and largely considered insecure.            |
| **25**  | SMTP     | Simple Mail Transfer Protocol | Used for routing email between mail servers.                                               |
| **80**  | HTTP     | HyperText Transfer Protocol   | The foundation of the World Wide Web. This is for unencrypted web traffic.                 |
| **443** | HTTPS    | HTTP Secure                   | The encrypted and secure version of HTTP. This is the modern standard for all web traffic. |

### Common Development Ports

These ports are not officially assigned by IANA but have become common conventions for local development environments. When you run a web application on your own machine, it often uses one of these ports by default.

| Port     | Common Use              | Notes                                                                                                                            |
| :------- | :---------------------- | :------------------------------------------------------------------------------------------------------------------------------- |
| **3000** | Web Development Server  | Often the default for modern web frameworks like React (Create React App), and Express.js.                                       |
| **8000** | Web Development Server  | A common alternative, used by Python's simple HTTP server, Django, and other frameworks.                                         |
| **8080** | Web Server / HTTP Proxy | A very common alternative to port 80, often used for application servers (like Apache Tomcat) or as a secondary web server port. |
