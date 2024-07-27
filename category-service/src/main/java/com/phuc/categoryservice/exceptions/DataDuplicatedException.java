package com.phuc.categoryservice.exceptions;

public class DataDuplicatedException extends  Exception{
    public DataDuplicatedException() {
        super("Data duplicated parentId and Id");
    }
}
