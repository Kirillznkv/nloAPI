create table if not exists anomaly (
    id bigserial not null primary key,
    SessionId text not null,
    Frequency double precision not null,
    Timestamp timestamp not null
);