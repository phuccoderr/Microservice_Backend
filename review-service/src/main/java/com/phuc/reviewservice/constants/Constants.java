package com.phuc.reviewservice.constants;

import lombok.experimental.UtilityClass;

@UtilityClass
public class Constants {
    public static final String API_REVIEWS = "/api/v1/reviews";


    public static final String GET_ALL_SUCCESS = "Get reviews successfully!";
    public static final String DELETE_SUCCESS = "Delete review successfully!";

    public static final String DB_HAS_CHILDREN = "Category cannot be deleted because it has child categories!";
    public static final String DB_ALREADY_EXISTS = "Data already exists!";
    public static final String DB_DUPLICATED = "Data duplicated parentId and id!";
    public static final String DB_NOT_FOUND = "Data not found!";

    public static final String PARAM_SORT_FAIL = "Param error sort: must be asc or desc!";

    public static final String TOKEN_EXPIRED = "JWT token is expired!";
    public static final String TOKEN_INVALID = "Invalid JWT token!";

    public static final String TOPIC_AVG = "average-rating-events-topics";


}
