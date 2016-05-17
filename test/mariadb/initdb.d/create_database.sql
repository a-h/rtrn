CREATE DATABASE rtrn;

CREATE USER rtrn IDENTIFIED BY 'rtrn';

GRANT ALL PRIVILEGES ON rtrn.* TO rtrn;

USE rtrn;

CREATE TABLE callbacks (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  url VARCHAR(4096),
  method VARCHAR(20),
  post_data INTEGER,
  retry_phases INTEGER
);

CREATE TABLE post_data (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  post_data_headers INTEGER
);

ALTER TABLE callbacks ADD CONSTRAINT post_data_in_callbacks FOREIGN KEY callbacks (post_data) REFERENCES post_data (id);

CREATE TABLE post_data_headers (
  id INTEGER AUTO_INCREMENT PRIMARY KEY,
  header VARCHAR(128),
  value VARCHAR(1024)
);

ALTER TABLE post_data ADD CONSTRAINT post_data_headers_in_post_data FOREIGN KEY post_data (post_data_headers) REFERENCES post_data_headers (id);

CREATE TABLE retry_phases (
  id INTEGER AUTO_INCREMENT PRIMARY KEY

)
