# goap-editor
A visual editor that generates and tests GOAP actions.

# Building
The client and server need to be build to run the editor.

## Client Build
To build the client navigate to the client folder:

cd client

Then initialise the client:

npm i

Then build the client:

npm run build

## Server Build
To build the server navigate to the root folder, if not already there:

cd ..

To get the dependencies run these commands:

go get -u github.com/gin-gonic/gin
go get -u github.com/gin-gonic/contrib/static

Then build the server:

TBD

# F.A.Q.
Answers to some common questions I ask myself while creating and updating the application:

## Why was Go-Lang chosen for the service?
I wanted to impress some one by learning Go and then show what I have learned.

## Did you impress that somebody?
¯\_(ツ)_/¯ Probably not.

## Did you learn something by using Go?
YES! I really like Go and possibly should have looked into it sooner.

## Will you use Go in all future server applications?
Maybe. I will look at using tech as needed.

## Why did you choose React and Typescript for the front end?
I didnt. AI chose that on my behalf. As Go, and hence the service, was the part of the application that I wanted to gain experience on I let AI generate the front end.

## So, you dont know React and Typescript?
Thats not what I said. I used AI to generate the front end because it was not the focus of my attention.

## Then, you do know React and Typescript?
Ugh! All the questions! Yes, I do, I guess. They have been part of my day job, however much credability that gives.

## Will you set up a Patreon so that people can help support this project?
What people? lol

## Why are you cearing this F.A.Q. and then getting 'snippy' with the answers? Do you expect any one besides you to read them?
Hmm... Yea, well, if you knew me then you would know... which I guess you do... because you, the reader, is me... my only reader... *sad face*

# TODO
A list of items that will improve the application:

## Single Build Command
Create a single build command that will build the front end as well as the back end.

## Security
Add login with authentication and authorisation using JWT tokens.

## Database Atomicity
Move functionality from the service into the database with ACID properties.

## Refactor
Look reducing some duplication of functions by weighing pros and cons.

## Magic Numbers
Look at moving the error message string into a dictionary that can centrally be referenced.

## Location Reference
The location in action and agent is not 'referenced' to the locations table. As part of data base atomicity this can be improved.