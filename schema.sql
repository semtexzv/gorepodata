-- Your SQL goes here
CREATE TABLE IF NOT EXISTS arch
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT                              NOT NULL UNIQUE,
    CHECK (name <> '')
);

INSERT OR
REPLACE
INTO arch (name)
VALUES ('noarch'),
       ('i386'),
       ('i486'),
       ('i586'),
       ('i686'),
       ('alpha'),
       ('alphaev6'),
       ('ia64'),
       ('sparc'),
       ('sparcv9'),
       ('sparc64'),
       ('s390'),
       ('athlon'),
       ('s390x'),
       ('ppc'),
       ('ppc64'),
       ('ppc64le'),
       ('pSeries'),
       ('iSeries'),
       ('x86_64'),
       ('ppc64iseries'),
       ('ppc64pseries'),
       ('ia32e'),
       ('amd64'),
       ('aarch64'),
       ('armv7hnl'),
       ('armv7hl'),
       ('armv7l'),
       ('armv6hl'),
       ('armv6l'),
       ('armv5tel'),
       ('src');

CREATE TABLE IF NOT EXISTS content_set
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    label TEXT                              NOT NULL CHECK ( label <> '' ),
    name  TEXT                              NOT NULL CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS repo
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    content_set_id INTEGER                           NOT NULL,
    url            TEXT                              NOT NULL CHECK ( url != '' ),
    basearch       TEXT                              NOT NULL CHECK ( basearch <> '' ),
    releasever     TEXT                              NOT NULL CHECK ( releasever <> '' ),
    revision       INTEGER,

    FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE IF NOT EXISTS package_name
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT                              NOT NULL CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS evr
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    epoch   TEXT                              NOT NULL CHECK ( epoch <> '' ),
    version TEXT                              NOT NULL CHECK ( version <> '' ),
    rel     TEXT                              NOT NULL CHECK ( rel <> '' )
);

CREATE TABLE IF NOT EXISTS package
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name_id INTEGER                           NOT NULL,
    evr_id  INTEGER                           NOT NULL,
    arch_id INTEGER                           NOT NULL,
    FOREIGN KEY (name_id) REFERENCES package_name (id),
    FOREIGN KEY (evr_id) REFERENCES evr (id),
    FOREIGN KEY (arch_id) REFERENCES arch (id)
);

CREATE TABLE IF NOT EXISTS package_repo
(
    package_id INTEGER NOT NULL,
    repo_id    INTEGER NOT NULL,
    FOREIGN KEY (package_id) REFERENCES package (id),
    FOREIGN KEY (repo_id) REFERENCES repo (id),
    PRIMARY KEY (package_id, repo_id)
) WITHOUT ROWID;


CREATE TABLE IF NOT EXISTS advisory_type
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS advisory_severity
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS advisory
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT      NOT NULL UNIQUE CHECK ( name <> '' ),

    type_id     INTEGER   NOT NULL,
    severity_id INTEGER,

    synopsis    TEXT,
    summary     TEXT,
    description TEXT,
    solution    TEXT,
    issued      TIMESTAMP NOT NULL,
    updated     TIMESTAMP NOT NULL,

    FOREIGN KEY (type_id) REFERENCES advisory_type (id),
    FOREIGN KEY (severity_id) REFERENCES advisory_severity (id)
);

CREATE TABLE IF NOT EXISTS advisory_repo
(
    advisory_id INTEGER NOT NULL,
    repo_id     INTEGER NOT NULL,

    FOREIGN KEY (advisory_id) REFERENCES advisory (id),
    FOREIGN KEY (repo_id) REFERENCES repo (id),
    PRIMARY KEY (advisory_id, repo_id)
);




-- Your SQL goes here
CREATE TABLE IF NOT EXISTS arch
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT                              NOT NULL UNIQUE,
    CHECK (name <> '')
);

INSERT OR
REPLACE
INTO arch (name)
VALUES ('noarch'),
       ('i386'),
       ('i486'),
       ('i586'),
       ('i686'),
       ('alpha'),
       ('alphaev6'),
       ('ia64'),
       ('sparc'),
       ('sparcv9'),
       ('sparc64'),
       ('s390'),
       ('athlon'),
       ('s390x'),
       ('ppc'),
       ('ppc64'),
       ('ppc64le'),
       ('pSeries'),
       ('iSeries'),
       ('x86_64'),
       ('ppc64iseries'),
       ('ppc64pseries'),
       ('ia32e'),
       ('amd64'),
       ('aarch64'),
       ('armv7hnl'),
       ('armv7hl'),
       ('armv7l'),
       ('armv6hl'),
       ('armv6l'),
       ('armv5tel'),
       ('src');

CREATE TABLE IF NOT EXISTS content_set
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    label TEXT                              NOT NULL CHECK ( label <> '' ),
    name  TEXT                              NOT NULL CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS repo
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    content_set_id INTEGER                           NOT NULL,
    url            TEXT                              NOT NULL CHECK ( url != '' ),
    basearch       TEXT                              NOT NULL CHECK ( basearch <> '' ),
    releasever     TEXT                              NOT NULL CHECK ( releasever <> '' ),
    revision       INTEGER,

    FOREIGN KEY (content_set_id) REFERENCES content_set (id)
);

CREATE TABLE IF NOT EXISTS package_name
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT                              NOT NULL CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS evr
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    epoch   TEXT                              NOT NULL CHECK ( epoch <> '' ),
    version TEXT                              NOT NULL CHECK ( version <> '' ),
    rel     TEXT                              NOT NULL CHECK ( rel <> '' )
);

CREATE TABLE IF NOT EXISTS package
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name_id INTEGER                           NOT NULL,
    evr_id  INTEGER                           NOT NULL,
    arch_id INTEGER                           NOT NULL,
    FOREIGN KEY (name_id) REFERENCES package_name (id),
    FOREIGN KEY (evr_id) REFERENCES evr (id),
    FOREIGN KEY (arch_id) REFERENCES arch (id)
);

CREATE TABLE IF NOT EXISTS package_repo
(
    package_id INTEGER NOT NULL,
    repo_id    INTEGER NOT NULL,
    FOREIGN KEY (package_id) REFERENCES package (id),
    FOREIGN KEY (repo_id) REFERENCES repo (id),
    PRIMARY KEY (package_id, repo_id)
) WITHOUT ROWID;


CREATE TABLE IF NOT EXISTS advisory_type
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS advisory_severity
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS advisory
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT      NOT NULL UNIQUE CHECK ( name <> '' ),

    type_id     INTEGER   NOT NULL,
    severity_id INTEGER,

    synopsis    TEXT,
    summary     TEXT,
    description TEXT,
    solution    TEXT,
    issued      TIMESTAMP NOT NULL,
    updated     TIMESTAMP NOT NULL,

    FOREIGN KEY (type_id) REFERENCES advisory_type (id),
    FOREIGN KEY (severity_id) REFERENCES advisory_severity (id)
);

CREATE TABLE IF NOT EXISTS advisory_repo
(
    advisory_id INTEGER NOT NULL,
    repo_id     INTEGER NOT NULL,

    FOREIGN KEY (advisory_id) REFERENCES advisory (id),
    FOREIGN KEY (repo_id) REFERENCES repo (id),
    PRIMARY KEY (advisory_id, repo_id)
);




