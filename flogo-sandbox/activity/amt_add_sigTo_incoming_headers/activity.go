package amt_add_sigTo_incoming_headers

import (
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
    "net/http"
)

var log = logger.GetLogger("amt_fetch_incoming_headers")

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

func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
    http.HandleFunc("/", handler)
  //  log.Error(http.ListenAndServe("localhost:8000", nil))
	return true, nil
}



func handler(w http.ResponseWriter, r *http.Request) {

	//r.Header.Add("x-signature", "actual sig on runtime")

    fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)

    //Iterate over all header fields
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
    }

    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
    //Get value for a specified token
    fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])
}

