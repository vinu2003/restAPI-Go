A basic article API written in GO Programming Language
you can perform,
    - Create an article
    - Retrive an article by ID
    - Retirve custom data of articles based on tag of the article and date it is posted.

KINDLY NOTE: Assumptions are highlighted as ***ASSUMPTION***

Directory Structure:
--------------------
awesomeProject/
    |- Godeps/             - Contains info about all dependencies of the project
    |- controller/              - Contains main API logic files
        |- handler.go      - Defines methods handling calls at various endpoints
        |- handler_test.go - DEfines unit tests for all handler functions
        |- model.go        - User and Product models
        |- dbstore.go      - Methods interacting with the database
        |- router.go       - Defines routes and endpoints
    |- errors/
        |- error.go        - implements functions to manipulate errors
    |- README.md
    |- articledatstore.js  - Script to populate local mongodb with dummy data
    |- main.go             - Entry point of the API



SETUP:
------
Local Machine I have used is macbookPro - Version 10.13.6

GoLang installment setup:
go version - go version go1.11.1 darwin/amd64

***ASSUMPTION*** : Go is requested to be the prmary programming language so I beleive the go setup must be available.
IDE used : GoLand
Set GOPATH to your src directory where the source files are placed.

Install MongoDB and start server by typing mongod.
* Reasons for using mongo DB is it is FAST, Web Scale friendly adhoc queries, real-time aggregation, fast parsing stores data as BSON etc..;


Additional Libraries used:
--------------------------
## Library for unit test
$ go get -u "github.com/stretchr/testify"

## Libraries to handle network routing
$ go get -u "github.com/gorilla/mux"
$ go get -u "github.com/gorilla/context"
$ go get -u "github.com/gorilla/handlers"

## mgo library for handling MongoDB
$ go get "gopkg.in/mgo.v2"


Database setup:
---------------
Mongo DB is started using mongod cmd.
logs can viewed same cmd.

Insert dummy data:
Use the script file articledatastore.js
$mongo < articledatstore.js


API DETAILS:
------------
AS mentioned above in 'Directory Structure' section, API entrypoint is in main.go
PORT details can be verified in main.go - if required feel free to change while providing CURL commands to test apis.

***ASSUMPTION***
mongo default port is 27017 - u can refer in dbstore.go file - make sure mongodb started with same port or else connection will fail.
If you are starting mongo with different port please update dbstore.go

As requested, the POST and GET methods (only BACKEND CODE) is implemented and not Frontend UI.
CURL command is used to verify the api in cmdline or safari/chrome installed on laptop can also be used.

***ASSUMPTION***
For POST method, duplicate entries are not inserted.

unit tests:
Implemented simple unit test frame work for validality handlers.
10 test case are implemented:

***ASSUMPTION***
for the above mentioned duplicate entries not allowed to be inserted, rerunning the unit test can have issue with 2 test cases
@@line218 TestHandler_ArticlesHandlerValidInput
@@line239 TestHandler_ArticlesHandlerDuplicateInput
  - what needs to be done ?
      make sure the data variable which has JSON input is slightly modified in ''both cases'' for each run.
      @@line218&240: data := []byte(`{"id":1,"title":"","date":"2018-03-14","body":"Change in climate and vegetation","tags":["world","climate","nature"]}`)

      WORKAROUND is implement DELETE method to lookup for match and delete which is not yet implemented as it is out of scope of this task.

Error handling:
Knowing that go doesnt have rich error handling mechanism, I tried to see how its done in case web based projects.
I prefer not to aggregate error status as one handler(which is most commonly used)  rather each handler have in their own method.
However, I have added custom handling mechanism:
  - this is just to make sure the details are added to error string on from database end as well which would be easy to debug using LOG messages.

Authentication:
username: test and password: password
 - this has to be provied in URL.

Examples:
--------
POST METHOD:
vinodhinis-MBP:awesomeProject vinodhinibalusamy$ curl -u test:password -H "Content-Type: application/json" -X POST -d '{"id":7,"title":"OL","date":"2018-10-05","bodi":"My STROL","tags":["aaa","ooo", "lll"]}' http://localhost:8982/articles
Added the article successfully...

GET method with ID:
vinodhinis-MBP:awesomeProject vinodhinibalusamy$ curl -u test:password -X GET http://localhost:8984/articles/1
{
  	"ID": 1,
  	"title": "latest science show that potato chips are better for you than sugar.",
  	"date": "2016-09-22",
  	"body": "some text, potentially containing simple markup about how potato chips are great.",
  	"tags": [
  		"health",
  		"fitness",
  		"science"
  	]
}

GET method with tag&date:
vinodhinis-MBP:awesomeProject vinodhinibalusamy$ curl -u test:password -X GET http://localhost:8981/tag/aaa/20181005
{
  	"tag": "aaa",
  	"count": 3,
  	"articles": [
  		"6",
  		"5",
  		"4"
  	],
  	"related_tags": [
  		"xxx",
  		"yyy",
  		"zzz",
  		"SSS",
  		"ooo",
  		"lll"
  	]

IMPROVEMENTS : I prefer to implement mock tests so that i can cover some corner cases in db query which I tried and not succesfull.
               I will try that later as side project.


Also, Documents/articles referred : I have signed up for Cloud Native Go with dockers and kubeternes course in linkedIn and followed up most of the concepts.
The reason is i havent written REST api using GoLang but only distributed software like JuJu.
I have developed REST api using SPring boot JAVA. So inorder to learn Web Dev using GO for current market requirement I am pursuing that specific course.

Requesting to provide any suggestions & improvements in the work I have done which will be useful for improvising my skills.

Thanks
-Vinodhini


