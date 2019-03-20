package amt_compute_sha_256_sig

import (
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"crypto/sha256"
	"fmt"
	"time"
	"strconv"

)

var log = logger.GetLogger("AMT Compute SHA56 Sig")

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

	sig := sha256.New()
	//var utc = time.Now().Unix()
	sig.Write([]byte("pvz3r3qgafb6qcaapgjt68nj" + "vNubFQXk7r" +strconv.FormatInt(time.Now().Unix(),10) ))
	fmt.Printf("%x", sig.Sum(nil))

	return true, nil
}

