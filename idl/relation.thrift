namespace go relation

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

struct Rela {
    1:i64 Follow_id
    2:i64 Follower_id
}

struct FollowRequest {
    1:i64 user_id
    2:i64 to_user_id
    3:i64 action_type
}

struct FollowResponse {
    1:BaseResponse base_resp
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

struct QueryUserListRequest {
    1:i64 user_id
    2:list<i64> user_ids
}

struct QueryUserListResponse {
    1:BaseResponse base_resp
    2:list<User> user_list
}

service relationService {
    FollowResponse Follow(1:FollowRequest req)
    QueryFollowResponse QueryFollow(1:QueryFollowRequest req)
    QueryFollowerResponse QueryFollower(1:QueryFollowerRequest req)
    QueryUserListResponse QueryUserList(1:QueryUserListRequest req)
}
