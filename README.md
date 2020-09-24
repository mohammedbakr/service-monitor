# service-monitor

## What is ICAP?

**ICAP** stands for **Internet Content Adaptation Protocol**. If a **content** (for example: file) you've requested over the internet
to download or whatever, needs **adaptation**(some kind of modification or analysis), the proxy server sends the content to the ICAP server for adaptation and after performing the required tasks on the content, the ICAP server sends it back to the proxy server so that it may return the adapted content back to the destination. This can occur both during request and response.

To know more about the ICAP protocol, [check this out](https://tools.ietf.org/html/rfc3507).

## Project Objectives:
The objective of this project is to create service to allow us to monitor production and pre-production services, this service should have the following features:<br />
1- Monitor the health and availability of other services, example: [https://www.gov.uk.glasswall-icap.com/](https://www.gov.uk.glasswall-icap.com/).<br />
2- It should be able to visualize in a web interface the availability and health of other services.<br />
3- Be able to do a L7 health checks against the protocols of interest: HTTP, HTTPS, ICAP.<br />
4- Verify that files are passing through the CDR engine for a specific website.<br />
5- Sending alerts when a specific service down criteria is met, maybe Slack based alerts, emails, Syslog, SNMP traps or similar.<br />
6- Take corrective actions, as restarting a service or execute a script when a service is down and relevant action is defined.<br />

### Prerequisites

* [Node & npm](https://nodejs.org)
* [Go](https://golang.org)

## Getting Started

Clone the repo:
```
git clone https://github.com/k8-proxy/service-monitor.git
```
Or clone the develop branch:
```
git clone -b develop https://github.com/k8-proxy/service-monitor.git
```

Enter back-end directory and copy this command inside the directory path.
To start the back-end server
```
go run main.go
```

Enter front-end directory and copy this command inside the directory path
```
npm install
npm start
```
Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.


## License

This project is licensed under the Apache License
