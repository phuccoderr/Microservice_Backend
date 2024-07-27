package com.phuc.categoryservice.exceptions;

public class DataNotFoundException extends Exception {
    public DataNotFoundException() {
        super("Data not found!");
    }
}
