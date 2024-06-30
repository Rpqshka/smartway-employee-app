CREATE TABLE employees
(
    id                  SERIAL      PRIMARY KEY,
    name                VARCHAR(31) NOT NULL,
    surname             VARCHAR(63) NOT NULL,
    employee_phone      VARCHAR(15) NOT NULL,
    company_id          INT         NOT NULL,
    passport_type       VARCHAR(31) NOT NULL,
    passport_number     VARCHAR(15) NOT NULL,
    department_name     VARCHAR(31) NOT NULL,
    department_phone    VARCHAR(15) NOT NULL
);