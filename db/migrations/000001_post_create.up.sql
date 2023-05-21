CREATE TABLE post
(
    id             BIGSERIAL    PRIMARY KEY,

    title               TEXT         NOT NULL,
    content             TEXT         NULL,
    media_url           TEXT         NULL,
    likes               INT          NOT NULL,

    created_at          TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at          TIMESTAMPTZ  NOT NULL DEFAULT now()
);
