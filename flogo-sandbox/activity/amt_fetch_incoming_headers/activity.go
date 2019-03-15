package amt_fetch_incoming_headers

import (
	"bytes"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"reflect"
	"strings"
)

const (
	HeadersPart1      = "IncomingHeaders"
	InputHeaders      = "headers_in"
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

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	key := "any"

			rawHeadersIn := context.GetInput(InputHeaders)
			log.Info("Received these headers: ", rawHeadersIn)
			//fmt.Println(rawHeadersIn)
			if rawHeadersIn != nil && reflect.ValueOf(rawHeadersIn).Kind() == reflect.Map {

				headersMap := make(map[string]string)

				// Convert the headers to forced-lowercase.
				strRawHeadersIn := rawHeadersIn.(map[string]string)
				for key, val := range strRawHeadersIn {
					headersMap[strings.ToLower(key)] = val
				}

				log.Info("Converted array of headers:")
				log.Info(headersMap)

				var sb bytes.Buffer

				for _, headerKey := range headerArr {
					//fmt.Println(headerKey)

					passedHeader, present := headersMap[strings.ToLower(headerKey.(string))]
					//fmt.Println(passedHeader)

					var delta string

					if present {
						delta = passedHeader
						log.Info("Found required header ", headerKey, " = ", delta)
					} else {
						log.Info("Required header ", headerKey, " was not passed in this request.")
						delta = "*"
					}

					sb.WriteString("/")
					sb.WriteString(delta)
				}

				key = sb.String()
			}
		
	log.Info("key:"+key)
	//context.SetOutput(HeadersPart1, key)

	return true, nil
}
