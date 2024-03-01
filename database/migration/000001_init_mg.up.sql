-- create file
-- migrate create -ext sql -dir database/migration/ -seq init_mg


-- generate all 
--  migrate -source file://migration -database postgres://postgres:postgres@localhost:5432/satudikti?sslmode=disable up 

-- up version
--  migrate -source file://migration -database postgres://postgres:postgres@localhost:5432/satudikti?sslmode=disable up 2

--down version
--  migrate -source file://migration -database postgres://postgres:postgres@localhost:5432/satudikti?sslmode=disable down 

-- change schema
-- set search_path to beasiswa