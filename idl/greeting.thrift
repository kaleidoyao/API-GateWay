struct GreetingRequest {
    1: string name (api.query='name')
}

struct GreetingResponse {
    1: string ResponseBody
}

service GreetingService {
    GreetingResponse greetingMethod(1: GreetingRequest request) (api.get="/greeting")
}