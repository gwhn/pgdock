-- postgres password format md5 -s "{user_name}{password}"
CREATE USER test_user WITH ENCRYPTED PASSWORD 'md57bfb43fbd55dd6ef8af906df058882c2';
CREATE DATABASE test_db;
GRANT ALL PRIVILEGES ON DATABASE test_db TO test_user;
