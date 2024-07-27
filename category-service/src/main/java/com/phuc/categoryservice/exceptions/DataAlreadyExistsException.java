package com.phuc.categoryservice.exceptions;

public class DataAlreadyExistsException extends Exception{

    public DataAlreadyExistsException() {
        super("Data already exists!");
    }
    public DataAlreadyExistsException(String message) {
        super(message);
    }
}
