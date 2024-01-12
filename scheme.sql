CREATE TABLE IF NOT EXISTS user(
    id INT,
    name string UNIQUE,
    password string
);

CREATE TABLE if NOT EXISTS person(
 id integer PRIMARY KEY,
 name VARCHAR,
 happiness int,
 experience INT,
 greed int,
 toxity int,
 courage int,
 currency float,
 owner_id int,
 FOREIGN KEY(owner_id) REFERENCES person(id) ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS company (
    id INTEGER PRIMARY KEY,
    title VARCHAR,
    sphere VARCHAR,
    currency FLOAT
);

CREATE TABLE IF NOT EXISTS education (
    id INTEGER PRIMARY key,
    title VARCHAR,
    grade VARCHAR
  
);
CREATE TABLE IF NOT EXISTS education_person (
    id INTEGER PRIMARY key,
    person_id INTEGER,
    education_id INTEGER,
    FOREIGN KEY(person_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(education_id) REFERENCES education(id) ON DELETE RESTRICT

);

CREATE TABLE IF NOT EXISTS vacancy (
    id INTEGER PRIMARY key,
    title VARCHAR,
    salary Float,
    company_id int,
    education_id int,
    FOREIGN KEY(company_id) REFERENCES company(id) ON DELETE RESTRICT,
    FOREIGN KEY(education_id) REFERENCES education(id) ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS loan (
    id INTEGER PRIMARY key,
    lender_id int,
    borrower_id int,
    amount FLOAT,
    term int,
    percent int,
    full_amount float,
    daily_payment float,
    FOREIGN KEY(lender_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(borrower_id) REFERENCES person(id) ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS job (
    id INTEGER PRIMARY key,
    salary float,
    vacancy_id int,
    person_id int,
    company_id int,
    FOREIGN KEY(vacancy_id) REFERENCES vacancy(id) ON DELETE RESTRICT,
    FOREIGN KEY(person_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(company_id) REFERENCES company(id) ON DELETE RESTRICT

);
CREATE TABLE IF NOT EXISTS estate (
    id INTEGER PRIMARY key,
    title VARCHAR,
    owner_id int,
    owner_type VARCHAR,
    cost FLOAT,
    loan_id int,
    FOREIGN KEY(loan_id) REFERENCES loan(id) ON DELETE RESTRICT

);

CREATE TABLE IF NOT EXISTS deal (
    id INTEGER PRIMARY key,
    cost float,
    seller_id int,
    buyer_id int,
    estate_id int,    
    FOREIGN KEY(seller_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(buyer_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(estate_id) REFERENCES estate(id) ON DELETE RESTRICT

);

CREATE TABLE IF NOT EXISTS loan (
    id INTEGER PRIMARY key,
    lender_id int,
    borrower_id int,
    amount FLOAT,    
    percent int,
    daily_payment float,
    FOREIGN KEY(lender_id) REFERENCES person(id) ON DELETE RESTRICT,
    FOREIGN KEY(borrower_id) REFERENCES person(id) ON DELETE RESTRICT
);



CREATE TABLE IF NOT EXISTS settings (
    id INTEGER PRIMARY key,
    owner_id int,
    bg_color VARCHAR,
    FOREIGN KEY(owner_id) REFERENCES user(id) ON DELETE RESTRICT
);