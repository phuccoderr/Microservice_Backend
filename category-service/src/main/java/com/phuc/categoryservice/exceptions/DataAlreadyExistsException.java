package com.phuc.categoryservice.exceptions;

public class DataAlreadyExistsException extends Exception{
    public DataAlreadyExistsException() {
        super("Data already exists!");
    }
}
