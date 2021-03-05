-- name: migrate
CREATE TABLE organizations
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT DEFAULT 'null',
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL UNIQUE,
    site VARCHAR(255) UNIQUE DEFAULT 'null',
    coordinates VARCHAR(255) NOT NULL,
    office VARCHAR(255) DEFAULT 'null',
    date_creation TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    date_update TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_enable BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE organization_documents
(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    file TEXT NOT NULL,
    organization_id INT REFERENCES organizations(id) ON DELETE CASCADE NOT NULL,
    is_secure BOOLEAN DEFAULT FALSE
);

CREATE TABLE roles
(
    id SERIAL NOT NULL PRIMARY KEY,
    organization_id INT REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    edit_organization BOOLEAN DEFAULT FALSE,
    delete_organization BOOLEAN DEFAULT FALSE,
    create_service BOOLEAN DEFAULT FALSE,
    edit_service BOOLEAN DEFAULT FALSE,
    delete_service BOOLEAN DEFAULT FALSE,
    create_role BOOLEAN DEFAULT FALSE,
    edit_role BOOLEAN DEFAULT FALSE,
    delete_role BOOLEAN DEFAULT FALSE,
    create_employee BOOLEAN DEFAULT FALSE,
    edit_employee BOOLEAN DEFAULT FALSE,
    delete_employee BOOLEAN DEFAULT FALSE
);

CREATE TABLE organizations_users
(
    id SERIAL NOT NULL PRIMARY KEY,
    organization_id INT REFERENCES organizations(id) ON DELETE CASCADE NOT NULL ,
    user_id INT NOT NULL,
    role_id INT REFERENCES roles(id) ON DELETE SET NULL,
    confirmed BOOLEAN DEFAULT FALSE
);


-- name: drop
DROP TABLE organizations_users;
DROP TABLE roles;
DROP TABLE organization_documents;
DROP TABLE organizations;