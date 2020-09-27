# service-monitor

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
To change urls or increment their number open the urls.txt file 
which will be cloned and insert each url in line

Enter back-end directory and copy this command inside the directory path in terminal.<br />
To start the back-end server
```
go run main.go
```
Leave the terminal open to run in background and then<br />
Enter front-end directory and copy this command inside the directory path in another terminal
```
npm install
npm start
```
Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

## Finished 
1-A front-end and back-end application communicates via API.<br />
2-list of URLs (to be monitored) are stored in a text file.<br />
3-The backend of the solution should periodically (every 10 seconds) do a HTTP GET request against those URLs<br />
(information should be recorded with each request).<br />
  4-URL.<br />
  5-Response Code.<br />
  6-Response Time (How much time does it took to complete the request), if a request is timing out, this should be the hard coded time out value.<br />
  7-Time Stamp: What was the time when the request took place.<br />
8-Test the solution with correct URLs, URL that resolve but there is no web server.<br />
9-Make sure that in all cases there is data recorded in a correct and consistent format.<br />
10-Configure the front end page the main one to reload every 30 seconds so it can get the new data.<br />

 
## In progress
1-Inserting timeout but working on handling if the request exceeded this time.
2-Still working on testing the solution with URL that doesn't resolve.
3-The data recorded should not be kept indefinitely, let us assume that we will keep only the last 200 records per URL.



## License

This project is licensed under the Apache License
