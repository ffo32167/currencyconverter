CREATE SCHEMA rates AUTHORIZATION postgres;

CREATE TABLE rates.rates (
	rate_date date NOT NULL,
	curr_code bpchar(3) NOT NULL,
	rate float8 NOT NULL
);

CREATE UNIQUE INDEX rates_rate_date_idx ON rates.rates USING btree (rate_date,curr_code) INCLUDE (rate);


INSERT INTO rates.rates (rate_date,curr_code,rate) VALUES
	 ('2021-09-12','RUB',73.17),
	 ('2021-09-12','EUR',0.85),
	 ('2021-09-12','JPY',109.94),
	 ('2021-09-11','RUB',72.77),
	 ('2021-09-11','EUR',0.88),
	 ('2021-09-11','JPY',106.94),
	 ('2021-09-14','USD',1.0),
	 ('2021-09-14','RUB',72.679),
	 ('2021-09-14','EUR',0.846769),
	 ('2021-09-14','JPY',109.99657143);
INSERT INTO rates.rates (rate_date,curr_code,rate) VALUES
	 ('2002-03-02','RUB',1.0),
	 ('2002-03-02','USD',30.9436),
	 ('2002-03-02','EUR',26.8343),
	 ('2002-03-02','JPY',0.231527);