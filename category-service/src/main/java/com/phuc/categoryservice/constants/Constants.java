package com.phuc.categoryservice.constants;

import lombok.experimental.UtilityClass;

@UtilityClass
public class Constants {
    public static final String API_CATEGORIES = "/api/v1/categories";

    public static final String CREATE_SUCCESS = "Success create category!";
    public static final String GET_ALL_SUCCESS = "Get all categories successfully!";
    public static final String GET_SUCCESS = "Get category successfully!";
    public static final String UPDATE_SUCCESS = "Update category successfully!";
    public static final String DELETE_SUCCESS = "Delete category successfully!";

    public static final String DB_HAS_CHILDREN = "Category cannot be deleted because it has child categories!";
    public static final String DB_ALREADY_EXISTS = "Data already exists!";
    public static final String DB_DUPLICATED = "Data duplicated parentId and id!";
    public static final String DB_NOT_FOUND = "Data not found!";

    public static final String TOKEN_EXPIRED = "JWT token is expired!";
    public static final String TOKEN_INVALID = "Invalid JWT token!";


}
