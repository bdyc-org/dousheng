namespace go relation

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
    3:i64 service_time
}

struct QueryFollowRequest {
    1:i64 user_id
}

struct QueryFollowResponse {
    1:BaseResponse base_resp
    2:list<i64> Follow_ids
}

struct QueryFollowerRequest {
    1:i64 user_id
}

struct QueryFollowerResponse {
    1:BaseResponse base_resp
    2:list<i64> Follower_ids
}

service relationService {
    QueryFollowResponse QueryFollow(1:QueryFollowRequest req)
    QueryFollowerResponse QueryFollower(1:QueryFollowerRequest req)
}
