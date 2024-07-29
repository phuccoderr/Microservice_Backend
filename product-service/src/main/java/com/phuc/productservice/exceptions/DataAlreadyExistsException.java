package com.phuc.productservice.exceptions;

public class DataAlreadyExistsException extends Exception{
    public DataAlreadyExistsException() {
        super("Data already exists!");
    }
}
