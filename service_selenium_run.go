package endly

import (
	"github.com/viant/toolbox/url"
	"fmt"
)

//SeleniumRunRequest represents group of selenium web elements calls
type SeleniumRunRequest struct {
	SessionID string
	Browser        string
	RemoteSelenium *url.Resource //remote selenium resource
	PageURL   string
	Actions   []*WebElementAction
}

//SeleniumRunResponse represents selenium call response
type SeleniumRunResponse struct {
	Data      map[string]*ElementResponse
	SessionID string
}

//SeleniumMethodCall represents selenium call.
type SeleniumMethodCall struct {
	Method     string
	Parameters []interface{}
	Wait       *SeleniumWait
}

//WebElementAction represents various calls on web element
type WebElementAction struct {
	Selector *WebElementSelector
	Calls    []*SeleniumMethodCall
}

//ElementResponse represents web element response
type ElementResponse struct {
	Selector *WebElementSelector
	Data     map[string]string
}


//Validate validates run request.
func (r *SeleniumRunRequest) Validate() error {
	if r.SessionID == ""  {
		if r.RemoteSelenium == nil {
			fmt.Errorf("both SessionID and RemoteSelenium were empty")
		}
		if r.Browser == "" {
			fmt.Errorf("both SessionID and Browser were empty")
		}
	}
	return nil
}

//NewSeleniumMethodCall creates a new method call
func NewSeleniumMethodCall(method string, wait *SeleniumWait, parameters ... interface{}) *SeleniumMethodCall {
	return &SeleniumMethodCall{
		Method:method,
		Wait:wait,
		Parameters:parameters,
	}
}