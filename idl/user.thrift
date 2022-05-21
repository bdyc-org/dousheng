namespace go user

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
    3:i64 service_time
}

struct User {
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct CreateUserRequest {
    1:string username
    2:string password
}

struct CreateUserResponse {
    1:BaseResponse base_resp
    2:i64 user_id
}

struct CheckUserRequest {
    1:string username
    2:string password
}

struct CheckUserResponse {
    1:BaseResponse base_resp
    2:i64 user_id
}

struct MGetUserRequest {
    1:i64 user_id
    2:list<i64> user_ids
}

struct MGetUserResponse {
    1:BaseResponse base_resp
    2:list<User> user_list
}

struct NewFollowRequest {
    1:i64 follow_id
    2:i64 follower_id
}

struct NewFollowResponse {
    1:BaseResponse base_resp
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    MGetUserResponse MGetUser(1:MGetUserRequest req)
    NewFollowResponse NewFollow(1:NewFollowerRequest req)
}
