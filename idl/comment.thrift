namespace go comment

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
    5:bool if_follow
}

struct Comment {
    1:i64 id
    2:User user
    3:string content
    4:string creat_date
}

struct CommentRequest {
    1:i64 user_id
    2:string token
    3:i64 video_id
    4:i32 action_type
    5:string comment_text
    6:i64 comment_id
}

struct CommentResponse {
    1:BaseResponse base_resp
    2:Comment comment
}

struct QueryCommentRequest {
    1:string token
    2:i64 video_id
}

struct QueryCommentResponse {
    1:BaseResponse base_resp
    2:list<Comment> comment_list
}

service CommentService {
    CommentResponse Comment(1:CommentRequest req)
    QueryCommentResponse QueryComment(1:QueryCommentRequest req)
}
