package com.phuc.reviewservice.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.phuc.reviewservice.dtos.PaginationDto;
import com.phuc.reviewservice.exeptions.ParamValidateException;

public interface IReviewRedisService {
    void clearAllReviews();
    void clearAllReviewsByProduct(String productId);
    PaginationDto getAllReviews(Integer page, Integer limit, String sort) throws ParamValidateException, JsonProcessingException;
    void saveAllReviews(PaginationDto paginationDto, Integer page, Integer limit, String sort) throws JsonProcessingException;
    PaginationDto getAllReviewsByProduct(String productId, Integer page, Integer limit) throws ParamValidateException, JsonProcessingException;
    void saveAllReviewsByProduct(PaginationDto paginationDto,String productId, Integer page, Integer limit) throws JsonProcessingException;
}
