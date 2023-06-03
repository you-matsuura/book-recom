CREATE DATABASE IF NOT EXISTS book_recom;
USE book_recom;

CREATE TABLE IF NOT EXISTS mst_user (
  msu_user_id int(10) AUTO_INCREMENT PRIMARY KEY COMMENT 'ユーザーID',
  msu_nick_name varchar(50) NULL COMMENT 'ニックネーム',
  msu_email varchar(255) NULL UNIQUE COMMENT 'メールアドレス',
  msu_hash_password varchar(255) COMMENT 'パスワードハッシュ値',
  msu_created_at datetime NOT NULL COMMENT '作成日時',
  msu_updated_at datetime NOT NULL COMMENT '更新日時'
);

CREATE TABLE IF NOT EXISTS trx_user_login_cookie (
  tul_user_login_cookie_id int(10) AUTO_INCREMENT PRIMARY KEY COMMENT 'ログインクッキーID',
  tul_user_id int(10) NOT NULL COMMENT 'ユーザーID',
  tul_login_cookie_value varchar(255) NOT NULL UNIQUE COMMENT 'ログインクッキー値',
  tul_hash_password varchar(255) NOT NULL COMMENT 'パスワードハッシュ値',
  tul_expiration_date date NOT NULL COMMENT 'クッキー有効期限',
  tul_created_at datetime NOT NULL COMMENT '作成日時',
  tul_updated_at datetime NOT NULL COMMENT '更新日時'
);

CREATE TABLE IF NOT EXISTS mst_book (
  mbo_book_id int(10) AUTO_INCREMENT PRIMARY KEY COMMENT 'BOOK ID',
  mbo_google_books_id varchar(100) COMMENT 'GoogleBooksAPIから取得したユニークなID',
  mbo_title varchar(255) NOT NULL COMMENT '書籍のタイトル',
  mbo_description text NULL COMMENT '書籍の説明',
  mbo_author varchar(255) NULL COMMENT '著者名',
  mbo_publisher varchar(255) NULL COMMENT '出版社',
  mbo_published_date date NULL COMMENT '出版日',
  mbo_image_url varchar(2083) NULL COMMENT '書籍の画像URL',
  mbo_created_at datetime NOT NULL COMMENT '作成日時',
  mbo_updated_at datetime NOT NULL COMMENT '更新日時'
);

CREATE TABLE IF NOT EXISTS trx_book_post (
  tpo_book_post_id int(10) AUTO_INCREMENT PRIMARY KEY COMMENT '投稿本ID',
  tpo_user_id int(10) NOT NULL COMMENT 'ユーザーID',
  tpo_google_books_id varchar(100) NOT NULL COMMENT 'GoogleBooksAPIから取得したユニークなID',
  tpo_poster_nickname varchar(50) NOT NULL COMMENT '投稿者のニックネーム',
  tpo_comment text NOT NULL COMMENT 'コメント',
  tpo_recommendation tinyint NOT NULL COMMENT 'おすすめ度',
  tpo_post_datetime datetime NOT NULL COMMENT '投稿日時',
  tpo_created_at datetime NOT NULL COMMENT '作成日時',
  tpo_updated_at datetime NOT NULL COMMENT '更新日時'
);