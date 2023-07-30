struct ReverseRequest {
    1: required string inputString;
}

struct ReverseResponse {
    1: required string outputString;
}

service ReverseService {
    ReverseResponse reverseMethod(1: ReverseRequest request) (api.post="/reverse")
}