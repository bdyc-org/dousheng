namespace go video

struct User {
    1:i64 user_id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct Video {
    1:i64 id
    2:string title
    3:string play_url
    4:string cover_url
    5:User author
    6:i64 favorite_count
    7:i64 comment_count
    8:bool is_favorite
}

struct douyin_feed_request {
    1:optional i64 latest_time
    2:optional string token
}

struct douyin_feed_response {
    1:i32 status_code
    2:optional string status_msg
    3:list<Video> video_list
    4:optional i64 next_time
}

struct douyin_publish_action_request {
    1:string file_name
    2:string token
    3:string title
}

struct douyin_publish_action_response {
    1:i32 status_code
    2:optional string status_msg
}

struct douyin_publish_list_request {
    1:i64 user_id
    2:string token
}

struct douyin_publish_list_response {
    1:i32 status_code
    2:optional string status_msg
    3:list<Video> video_list
}

service VideoService {
    douyin_feed_response FeedVideo(1:douyin_feed_request req)
    douyin_publish_action_response PublishAction(1:douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1:douyin_publish_list_request req)
}