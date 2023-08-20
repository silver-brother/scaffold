-- name: GetUserById :one
select *
from user
where id = ?
limit 1;

-- name: GetUserByIdCard :one
select *
from user
where id_card = ?;

-- name: ListUserByIds :many
select *
from user
where id in (sqlc.slice('ids'));


-- name: InsertUser :exec
insert into user (name, sex, birth_date, id_card, mobile, avatar, description)
values (?, ?, ?, ?, ?, ?, ?);


