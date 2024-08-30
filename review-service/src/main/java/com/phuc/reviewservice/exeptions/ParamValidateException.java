package com.phuc.reviewservice.exeptions;

public class ParamValidateException extends Exception {
    public ParamValidateException(String message) {
        super(message);
    }
}