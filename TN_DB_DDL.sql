-- MySQL Script generated by MySQL Workbench
-- Tue 28 Aug 2018 02:14:14 PM WIB
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

/* SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0; */
/* SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0; */
/* SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES'; */

-- -----------------------------------------------------
-- Schema public
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Table `public`.`account`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS public.account (
  account_number VARCHAR(50) NOT NULL ,
  email VARCHAR(100) NOT NULL,
  create_date TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  modi_date TIMESTAMP(0) NULL ,
  PRIMARY KEY (account_number))
;


-- -----------------------------------------------------
-- Table `public`.`tx_deposit`
-- -----------------------------------------------------
CREATE SEQUENCE public.tx_deposit_seq;

CREATE TABLE IF NOT EXISTS public.tx_deposit (
  seqno_deposit BIGINT NOT NULL DEFAULT NEXTVAL ('public.tx_deposit_seq') ,
  account_number VARCHAR(50) NOT NULL ,
  deposit_date DATE NOT NULL ,
  cash DOUBLE PRECISION NOT NULL ,
  create_date TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  modi_date TIMESTAMP(0) NULL ,
  PRIMARY KEY (seqno_deposit),
  CONSTRAINT fk_transaction_account
    FOREIGN KEY (account_number)
    REFERENCES public.account (account_number)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
;

CREATE INDEX fk_transaction_account_idx ON public.tx_deposit (account_number ASC);


/* SET SQL_MODE=@OLD_SQL_MODE; */
/* SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS; */
/* SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS; */

