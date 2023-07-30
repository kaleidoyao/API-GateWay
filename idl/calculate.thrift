struct CalculateRequest {
    1: i32 operand_1 (api.body="operand_1")
    2: i32 operand_2 (api.body="operand_2")
}

struct CalculateResponse {
    1: bool success (api.body="success")
    2: string message (api.body="message")
    3: i32 data (api.body="data")
}

service calculateService {
    CalculateResponse calculateMethod(1: CalculateRequest request) (api.post="/calculate")
}