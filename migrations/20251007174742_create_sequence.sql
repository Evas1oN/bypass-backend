-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE ipaddress AS smallint START WITH 1 INCREMENT BY 1 NO CYCLE
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SEQUENCE ipaddress
-- +goose StatementEnd
