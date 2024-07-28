package com.phuc.categoryservice.exceptions;

public class DataHasChildrenException extends Exception {
    public DataHasChildrenException() {
        super("Category cannot be deleted because it has child categories.");
    }
}
