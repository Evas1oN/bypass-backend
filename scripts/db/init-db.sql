create user bypass with encrypted password 'bypass';
create database bypass with owner bypass;
\c bypass
create schema authorization bypass;
alter user bypass set search_path = 'bypass';