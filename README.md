# service-monitor

## Project Objectives:
The objective of this project is to create service to allow us to monitor production and pre-production services, this service should have the following features:<br />
1- Monitor the health and availability of other services, example: [https://www.gov.uk.glasswall-icap.com/](https://www.gov.uk.glasswall-icap.com/).<br />
2- It should be able to visualize in a web interface the availability and health of other services.<br />
3- Be able to do a L7 health checks against the protocols of interest: HTTP, HTTPS, ICAP.<br />
4- Verify that files are passing through the CDR engine for a specific website.<br />
5- Sending alerts when a specific service down criteria is met, maybe Slack based alerts, emails, Syslog, SNMP traps or similar.<br />
6- Take corrective actions, as restarting a service or execute a script when a service is down and relevant action is defined.<br />

## Prerequisites

* Apache Server
* Mysql Database
* [Node & npm](https://nodejs.org)
* [Go](https://golang.org)


## Getting Started

## database

For Windows:<br />
Install [XAMPP](https://www.apachefriends.org/index.html) for an easy quickstart<br />
open terminal in path "C:\xampp\mysql" and run
```
bin\mysql -u <username> -p <password>
source <source of healthcheck.sql>
```

for Linux: open terminal and run
```
mysql -u <username> -p <password>
source <source of healthcheck.sql>
```

### After setting up database

For windows: open terminal and copy
```
echo %GOPATH%
```
for Linux: open terminal and copy
```
echo $GOPATH
```
Create in the above path three directories (pkg,src,bin)<br />
inside src directory create directory named github.com<br />

Clone the repo in github.com directory :
```
git clone https://github.com/k8-proxy/service-monitor.git
```
Or clone the develop branch:
```
git clone -b develop https://github.com/k8-proxy/service-monitor.git
```

Enter back-end directory and copy this command inside the directory path in terminal.<br />

```
go get github.com/gorilla/mux
go get github.com/gorilla/handlers
go get github.com/go-sql-driver/mysql
```
To start the back-end server in the same terminal
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
8-Test the solution with correct URLs.<br />
9-Make sure that in all cases there is data recorded in a correct and consistent format.<br />
10-CRUD operations from text file then shifting to CRUD operations using MySQL database.<br />
11-Insert time interval for create and update.<br />
12-Configure the front end page the main one to reload every 30 seconds so it can get the new data.<br />

 
## In progress

1-testing the solution with URL that doesn't resolve the url will be skipped and the healthcheck will continue, still working on inserting their own responses


## License

This project is licensed under the Apache License
