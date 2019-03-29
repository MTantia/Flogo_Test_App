package amt_compute_sha_256_sig

import (

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"crypto/sha256"
	"fmt"
	"time"
	"strconv"

)
const (
	sha256_SignatureOutput = "sha256_SignatureOutput"
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

	sha256_Signature := sha256.New()
	sha256_Signature.Write([]byte("pvz3r3qgafb6qcaapgjt68nj" + "vNubFQXk7r"+strconv.FormatInt(time.Now().Unix(),10)))
	fmt.Printf("%x", sha256_Signature.Sum(nil))
	context.SetOutput("sha256_SignatureOutput", sha256_Signature.Sum(nil));
	fmt.Println(context.GetOutput(sha256_SignatureOutput));
	return true, nil
}

