namespace go favorite

struct BaseResponse {
    1:i64 status_code
    2:string status_msg
    3:string service_time
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

struct Video {
    1:i64 id
    2:User author
    3:string play_url
    4:string cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
    8:string title
}

struct FavoriteOperationRequest {
    1:i64 user_id
    2:i64 video_id
    3:i64 action_type
}

struct FavoriteOperationResponse {
    1:BaseResponse base_resp
}

struct FavoriteListRequest {
    1:i64 user_id
}

struct FavoriteListResponse {
    1:BaseResponse base_resp
    2:list<Video> video_list
}

struct FavoriteJudgeRequest {
    1:i64 user_id
    2:list<i64> video_ids
}

struct FavoriteJudgeResponse {
    1:BaseResponse base_resp
    2:list<i64> video_ids
}

service FavoriteService {
    FavoriteOperationResponse Favorite(1:FavoriteOperationRequest req)
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
    FavoriteJudgeResponse FavoriteJudge(1:FavoriteJudgeRequest req)
}
