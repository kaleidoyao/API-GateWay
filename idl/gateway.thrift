include "calculate.thrift"
include "greeting.thrift"
include "reverse.thrift"

service GatewayService {
    calculate.CalculateResponse calculateMethod(1: calculate.CalculateRequest request) (api.post="/calculate");
    greeting.GreetingResponse greetingMethod(1: greeting.GreetingRequest request) (api.get="/greeting")
    reverse.ReverseResponse reverseMethod(1: reverse.ReverseRequest request) (api.post="/reverse")
}