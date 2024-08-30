package com.phuc.reviewservice.exeptions;

public class DataErrorException extends Exception {
    public DataErrorException(String message) {
        super(message);
    }
}
