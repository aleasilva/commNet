
# commNet

    CommNet is a application to recover the source and the content of distress call message. 

    This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.

## Running Locally

    Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

    ```sh
    $ git clone https://github.com/heroku/go-getting-started.git
    $ cd go-getting-started
    $ go build -o bin/go-getting-started -v 
    $ heroku local
    ```
    Your app should now be running on [localhost:5000](http://localhost:5000/).

## Running from Heroku
    This instrucions is for running the application direct from heroku.
    You can test application using curl or graphic interface like [postman](https://www.postman.com/)

### Check that the server is running:
    First step is certify that app is running for this call [ping](https://comm-net.herokuapp.com/ping).
    After few seconds the server will response a json like this: {"message": "pong"}

### Retrieving a message:
    After checkin tha app is runnig is time to recover the messa sent, for this you can use [topSecret](https://comm-net.herokuapp.com/topsecret)
 
    This shoud be a post message, don´t forget to set the Header Content-Type = application/json.
    The protocol for this comm should be like something like this:
    Request:    
            {
            "satellites":[ 
                {
                    "name": "kenobi",
                    "distance": 100.0,
                    "message": ["este", "", "", "mensaje", ""]                       
                },
                {
                    "name": "skywalker",  
                    "distance": 115.5,
                    "message": ["", "es", "", "", "secreto"]
                },
                {
                    "name": "sato",  
                    "distance": 142.7,
                    "message": ["este", "", "un", "", ""]
                }
                ]
            }

    After ours system process the instructions you shoud see the answer some like this:
      
      Response:
        {
            "Message": "Message Response",
            "Position": {
                "X": 10,
                "Y": 15
            }
        } 


[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)