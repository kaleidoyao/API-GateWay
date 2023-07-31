struct CalculateRequest {
    1: i32 operand_1 (api.body="operand_1")
    2: i32 operand_2 (api.body="operand_2")
}

struct CalculateResponse {
    1: i32 outcome (api.body="outcome")
}

service calculateService {
    CalculateResponse calculateMethod(1: CalculateRequest request) (api.post="/calculate")
}