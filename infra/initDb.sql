DROP TABLE if exists userData;
CREATE TABLE userData (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(18) NOT NULL,
    private BOOLEAN NOT NULL,
    incompleto BOOLEAN NOT NULL,
    dataDaUltimaCompra VARCHAR(10),
    ticketMedio NUMERIC(10, 2),
	ticketDaUltimaCompra NUMERIC(10, 2),
	lojaMaisFrequente VARCHAR(18),
	lojaDaUltimaCompra VARCHAR(18)
);
