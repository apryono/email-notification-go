# Golang Email Notification

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/apryono/email-notification-go)

## Email Notification use net/http

  - Feature to send email notification with [sendgrid/sendgrid-go][sg]
  - Use net/http, bytes
  - Logging handler use [sirupsen/logrus][logrus]

## How to execute ? 

  - Make sure the go language driver is installed in your locale
  - If not already you can click below [Go Language][go]
  - You can use any text editor what you like to use
  - After you installing the go, you can clone this repository
  - before clone this repo, firstly create folder github.com/ in src folder, and then you can git clone at github.com folder.

You need :
Library of Logrus for logging handler or you can use log at go library :

        go get https://github.com/sirupsen/logrus

### After Clone Repository
        cd github.com/email-notification-go

-  after that you can change the token with yours
-  and change the mail sender too
-  and you can modify in your local
- and then run the main.go

        go run main.go

### Installation

Install the dependencies and devDependencies and start the server.

```sh
$ cd github.com/email-notification-go
$ go get https://github.com/sirupsen/logrus
$ go run main.go
```
MIT


**Free Software, Hell Yeah!**


   [sg]: <https://github.com/sendgrid/sendgrid-go>
   [logrus]: <github.com/sirupsen/logrus>
   [go]: <https://golang.org/dl/>
