package com.phuc.productservice.exceptions;

public class ParamValidateException extends Exception {
    public ParamValidateException() {
        super("Param error sort: must be asc or desc");
    }
}
