# chapter 3

- learned about the flag parser,
  // addr := flag.string("addr", ":4000", "Http network address")
  //flag.parse()
  - this will enable us to give the addr from the CLI
    //go run cmd/web/\* -addr=":9999"

# 3.2

Leveled Logging

- infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

- errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

2 types of log streams for seperating the logs from the applciation

#3.3 dependency injection

- Inorder to provide handlers with dependencies we will use dependency injection
- In go we create and applcaition struct and define handler methods against that applcation
  //type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
  }
  //app := &application{
  errorLog: errorLog,
  infoLog: infoLog,
  }

  //func (app *application) home(w http.ResponseWriter, r *http.Request)

#3.4 Centralized Error Handling

- we introduced helper.go where we cleaned up some of the error handling to the helper methods
  -serverError
  -clientError
  -notfound

#3.5 Isolating application routes
The main file should handle

- Parsing the runtime configuration settings for the application;
- Establishing the dependencies for the handlers; and
- Running the HTTP server.

# 4 Database driven responses

- create a db and assign dummy values to it via CLI
  // create a new user and allow access to read and insert only
  -- - CREATE USER 'web'@'localhost';
  -- GRANT SELECT, INSERT ON snippetbox.\* TO 'web'@'localhost';
  -- -- Important: Make sure to swap 'pass' with a password of your own choosing.
  -- ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

- learned about go get // go mod tidy package maangers
- The \_ import runs the driver’s setup without requiring you to reference it in your code, avoiding compiler errors.
