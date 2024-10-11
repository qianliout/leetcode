use stack;
create table names
(
    id           bigint       not null auto_increment primary key,
    code         varchar(10)  not null default '',
    name         varchar(100) not null default '',
    price        bigint       not null default 0,
    plate        varchar(200) not null default '',
    industry     varchar(200) not null default '',
    exchange     varchar(200) not null default '',
    crawl_at     bigint       not null default 0,
    `created_at` bigint       not null default 0,
    `updated_at` bigint       not null default 0,
    unique key `idx_code` (`code`)
);

create table profiles
(
    id                 bigint not null auto_increment primary key,
    name               varchar(100),
    code               varchar(10),
    reporting_period   varchar(20),
    operate_all_income bigint,
    operate_income     bigint,
    operate_all_cost   bigint,
    operate_cost       bigint,
    tax                bigint,
    sales_expense      bigint,
    manage_expense     bigint,
    financial_expense  bigint,
    rd_expense         bigint,
    diluted_earn       float,
    unique key ` idx_code ` (` code `, ` reporting_period `)
);

create table balances
(
    id               bigint not null auto_increment primary key,
    name             varchar(20),
    code             varchar(10),
    reporting_period varchar(50),
    money_funds      bigint,
    trans_finance    bigint,
    stock            bigint,
    short_loan       bigint,
    long_loan        bigint,
    capital          bigint,
    unique key ` idx_code ` (` code `, ` reporting_period `)
);


