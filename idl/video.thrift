namespace go video

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

struct FeedRequest {
    1:i64 latest_time
    2:i64 user_id
}

struct FeedResponse {
    1:BaseResponse base_resp
    2:i64 next_time
    3:list<Video> video_list
}

struct CreateVideoRequest {
    1:i64 author_id
    2:string play_url
    3:string cover_url
    4:string title
}

struct CreateVideoResponse {
    1:BaseResponse base_resp
}

struct PublishListRequest {
    1:i64 user_id
    2:i64 author_id
}

struct PublishListResponse {
    1:BaseResponse base_resp
    2:list<Video> video_list
}

struct MGetVideoRequest {
    1:i64 user_id
    2:list<i64> video_ids
}

struct MGetVideoResponse {
    1:BaseResponse base_resp
    2:list<Video> video_list
}

struct FavoriteOperationRequest {
    1:i64 video_id
    2:i64 action_type
}

struct FavoriteOperationResponse {
    1:BaseResponse base_resp
}

struct CommentOperationRequest {
    1:i64 video_id
    2:i64 action_type
}

struct CommentOperationResponse {
    1:BaseResponse base_resp
}

service VideoService {
    FeedResponse Feed(1:FeedRequest req)
    CreateVideoResponse CreateVideo(1:CreateVideoRequest req)
    PublishListResponse PublishList(1:PublishListRequest req)
    MGetVideoResponse MGetVideo(1:MGetVideoRequest req)
    FavoriteOperationResponse Favorite(1:FavoriteOperationRequest req)
    CommentOperationResponse comment(1:CommentOperationRequest req)
}