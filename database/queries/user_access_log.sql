-- name: ListUserAccessLogByPagination :many
select user_access_log.*, u.name as user_name
from user_access_log
         left join user u on u.id = user_access_log.user_id
limit ? offset ?;