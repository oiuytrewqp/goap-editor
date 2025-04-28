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
