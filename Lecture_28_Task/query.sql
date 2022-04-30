-- name: ListAllStories :many
SELECT * FROM topstories;

-- name: InsertData :execresult
INSERT INTO topstories (STORY_ID, TITLE, SCORE, URL, TimeStamp) VALUES (?,?,?,?,?);

-- name: DeleteAllRecords :exec
DELETE FROM topstories;

-- name: GetTimeFromDB :one
SELECT DISTINCT TimeStamp FROM topstories LIMIT 1;