create table currencies (id string, code string, type string, dollar_based_proportion number);

insert into currencies (id, code, type, dollar_based_proportion) values ('1', 'USD', 'FIAT', 0);
insert into currencies (id, code, type, dollar_based_proportion) values ('2', 'BRL', 'FIAT', 0);
insert into currencies (id, code, type, dollar_based_proportion) values ('3', 'EUR', 'FIAT', 0);
insert into currencies (id, code, type, dollar_based_proportion) values ('4', 'BTC', 'CRYPTO', 0);
insert into currencies (id, code, type, dollar_based_proportion) values ('5', 'ETH', 'CRYPTO', 0);
insert into currencies (id, code, type, dollar_based_proportion) values ('6', 'RUAN', 'FICTICIOUS', 1);
