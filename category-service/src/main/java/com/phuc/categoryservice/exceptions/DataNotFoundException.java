package com.phuc.categoryservice.exceptions;

import com.phuc.categoryservice.constants.Constants;

public class DataNotFoundException extends Exception {
    public DataNotFoundException() {
        super(Constants.DB_NOT_FOUND);
    }
}
