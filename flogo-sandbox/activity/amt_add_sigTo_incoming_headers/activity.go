package amt_add_sigTo_incoming_headers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"net/http"
	"strconv"
	"time"
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

   // http.HandleFunc("/", handler)
   //log.Error(http.ListenAndServe("localhost:8000", nil))
	return true, nil
}



func handler(w http.ResponseWriter, input_Headers *http.Request) {

	sha256_Signature := sha256.New()
	//test environment details hardcoded
	sha256_Signature.Write([]byte("pvz3r3qgafb6qcaapgjt68nj" + "vNubFQXk7r" +strconv.FormatInt(time.Now().Unix(),10) ))
	fmt.Printf("%x", hex.EncodeToString(sha256_Signature.Sum(nil)))

	input_Headers.Header.Add("x-signature", hex.EncodeToString(sha256_Signature.Sum(nil)))
}

