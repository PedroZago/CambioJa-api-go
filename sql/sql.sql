-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema sistemaCambio
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema sistemaCambio
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sistemaCambio` DEFAULT CHARACTER SET utf8 ;
USE `sistemaCambio` ;

-- -----------------------------------------------------
-- Table `sistemaCambio`.`usuario`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sistemaCambio`.`usuario` (
  `usuarioID` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(50) NOT NULL,
  `email` VARCHAR(50) NOT NULL UNIQUE,
  `senha` VARCHAR(100) NOT NULL,
  `sexo` VARCHAR(50) NOT NULL,
  `eAdmin` INT NULL,
  `dataCriacao` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (`usuarioID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sistemaCambio`.`moeda`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sistemaCambio`.`moeda` (
  `moedaID` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(50) NOT NULL UNIQUE,
  `codISO` VARCHAR(10) NOT NULL UNIQUE,
  `cotacao` DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (`moedaID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sistemaCambio`.`cambio`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sistemaCambio`.`cambio` (
  `cambioID` INT NOT NULL AUTO_INCREMENT,
  `dataCambio` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
  `valorTransferido` DECIMAL(10,2) NOT NULL,
  `resultadoConversao` DECIMAL(10,2) NOT NULL,
  `usuarioID` INT NOT NULL,
  `moedaID` INT NOT NULL,
  PRIMARY KEY (`cambioID`),
  INDEX `fk_cambio_1_idx` (`usuarioID` ASC) VISIBLE,
  INDEX `fk_cambio_2_idx` (`moedaID` ASC) VISIBLE,
  CONSTRAINT `fk_cambio_1`
    FOREIGN KEY (`usuarioID`)
    REFERENCES `sistemaCambio`.`usuario` (`usuarioID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_cambio_2`
    FOREIGN KEY (`moedaID`)
    REFERENCES `sistemaCambio`.`moeda` (`moedaID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
