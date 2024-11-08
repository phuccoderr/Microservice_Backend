package com.phuc.productservice.constants;

import lombok.experimental.UtilityClass;

@UtilityClass
public class Constants {

    public static final String API_PRODUCTS = "/api/v1/products";
    public static final String API_DISCOUNTS = "/api/v1/discounts";

    public static final String GET_ALL_SUCCESS = "Get products successfully!";
    public static final String GET_SUCCESS = "Get product successfully!";
    public static final String CREATE_SUCCESS = "Create product successfully!";
    public static final String UPDATE_SUCCESS = "Update product successfully!";
    public static final String DELETE_SUCCESS = "Delete product successfully!";
    public static final String ADD_FILES_SUCCESS = "Add files product successfully!";
    public static final String DELETE_FILES_SUCCESS = "Delete files product successfully!";


    public static final String CREATE_DISCOUNT_SUCCESS = "Create discount successfully!";
    public static final String GET_DISCOUNT_SUCCESS = "Get discount successfully!";
    public static final String DELETE_DISCOUNT_SUCCESS = "Delete discount successfully!";
    public static final String APPLY_DISCOUNT_SUCCESS = "Apply discount fail!";

    public static final String CLOUD_FOLDER = "microservice/product/";
    public static final String UPLOAD_FILE_FAIL = "File upload fail!";
    public static final String FILE_EXTENSION_FAIL = "Only jpg, png, gif, bmp files are allowed!";
    public static final String FILE_MAX_SIZE = "Max file size is 2MB!";

    public static final String DB_NOT_FOUND = "Data not Found!";
    public static final String DB_ALREADY_EXISTS = "Data already exists!";

    public static final String PARAM_SORT_FAIL = "Param error sort: must be asc or desc!";

    public static final String TOKEN_EXPIRED = "JWT token is expired!";
    public static final String TOKEN_INVALID = "Invalid JWT token!";



}
