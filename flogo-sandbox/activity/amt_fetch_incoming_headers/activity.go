package amt_fetch_incoming_headers

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"fmt"
    "log"
    "net/http"
)


log = logger.GetLogger("amt_fetch_incoming_headers")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

 	http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return true, nil
}


func handler(w http.ResponseWriter, inputHeaders *http.Request) {


    fmt.Fprintf(w, "%s %s %s \n", inputHeaders.Method, inputHeaders.URL, inputHeaders.Proto)

    //Iterate over all header fields
    for k, v := range inputHeaders.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }

    fmt.Fprintf(w, "Host = %q\n", inputHeaders.Host)
    fmt.Fprintf(w, "RemoteAddr= %q\n", inputHeaders.RemoteAddr)
    //Get value for a specified token
    fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", inputHeaders.Header["Accept"])
}
