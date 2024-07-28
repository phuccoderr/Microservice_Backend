package com.phuc.categoryservice.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.phuc.categoryservice.dtos.CategoryDto;
import com.phuc.categoryservice.dtos.PaginationDto;
import com.phuc.categoryservice.exceptions.ParamValidateException;

import java.util.List;

public interface ICategoryRedisService {

    void clear();
    PaginationDto getAllCategories(Integer page, Integer limit, String sort) throws ParamValidateException, JsonProcessingException;
    void saveAllCategories(PaginationDto paginationDto, Integer page, Integer limit, String sort) throws JsonProcessingException;
}
