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
    6:string avatar
    7:string signature
    8:string background_image
    9:i64 total_favorited
    10:i64 favorite_count
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

struct FollowOperationRequest {
    1:i64 follow_id
    2:i64 follower_id
    3:i64 action_type
}

struct FollowOperationResponse {
    1:BaseResponse base_resp
}

struct AuthenticationRequest {
    1:string username
}

struct AuthenticationResponse {
    1:BaseResponse base_resp
    2:i64 user_id
}

struct FavoriteOperationRequest {
    1:i64 user_id
    2:i64 video_auther
    3:i64 action_type
}

struct FavoriteOperationResponse {
    1:BaseResponse base_resp
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
    MGetUserResponse MGetUser(1:MGetUserRequest req)
    FollowOperationResponse Follow(1:FollowOperationRequest req)
    AuthenticationResponse Authentication(1:AuthenticationRequest req)
    FavoriteOperationResponse Favorite(1:FavoriteOperationRequest req)
}
