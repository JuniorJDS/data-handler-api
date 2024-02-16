DROP TABLE if exists userData;
CREATE TABLE userData (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(14) NOT NULL,
    private BOOLEAN NOT NULL,
    incompleto BOOLEAN NOT NULL,
    dataDaUltimaCompra DATE,
    ticketMedio NUMERIC(10, 2),
	ticketDaUltimaCompra NUMERIC(10, 2),
	lojaMaisFrequente VARCHAR(18),
	lojaDaUltimaCompra VARCHAR(18),
    isValidCPForCNPJ BOOLEAN NOT NULL
);
